**<h1>Promise</h1>**

## **<h2>Method: 'NewPromise'</h2>**

- _executor func(resolve Resolver, reject Rejecter)_
- Returns: _\*Promise_

<span >Returns new Promise instance</span>

```golang
  ...
p := NewPromise(func(resolve Resolver, reject Rejecter) {
		resolve("5")
	})
  ...
```
