package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

type contextGlobal struct {
	sync.Mutex
	internal *context.CancelFunc
}

func newContextGlobal() *contextGlobal {
	_, cancel := context.WithCancel(context.Background())
	return &contextGlobal{
		internal: &cancel,
	}
}

func (c *contextGlobal) setContext(cancel context.CancelFunc) {
	c.Lock()
	c.internal = &cancel
	c.Unlock()
}

func (c *contextGlobal) getContext() (value context.CancelFunc) {
	c.Lock()
	value = *c.internal
	c.Unlock()
	return value
}

var userContext = newContextGlobal()

func main() {

	context, cancel := context.WithCancel(context.Background())
	userContext.setContext(cancel)

	i := 0
	go func() {
		for {
			select {
			case <-context.Done():
				{
					fmt.Printf("captured %v, exiting...\n", context.Done())
					fmt.Println("iterations: ", i)
					os.Exit(1)
					return
				}
			default:
				i = i + 1
			}
		}

	}()

	http.HandleFunc("/stop", trapSignal)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func trapSignal(w http.ResponseWriter, r *http.Request) {
	userContext.getContext()()
}
