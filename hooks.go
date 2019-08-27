package hooks

import (
	"sync"
)

// UseState interface state dari hooks.
type UseState interface{}

// UseReducer berpungsi sebagai penerima actions dispatched
// gunakan Dispatch() dan updates state.
type UseReducer func(UseState, Action) UseState

// Action :  adalah triger data updates didalam hooks itu sendiri.
type Action struct {
	ID   string
	Data interface{}
}

// Hooks : immutable data of state.
// Keadaan Hooks dapat menerima dengan UseState()
// state bisa dirubah oleh UseReducer sebagai result dari Actions Dispatch.
type Hooks struct {
	mu      sync.RWMutex
	reducer UseReducer
	state   UseState
	update  func(UseState)
}

// New instantiates a new hooks Hooks.
func New(initialState UseState) *Hooks {
	st := Hooks{
		reducer: func(s UseState, a Action) UseState {
			return s
		},
		state: initialState,
	}
	return &st
}

// UseReducer sets hook's reducer function.
func (st *Hooks) UseReducer(r UseReducer) {
	st.reducer = r
}

// UseEffect sets Hooks's update func.
// dispatch akan mengopi state baru
func (st *Hooks) UseEffect(update func(UseState)) {
	st.update = update
}

// setState mengembalikan copy dari Hooks's map.
func (st *Hooks) setState() UseState {
	return st.state
}

// UseState mengembalikan dari copy dari state.
func (st *Hooks) UseState() UseState {
	st.mu.RLock()
	defer st.mu.RUnlock()
	return st.setState()
}

// Dispatch Action Hooks.
func (st *Hooks) Dispatch(action Action) {
	st.mu.Lock()
	defer st.mu.Unlock()
	st.state = st.reducer(st.setState(), action)
	if st.update != nil {
		st.update(st.setState())
	}
}
