package rice

import (
	"errors"
	"sync"
)

type IFSM interface{}
type StateType int
type EventType int
type Transition int
type CallbackType int

const (
	None CallbackType = iota
	BeforeEvent
	LeaveState
	EnterState
	AfterEvent
)

type EventDesc struct {
	Name EventType
	Src  []StateType
	Dst  StateType
}

type EventKey struct {
	event EventType
	src   StateType
}

type CallbackKey struct {
	target       any          // 状态或事件
	callbackType CallbackType // 动作类型
}

type Callback func(*Event)
type Callbacks map[string]Callback

// Event is the info that get passed as a reference in the callbacks.
type Event struct {
	FSM      *FSM
	Event    EventType
	Src      StateType
	Dst      StateType
	Err      error
	Args     []interface{}
	canceled bool
	async    bool
}

type FSM struct {
	previous    StateType
	current     StateType
	transitions map[EventKey]StateType
	callbacks   map[CallbackKey]Callback
	metadata    map[string]interface{}

	stateMu    sync.RWMutex
	eventMu    sync.Mutex
	metadataMu sync.RWMutex
}

func (fsm *FSM) Previous() StateType {
	return fsm.previous
}

func (fsm *FSM) Current() StateType {
	fsm.stateMu.RLock()
	defer fsm.stateMu.RUnlock()
	return fsm.current
}

func (fsm *FSM) Is(state StateType) bool {
	fsm.stateMu.RLock()
	defer fsm.stateMu.RUnlock()
	return state == fsm.current
}

func (f *FSM) Metadata(key string) (any, bool) {
	f.metadataMu.RLock()
	defer f.metadataMu.RUnlock()
	dataElement, ok := f.metadata[key]
	return dataElement, ok
}

func (fsm *FSM) Transition(event EventType) {
}

func (fsm *FSM) Event(eventType EventType, args ...any) error {
	fsm.eventMu.Lock()
	defer fsm.eventMu.Unlock()

	fsm.stateMu.RLock()
	defer fsm.stateMu.RUnlock()

	// 在 transitions 里查询 现态 + 事件 ==> 次态
	dst, ok := fsm.transitions[EventKey{event: eventType, src: fsm.current}]
	if !ok {
		return errors.New("no eventkey")
	}

	e := &Event{fsm, eventType, fsm.current, dst, nil, args, false, false}

	err := fsm.beforeEventCallbacks(e)
	if err != nil {
		return err
	}

	return nil
}

func (fsm *FSM) beforeEventCallbacks(e *Event) error {

	// 在 callbacks 里查询 事件+动作类型==>动作 是否存在，如果存在，就执行相应的动作
	if fn, ok := fsm.callbacks[CallbackKey{e.Event, BeforeEvent}]; ok {
		fn(e)
		if e.canceled {
			return errors.New("event canceled")
		}
	}

	return nil
}

func NewFSM(initState StateType, events []EventDesc, callbacks map[CallbackKey]Callback) *FSM {

	f := &FSM{
		previous:    0,
		current:     initState,
		transitions: make(map[EventKey]StateType),
		callbacks:   make(map[CallbackKey]Callback),
	}

	allEvents := make(map[EventType]bool)
	allStates := make(map[StateType]bool)
	for _, e := range events {
		for _, src := range e.Src {
			f.transitions[EventKey{e.Name, src}] = e.Dst
			allStates[src] = true
			allStates[e.Dst] = true
		}
		allEvents[e.Name] = true
	}

	f.callbacks = callbacks

	return f
}

type FSMManage struct{}

var (
	ErrInvalidTransition = errors.New("invalid transition")
)

// InvalidEventError is returned by FSM.Event() when the event cannot be called
// in the current state.
type InvalidEventError struct {
	Event string
	State string
}

func (e InvalidEventError) Error() string {
	return "event " + e.Event + " inappropriate in current state " + e.State
}

// UnknownEventError is returned by FSM.Event() when the event is not defined.
type UnknownEventError struct {
	Event string
}

func (e UnknownEventError) Error() string {
	return "event " + e.Event + " does not exist"
}

// InTransitionError is returned by FSM.Event() when an asynchronous transition
// is already in progress.
type InTransitionError struct {
	Event string
}

func (e InTransitionError) Error() string {
	return "event " + e.Event + " inappropriate because previous transition did not complete"
}

// NotInTransitionError is returned by FSM.Transition() when an asynchronous
// transition is not in progress.
type NotInTransitionError struct{}

func (e NotInTransitionError) Error() string {
	return "transition inappropriate because no state change in progress"
}

// NoTransitionError is returned by FSM.Event() when no transition have happened,
// for example if the source and destination states are the same.
type NoTransitionError struct {
	Err error
}

func (e NoTransitionError) Error() string {
	if e.Err != nil {
		return "no transition with error: " + e.Err.Error()
	}
	return "no transition"
}

// CanceledError is returned by FSM.Event() when a callback have canceled a
// transition.
type CanceledError struct {
	Err error
}

func (e CanceledError) Error() string {
	if e.Err != nil {
		return "transition canceled with error: " + e.Err.Error()
	}
	return "transition canceled"
}

// AsyncError is returned by FSM.Event() when a callback have initiated an
// asynchronous state transition.
type AsyncError struct {
	Err error
}

func (e AsyncError) Error() string {
	if e.Err != nil {
		return "async started with error: " + e.Err.Error()
	}
	return "async started"
}

// InternalError is returned by FSM.Event() and should never occur. It is a
// probably because of a bug.
type InternalError struct{}

func (e InternalError) Error() string {
	return "internal error on state transition"
}
