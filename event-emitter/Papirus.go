package eventemitter

import "fmt"

func BAD_ARGS() string {
	return "Not valid passed args"
}
func NOT_LISTENERS(name string) string {
	return fmt.Sprintf("Listeners does not exist for %s event", name)
}
func NOT_LISTENER(name string) string {
	return fmt.Sprintf("Listener does not exist for %s event", name)
}
func SET_MAX_LISTENERS(args ...[]interface{}) string {
	return fmt.Sprintf("Already set maxListeners %d for event &s", args)
}
