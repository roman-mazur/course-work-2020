**<h1>Event Emitter</h1>**

## **<h2>Method: 'GetMaxListeners'</h2>**

- Returns: _int_

<span >Returns the current max listener count for the EventEmitter which is either set by [SetMaxListeners(n](#h2method-setmaxlistenersh2)) or defaults value(10)</span>

```golang
  ...
count := ee.GetMaxListeners()
  ...
```

## **<h2>Method: 'SetMaxListeners'</h2>**

- n _int_
- Returns: _error_ | _nil_

<span>The SetMaxListeners(n) method limits listener's count for EventEmitter instance</span>

```golang
  ...
err := ee.SetMaxListeners(15)
if err != nil {
	errorHandler(err)
}  ...
```

## **<h2>Method: 'Once'</h2>**

- name _string_
- cb _Listener_
- async _bool_
- Returns: _error_ | _nil_

<span>Adds a one-time listener function for passed event. The next time event is triggered, this listener is removed and then invoked. The listener marked as async, depending on last arg "async"</span>

```golang
  ...
err := ee.Once("event", func(name string, data ...interface{}) {
	fmt.Println(data)
}, false)
if err != nil {
	errorHandler(err)
}
  ...
```

## **<h2>Method: 'On'</h2>**

- name _string_
- cb _Listener_
- async _bool_
- Returns: _error_ | _nil_

<span>Adds a listener function to Events map for the passed event. The next time event is triggered, this listener will be invoked, in the order it was registered, multiple times. The listener marked as async, depending on last arg "async"</span>

```golang
  ...
err := ee.On("event", func(name string, data ...interface{}) {
	fmt.Println(data)
}, false)
if err != nil {
	errorHandler(err)
}
  ...
```

## **<h2>Method: 'EventNames'</h2>**

- Returns: _[]string_

<span >Returns a string array listing the events for which the EventEmitter has registered listeners</span>

```golang
  ...
events := ee.EventNames()
  ...
```

## **<h2>Method: 'ListenerCount'</h2>**

- name _string_
- Returns: _int_

<span >Returns the count of listeners listening to the passed event</span>

```golang
  ...
count := ee.ListenerCount("event")
  ...
```

## **<h2>Method: 'RemoveAllListeners'</h2>**

- name _string_
- Returns: _error | nil_

<span >Removes all listeners for passed event</span>

```golang
  ...
err := ee.RemoveAllListeners("event")
if err != nil {
	errorHandler(err)
}
  ...
```

## **<h2>Method: 'RemoveListener'</h2>**

- name _string_
- cb _Listener_
- Returns: _error | nil_

<span >Removes the passed listener from the listener array for the passed event</span>

```golang
  ...
err := ee.RemoveListener("event", cb)
if err != nil {
	errorHandler(err)
}
  ...
```

## **<h2>Method: 'Emit'</h2>**

- name _string_
- data _interface{}_
- Returns: _error | nil_

<span >Calls each of the listeners registered for the passed event, in the order they were registered, passing the supplied data to each. Data may be any Type. Depending on listener's mark "async", listeners calls in common or like goroutine</span>

```golang
  ...
err := ee.Emit("event", struct{ msg string }{"event data"})
if err != nil {
	errorHandler(err)
}
  ...
```

## **<h2>Method: 'NewEventEmitter'</h2>**

- Returns: _\*EventEmitter_

<span >Returns new EventEmitter instance</span>

```golang
  ...
ee := NewEventEmitter()
  ...
```

## **<h2>Method: 'AddListener'</h2>**

- name _string_
- cb _Listener_
- async _bool_
- once _bool_
- Returns: _error_ | _nil_

<span >Helper for [On](#h2method-onh2), [Once](#h2method-onceh2) Methods</span>
