package main

//kvgo -> key value go project
import (
	"encoding/base64"
	"fmt"
	"redis-kvgo/store/kv"
	"redis-kvgo/store/ttl"
	"time"
)

// interfaces in go are implicit
// that's mean whatever struct implements its signature, will be recognized as implementation of the interface
/*
	  Storer (Get() Set())
     ___|___
    |       |
TtlStore  Store
*/
type Storer interface {
	Set(key, value string) error
	Get(key string) (string, error)
}

func main() {
	fmt.Println("Gokv project")

	store := kv.NewStore(0)
	ttlStore := ttl.NewTtlStore(10 * time.Second)
	valStore, err1 := SetWithEncryption(store, "1", "1")
	valTtlStore, err2 := SetWithEncryption(ttlStore, "2", "nycolas")

	if err1 != nil || err2 != nil {
		fmt.Println(err1)
		fmt.Println(err2)
	}

	fmt.Println("store value: " + valStore)
	fmt.Println("ttlstore value: " + valTtlStore)

	fmt.Println(store)  // pointer to a store -> address
	fmt.Println(*store) // dereference -> value
	fmt.Println(&store) // address itself

	store.Set("name", "Ricardo")
	fmt.Println(*store) // dereference -> value

	store.Set("lastName", "Albe")
	fmt.Println(*store) // dereference -> value

	value, returned := store.Get("nam")
	fmt.Println(value, returned)

	fmt.Println("Renaming a key does not exists")
	store.Rename("go1", "go2")

	if _, ok := store.Pop("go1"); !ok {
		fmt.Println("Popping a key does not exists")
	}

	fmt.Println(store.GetData())
	store.Delete("name")
	fmt.Println(store.GetData())
	store.Delete("nam")
	fmt.Println(store.GetData())

	fmt.Println(0x62)
}

func SetWithEncryption(s Storer, key, value string) (string, error) {
	encrypted := base64.StdEncoding.EncodeToString([]byte(value))
	if err := s.Set(key, encrypted); err != nil {
		return "", err
	}
	return s.Get(key)
}
