**<h1>Promise</h1>**

## **<h2>Method: 'NewPromise'</h2>**

- executor _func(resolve Resolver, reject Rejecter)_
- Returns: _\*Promise_

<span >Returns new Promise instance</span>

```golang
  ...
p := NewPromise(func(resolve Resolver, reject Rejecter) {
		resolve("5")
	})
  ...
```

## **<h2>Method: 'Resolve'</h2>**

- result _interface{}_
- Returns: _\*Promise_

<span >The Resolve method returns a Promise that is resolved with a given value</span>

```golang
  ...
p := Resolve("5")
  ...
```

## **<h2>Method: 'Reject'</h2>**

- result _string_
- Returns: _\*Promise_

<span>The Reject method returns a Promise that is rejected with a given reason</span>

```golang
  ...
p := Reject("5")
  ...
```

## **<h2>Method: 'Then'</h2>**

- onFulfilled _func(data interface{}) interface{}_
- onRejected _func(err error) interface{}_
- Returns: _\*Promise_

<span>The Then method returns a Promise. It takes up to two arguments: callback functions for the success and failure cases of the Promise</span>

```golang
  ...
NewPromise(func(resolve Resolver, reject Rejecter) {
		reject(errors.New("5"))
	}).Then(func(data interface{}) interface{} {
		return nil
	}, func(err error) interface{} {
		return Resolve("6")
	}).Then(func(data interface{}) interface{} {
		return data
	}, func(err error) interface{} {
		return nil
	}).Then(func(data interface{}) interface{} {
		fmt.Println(data)
		return nil
	}, func(err error) interface{} {
		return nil
	})
  ...
```
