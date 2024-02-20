package emitter

import (
	"fmt"
	"log"
	"sync"
)

type HandlerCallback[T any] func(t T) error

type EventMap struct {
	lock            sync.RWMutex
	listeners       map[string][]HandlerCallback[any]
	ActiveListeners int
}

var listenersMap *EventMap

func New() *EventMap {
	*listenersMap = EventMap{
		lock:            sync.RWMutex{},
		listeners:       make(map[string][]HandlerCallback[any]),
		ActiveListeners: 0,
	}

	return listenersMap
}

// AddListener adds a listener to the map of listeners
// based on the handler type
func AddListener[T any](h HandlerCallback[T]) {
	name := fmt.Sprintf("%T", *new(T))
	// We lock since we are modifying the map
	listenersMap.lock.Lock()
	defer listenersMap.lock.Unlock()

	if ls := listenersMap.listeners; ls == nil {
		listenersMap.listeners = make(map[string][]HandlerCallback[any])
		listenersMap.listeners[name] = make([]HandlerCallback[any], 0)
	}

	if ok := listenersMap.listeners[name]; ok == nil {
		listenersMap.listeners[name] = make([]HandlerCallback[any], 1)
		listenersMap.ActiveListeners++
	}

	listenersMap.listeners[name] = append(listenersMap.listeners[name], wrap(h))
}

// Emit emits an event to all listeners using the callback type
// of the event as name to find the handlers
func Emit[T any](t T) {
	name := fmt.Sprintf("%T", t)
	handlers := listenersMap.listeners[name]

	for _, v := range handlers {
		go func(f HandlerCallback[any]) {
			err := f(t)
			if err != nil {
				log.Fatalf("Error: %v", err)
			}
		}(v)
	}

}

// wraps the handler callback to a generic type
// https://github.com/joaofnds/chev/blob/main/chev.go#L105
func wrap[T any](f HandlerCallback[T]) HandlerCallback[any] {
	return func(t any) error { return f(t.(T)) }
}
