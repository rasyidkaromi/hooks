# hooks
Go Hooks dengan UseReducer, UseState, UseEffect dan Dispatch

Install
---------------
go get github.com/rasyidkaromi/hooks
        
sebuah Implementasi style font-end  

Usage
---------------

          import (
            "github.com/rasyidkaromi/hooks"
          )

          // Membuat Init state
          type InitState struct {
            count int
          }

          // Inisiasi New Hooks
          Hook := hooks.New(InitState{0})

          // Membuat a UseReducer dengan increments "count" ketika menerima "increment" dari Dispatch. 
          Hook.UseReducer(func(state hooks.UseState, action hooks.Action) hooks.UseState {
            switch action.ID {
            case "increment":
              return InitState{state.(InitState).count + action.Data.(int)}
            case "decrement":
              return InitState{state.(InitState).count - action.Data.(int)}
            default:
              return state
            }
          })

          Hook.Dispatch(Action{"increment", 5})
          Hook.Dispatch(Action{"decrement", 2})

          fmt.Println(Hook.UseState().(InitState).count) // prints 3

          // Register func UseEffect dipanggil setelah state lainnya update
          Hook.UseEffect(func(state UseState) {
            fmt.Println(state.(InitState).count)
          })
          Hook.Dispatch(Action{"decrement", 2})




dual core intel p6300 
goos: windows
goarch: amd64
BenchmarkDispatch-2 
 20000	     67184 ns/op	      16 B/op	       2 allocs/op
