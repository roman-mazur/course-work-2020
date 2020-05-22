package eventemitter

type Listener func(string, ...interface{})

type EventState struct {
	cb    Listener
	once  bool
	async bool
}
type Events map[string][]EventState

type EventEmitter struct {
	Events
	defaultMaxListeners int
}
