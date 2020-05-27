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
