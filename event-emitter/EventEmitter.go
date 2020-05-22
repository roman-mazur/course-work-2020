package eventemitter

import (
	"errors"
	"reflect"
	"sync"
)

func (ee *EventEmitter) GetMaxListeners() int {
	return ee.defaultMaxListeners
}

func (ee *EventEmitter) SetMaxListeners(n int) error {
	if n < 0 {
		return errors.New(BAD_ARGS())
	}
	ee.defaultMaxListeners = n
	return nil
}

func (ee *EventEmitter) Once(name string, cb Listener, async bool) error {
	err := ee.AddListener(name, cb, async, true)
	return err
}

func (ee *EventEmitter) On(name string, cb Listener, async bool) error {
	err := ee.AddListener(name, cb, async, false)
	return err
}

func (ee *EventEmitter) EventNames() []string {
	keys := []string{}
	for ev := range ee.Events {
		keys = append(keys, ev)
	}
	return keys
}
func (ee *EventEmitter) ListenerCount(name string) int {
	if ee.Events[name] == nil {
		return 0
	}
	return len(ee.Events[name])
}
func (ee *EventEmitter) RemoveAllListeners(name string) error {
	if name == "" {
		return errors.New(BAD_ARGS())
	}
	if ee.Events[name] == nil {
		return errors.New(NOT_LISTENERS(name))
	}
	ee.Events[name] = nil
	return nil
}

func (ee *EventEmitter) RemoveListener(name string, cb Listener) error {
	if name == "" || cb == nil {
		return errors.New(BAD_ARGS())
	}
	if ee.Events[name] == nil {
		return errors.New(NOT_LISTENERS(name))
	}
	listeners := ee.Events[name]
	cbP := reflect.ValueOf(cb).Pointer()
	var q EventState
	for i, lis := range listeners {
		lisP := reflect.ValueOf(lis.cb).Pointer()
		if cbP == lisP {
			listeners[i] = listeners[len(listeners)-1]
			listeners[len(listeners)-1] = q
			listeners = listeners[:len(listeners)-1]
			ee.Events[name] = listeners
			return nil
		}
	}
	return errors.New(NOT_LISTENER(name))
}

func (ee *EventEmitter) Emit(name string, data ...interface{}) error {
	if name == "" || data == nil {
		return errors.New(BAD_ARGS())
	}
	var wg sync.WaitGroup
	for _, lis := range ee.Events[name] {
		if lis.async {
			asyncWrapper(&wg, lis.cb, name, data)
			continue
		} else {
			lis.cb(name, data)
		}
	}
	wg.Wait()
	for _, lis := range ee.Events[name] {
		if lis.once {
			ee.RemoveListener(name, lis.cb)
		}
	}

	return nil
}

func NewEventEmitter() *EventEmitter {
	ee := EventEmitter{
		make(Events),
		10,
	}
	return &ee
}

func asyncWrapper(wg *sync.WaitGroup, cb Listener, args ...interface{}) {
	wg.Add(1)
	params := make([]reflect.Value, len(args))
	for n := range args {
		params[n] = reflect.ValueOf(args[n])
	}
	go func() {
		defer wg.Done()
		reflect.ValueOf(cb).Call(params)
	}()
}

func (ee *EventEmitter) AddListener(name string, cb Listener, async, once bool) error {
	if name == "" || cb == nil {
		return errors.New(BAD_ARGS())
	}
	if len(ee.Events[name]) > ee.defaultMaxListeners {
		return errors.New(SET_MAX_LISTENERS([]interface{}{ee.defaultMaxListeners, name}))
	}

	ee.Events[name] = append(ee.Events[name], EventState{cb, once, async})
	return nil
}
