package main

import (
	"fmt"
	"log"
	"redis-kvgo/store"
	//"time"
	"sync"
)

// which parses and routes each Command to the appropriate store method.

type Command struct {
	Id    int
	Op    string
	Key   string
	Value string
}

// tiny switch which turns a Command into a store call
func dispatch(s store.Storer, c Command) {
	fmt.Printf("Dispatcher for Command %d starting\n", c.Id)
	switch c.Op {
	case "SET":
		log.Printf("SET key: %s, value: %s", c.Key, c.Value)
		if err := s.Set(c.Key, c.Value); err != nil {
			return
		}
	case "GET":
		log.Printf("GET key: %s", c.Key)
		if _, err := s.Get(c.Key); err != nil {
			return
		}
	}
	fmt.Printf("Dispatcher for Command %d done\n", c.Id)
}

// CommandOnBoot replays a slice of Commands against the store
func CommandOnBoot(s store.Storer, cmds []Command) {
	var wg sync.WaitGroup
	for _, c := range cmds {
		//go dispatch(s, c)
		wg.Go(func() { dispatch(s, c) })
	}

	wg.Wait()

	// workaround for synchronization -> sleep to simulate heavy task
	//time.Sleep(2 * time.Second)
}
