package eventemitter

import (
	"fmt"
	"testing"
	"time"
)

func errorHandler(err error) {
	fmt.Print(err)
}

func cb1(name string, data ...interface{}) {
	fmt.Println(fmt.Sprintf("handler 1; Event: %s; Data: %v", []interface{}{name, data}...))
}
func cb2(name string, data ...interface{}) {
	fmt.Println(fmt.Sprintf("handler 2; Event: %s; Data: %v", []interface{}{name, data}...))
}
func cb3(name string, data ...interface{}) {
	fmt.Println(fmt.Sprintf("handler 3; Event: %s; Data: %v", []interface{}{name, data}...))
}
func cbLong(name string, data ...interface{}) {
	time.Sleep(time.Second)
	fmt.Println(fmt.Sprintf("cbLong; Event: %s; Data: %v", []interface{}{name, data}...))
}

func TestMaxListeners(t *testing.T) {
	ee := NewEventEmitter()
	err := ee.SetMaxListeners(15)
	if err != nil {
		errorHandler(err)
	}
	if ee.GetMaxListeners() != 15 {
		t.Error("Max listeners must be 15")
	}
}
func TestOn(t *testing.T) {
	ee := NewEventEmitter()
	err := ee.On("event1", cb1, false)
	if err != nil {
		errorHandler(err)
	}
	if len(ee.EventNames()) != 1 {
		t.Error("Must be 1 event")
	}
}

func TestOnAsync(t *testing.T) {
	ee := NewEventEmitter()
	err := ee.On("event1", cbLong, true)
	if err != nil {
		errorHandler(err)
	}
	err = ee.On("event1", cb2, false)
	if err != nil {
		errorHandler(err)
	}
	//cbLong started first
	//handler finished faster than cbLong
	ee.Emit("event1", struct{ A int }{6})
	count := ee.ListenerCount("event1")
	if count != 2 {
		t.Error("Must be 2 listener")
	}
}

func TestOnce(t *testing.T) {
	ee := NewEventEmitter()
	err := ee.Once("event1", cb1, false)
	if err != nil {
		errorHandler(err)
	}
	err = ee.On("event1", cb2, false)
	if err != nil {
		errorHandler(err)
	}
	count := ee.ListenerCount("event1")
	if count != 2 {
		t.Error("Must be 2 listeners")
	}
	ee.Emit("event1", struct{ A int }{5})
	count = ee.ListenerCount("event1")
	if count != 1 {
		t.Error("Must be 1 listener")
	}
}

func TestRemoveListener(t *testing.T) {
	ee := NewEventEmitter()
	err := ee.On("event1", cb1, false)
	if err != nil {
		errorHandler(err)
	}
	err = ee.On("event1", cb2, false)
	if err != nil {
		errorHandler(err)
	}
	err = ee.On("event1", cb3, false)
	if err != nil {
		errorHandler(err)
	}
	count := ee.ListenerCount("event1")
	if count != 3 {
		t.Error("Must be 3 listeners")
	}
	err = ee.RemoveListener("event1", cb1)
	if err != nil {
		errorHandler(err)
	}
	count = ee.ListenerCount("event1")
	if count != 2 {
		t.Error("Must be 2 listeners")
	}
}

func TestRemoveAllListener(t *testing.T) {
	ee := NewEventEmitter()
	err := ee.On("event1", cb1, false)
	if err != nil {
		errorHandler(err)
	}
	err = ee.On("event1", cb2, false)
	if err != nil {
		errorHandler(err)
	}
	err = ee.On("event1", cb3, false)
	if err != nil {
		errorHandler(err)
	}
	count := ee.ListenerCount("event1")
	if count != 3 {
		t.Error("Must be 3 listeners")
	}
	err = ee.RemoveAllListeners("event1")
	if err != nil {
		errorHandler(err)
	}
	count = ee.ListenerCount("event1")
	if count != 0 {
		t.Error("Must be 0 listeners")
	}
}
