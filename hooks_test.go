package hooks

import (
	"log"
	"math/rand"
	"testing"
	"time"
)


func BenchmarkDispatch(b *testing.B) {
	type counterState struct {
		count int
	}
	Hook := New(counterState{0})
	Hook.UseReducer(func(state UseState, action Action) UseState {
		log.Println(state)
		switch action.ID {
		case "increment":
			return counterState{state.(counterState).count +1}
		default:
			return state
		}
	})

	for i := 0; i < b.N; i++ {
		Hook.Dispatch(Action{"increment", nil})
	}
}
