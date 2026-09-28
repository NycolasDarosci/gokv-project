package main

import (
	"log"
	"redis-kvgo/store"
	"time"
)

// which parses and routes each Command to the appropriate store method.

type Command struct {
	Op    string
	Key   string
	Value string
}

// tiny swith which turns a Command into a store call
func dispatch(s store.Storer, c Command) {
	switch c.Op {
	case "SET":
		log.Printf("SET data: %s-%s", c.Key, c.Value)
		if err := s.Set(c.Key, c.Value); err != nil {
			return
		}
	case "GET":
		log.Printf("GET key: %s", c.Key)
		if _, err := s.Get(c.Key); err != nil {
			return
		}
	}
}

func CommandOnBoot(s store.Storer, cmds []Command) {
	for _, c := range cmds {
		go dispatch(s, c)
	}

	time.Sleep(2 * time.Second)
}
