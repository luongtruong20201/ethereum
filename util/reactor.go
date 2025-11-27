package util

import "sync"

type ReactorEvent struct {
	mutex sync.Mutex
	event string
	chans []chan React
}

func (e *ReactorEvent) Post(react React) {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	for _, ch := range e.chans {
		go func(ch chan React) {
			ch <- react
		}(ch)
	}
}

func (e *ReactorEvent) Add(ch chan React) {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	e.chans = append(e.chans, ch)
}

func (e *ReactorEvent) Remove(ch chan React) {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	for i, c := range e.chans {
		if c == ch {
			e.chans = append(e.chans[:i], e.chans[i+1:]...)
		}
	}
}

type React struct {
	Resource interface{}
	Event    string
}

type ReactorEngine struct {
	patterns map[string]*ReactorEvent
}

func NewReactorEngine() *ReactorEngine {
	return &ReactorEngine{patterns: make(map[string]*ReactorEvent)}
}

func (reactor *ReactorEngine) Subscribe(event string, ch chan React) {
	ev := reactor.patterns[event]
	if ev == nil {
		ev = &ReactorEvent{event: event}
		reactor.patterns[event] = ev
	}
	ev.Add(ch)
}

func (reactor *ReactorEngine) Unsubscribe(event string, ch chan React) {
	ev := reactor.patterns[event]
	if ev != nil {
		ev.Remove(ch)
	}
}

func (reactor *ReactorEngine) Post(event string, resource interface{}) {
	ev := reactor.patterns[event]
	if ev != nil {
		ev.Post(React{Resource: resource, Event: event})
	}
}
