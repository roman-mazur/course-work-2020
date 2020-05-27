package promise

import (
	"errors"
	"fmt"
	"time"
)

func checkOnError(value interface{}) error {
	if err, ok := value.(error); ok && err != nil {
		return err
	}
	return nil
}
func checkOnChan(value interface{}) chan interface{} {
	if ch, ok := value.(chan interface{}); ok && ch != nil {
		return ch
	}
	return nil
}

func (p *Promise) fulfill(result interface{}) {
	p.state = fulfilled
	p.value = result
	p.resolved = true
}
func (p *Promise) reject(err error) {
	p.state = rejected
	p.err = err
	p.resolved = true
}
func (p *Promise) resetState() {
	p.state = pending
}

func NewPromise(executor Executor) *Promise {
	p := &Promise{
		state:    pending,
		executor: executor,
		done:     make(chan bool, 1),
		resolved: false,
	}

	go func() {
		defer close(p.done)
		p.executor(p.fulfill, p.reject)
		p.done <- true
	}()
	<-p.done
	return p
}

func Resolve(result interface{}) *Promise {
	return NewPromise(func(resolve Resolver, reject Rejecter) {
		resolve(result)
	})
}

func Reject(result string) *Promise {
	return NewPromise(func(resolve Resolver, reject Rejecter) {
		reject(errors.New(result))
	})
}

func doResolve(resolve Resolver, reject Rejecter, cb func() interface{}, done chan bool) {
	go func() {
		defer func() {
			done <- true
			if err := recover(); err != nil {
				reject(errors.New(fmt.Sprint(err)))
			}
			close(done)
		}()
		response := cb()
		if p, isPromise := response.(*Promise); isPromise {
			response = await(p)
			if ch := checkOnChan(response); ch != nil {
				response = <-ch
			}
		}
		if err := checkOnError(response); err != nil {
			reject(err)
		} else {
			resolve(response)
		}
	}()
}

func (p *Promise) Then(onFulfilled OnFulfilled, onRejected OnRejected) *Promise {
	if p, ok := waitResolves([]*Promise{p}, 1, 1<<63-1); !ok && p != nil {
		return nil
	}
	done := make(chan bool, 1)
	next := NewPromise(func(resolve Resolver, reject Rejecter) {
		doResolve(resolve, reject, func() interface{} {
			if p.err != nil {
				return onRejected(p.err)
			} else {
				return onFulfilled(p.value)
			}
		}, done)
	})
	<-done
	next.resetState()
	return next
}
func await(p *Promise) chan interface{} {
	done := make(chan interface{}, 1)
	go func() {
		defer close(done)
		p.Then(func(data interface{}) interface{} {
			if pr, isPromise := data.(*Promise); isPromise {
				done <- await(pr)
			} else {
				if ch := checkOnChan(data); ch != nil {
					done <- <-ch
				} else {
					done <- data
				}
			}
			return nil
		}, func(err error) interface{} {
			done <- err
			return nil
		})
	}()
	return done
}
func waitResolves(promises []*Promise, waitCount int, timeout time.Duration) (*Promise, bool) {
	var pr *Promise
	start := time.Now()
	for i := 0; i < waitCount; i = 0 {
		if time.Since(start) > timeout {
			return pr, false
		}
		for _, p := range promises {
			if !p.resolved {
				continue
			}
			pr = p
			i++
		}
		if i == waitCount {
			break
		}
	}
	return pr, true
}

func All(promises []*Promise, timeout time.Duration) *Promise {
	if p, ok := waitResolves(promises, len(promises), timeout); !ok && p != nil {
		return nil
	}
	var result []interface{}
	var chans []chan interface{}
	for _, p := range promises {
		chans = append(chans, await(p))
	}
	for i := range chans {
		data := <-chans[i]
		if err := checkOnError(data); err != nil {
			return Reject(err.Error())
		}
		result = append(result, data)
	}
	return Resolve(result)
}
