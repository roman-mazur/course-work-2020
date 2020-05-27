package promise

const (
	pending = iota
	fulfilled
	rejected
)

type AllSettledResponse struct {
	status string
	value  interface{}
	reason string
}

type Promise struct {
	state    int
	value    interface{}
	err      error
	done     chan bool
	resolved bool
	executor Executor
}
type Executor func(resolve Resolver, reject Rejecter)
type Resolver func(value interface{})
type Rejecter func(err error)
type OnFulfilled func(data interface{}) interface{}
type OnRejected func(err error) interface{}
