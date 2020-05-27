package promise

import (
	"errors"
	"fmt"
	"testing"
	"time"
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

func TestAll(t *testing.T) {
	p1 := NewPromise(func(resolve Resolver, reject Rejecter) {
		time.AfterFunc(3*time.Second, func() {
			reject(errors.New("1"))
		})
	})
	p2 := NewPromise(func(resolve Resolver, reject Rejecter) {
		time.AfterFunc(1*time.Second, func() {
			resolve("2")
		})
	})
	p3 := Resolve("3")
	arrP := []*Promise{
		p1, p2, p3,
	}
	All(arrP, 5*time.Second).Then(func(data interface{}) interface{} {
		if data != nil {
			t.Error("Unexpected data")
		}
		return nil
	}, func(err error) interface{} {
		if err != nil && err.Error() != "1" {
			t.Error(fmt.Sprintf("Unexpected err: %v, must be 1", err))
		}
		return nil
	})
}

func TestRace(t *testing.T) {
	p1 := NewPromise(func(resolve Resolver, reject Rejecter) {
		time.AfterFunc(3*time.Second, func() {
			reject(errors.New("1"))
		})
	})
	p2 := NewPromise(func(resolve Resolver, reject Rejecter) {
		time.AfterFunc(1*time.Second, func() {
			resolve("2")
		})
	})
	p3 := Reject("3")
	arrP := []*Promise{
		p1, p2, p3,
	}
	Race(arrP, 5*time.Second).Then(func(data interface{}) interface{} {
		if data != nil {
			t.Error("Unexpected data")
		}
		return nil
	}, func(err error) interface{} {
		if err != nil && err.Error() != "3" {
			t.Error(fmt.Sprintf("Unexpected err: %v, must be 3", err))
		}
		return nil
	})
}

func TestAllSettled(t *testing.T) {
	p1 := NewPromise(func(resolve Resolver, reject Rejecter) {
		time.AfterFunc(3*time.Second, func() {
			reject(errors.New("1"))
		})
	})
	p2 := NewPromise(func(resolve Resolver, reject Rejecter) {
		time.AfterFunc(1*time.Second, func() {
			resolve("2")
		})
	})
	p3 := Reject("3")
	arrP := []*Promise{
		p1, p2, p3,
	}
	AllSettled(arrP, 5*time.Second).Then(func(data interface{}) interface{} {
		d := data.([]AllSettledResponse)
		if d[0].status != "rejected" || d[0].reason != "1" {
			t.Error(fmt.Sprintf("Unexpected data: %+v", d[0]))
		}
		if d[1].status != "fulfilled" || d[1].value != "2" {
			t.Error(fmt.Sprintf("Unexpected data: %+v", d[1]))
		}
		if d[2].status != "rejected" || d[2].reason != "3" {
			t.Error(fmt.Sprintf("Unexpected data: %+v", d[2]))
		}
		return nil
	}, func(err error) interface{} {
		if err != nil {
			t.Error("Unexpected err")
		}
		return nil
	})
}
