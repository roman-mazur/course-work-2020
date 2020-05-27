package promise

import (
	"errors"
	"fmt"
	"testing"
)

func TestResolve(t *testing.T) {
	Resolve("5").Then(func(data interface{}) interface{} {
		if data != "5" {
			t.Error(fmt.Sprintf("Unexpected data: %v, must be 5", data))
		}
		return nil
	}, func(err error) interface{} {
		if err != nil {
			t.Error("Unexpected error")
		}
		return nil
	})
}

func TestReject(t *testing.T) {
	Reject("5").Then(func(data interface{}) interface{} {
		if data != nil {
			t.Error("Unexpected data")
		}
		return nil
	}, func(err error) interface{} {
		if err != nil && err.Error() != "5" {
			t.Error(fmt.Sprintf("Unexpected err: %v, must be 5", err))
		}
		return nil
	})
}

func TestNewPromise(t *testing.T) {
	NewPromise(func(resolve Resolver, reject Rejecter) {
		resolve("5")
	}).Then(func(data interface{}) interface{} {
		if data != "5" {
			t.Error(fmt.Sprintf("Unexpected data: %v, must be 5", data))
		}
		return nil
	}, func(err error) interface{} {
		if err != nil {
			t.Error("Unexpected error")
		}
		return nil
	})

	NewPromise(func(resolve Resolver, reject Rejecter) {
		reject(errors.New("5"))
	}).Then(func(data interface{}) interface{} {
		if data != nil {
			t.Error("Unexpected data")
		}
		return nil
	}, func(err error) interface{} {
		if err != nil && err.Error() != "5" {
			t.Error(fmt.Sprintf("Unexpected err: %v, must be 5", err))
		}
		return nil
	})
}

func TestThenChain(t *testing.T) {
	NewPromise(func(resolve Resolver, reject Rejecter) {
		resolve("5")
	}).Then(func(data interface{}) interface{} {
		return data
	}, func(err error) interface{} {
		return nil
	}).Then(func(data interface{}) interface{} {
		return data
	}, func(err error) interface{} {
		return nil
	}).Then(func(data interface{}) interface{} {
		if data != "5" {
			t.Error(fmt.Sprintf("Unexpected data: %v, must be 5", data))
		}
		return nil
	}, func(err error) interface{} {
		return nil
	})

	NewPromise(func(resolve Resolver, reject Rejecter) {
		resolve("5")
	}).Then(func(data interface{}) interface{} {
		return data
	}, func(err error) interface{} {
		return nil
	}).Then(func(data interface{}) interface{} {
		panic("ooops")
	}, func(err error) interface{} {
		return nil
	}).Then(func(data interface{}) interface{} {
		if data != nil {
			t.Error("Unexpected data")
		}
		return nil
	}, func(err error) interface{} {
		if err != nil && err.Error() != "ooops" {
			t.Error(fmt.Sprintf("Unexpected err: %v, must be ooops", err))
		}
		return nil
	})
}

func TestThenReturnPromise(t *testing.T) {
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
		if data != "6" {
			t.Error(fmt.Sprintf("Unexpected data: %v, must be 6", data))
		}
		return nil
	}, func(err error) interface{} {
		return nil
	})

}
