package promise

import "errors"

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
