package main

//kvgo -> key value go project
import (
	"fmt"
	m "redis-kvgo/middlewares"
	"redis-kvgo/store"
	"redis-kvgo/store/kv"
	"redis-kvgo/store/ttl"
	"strings"
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

func main() {
	fmt.Println("Gokv project")

	s := kv.NewStore(0)
	sLog := m.NewLoggingMiddleware(s)
	sMetric := m.NewMetricsMiddleware(sLog)

	ttlStore := ttl.NewTtlStore(10 * time.Second)
	ttlLog := m.NewLoggingMiddleware(ttlStore)
	ttlMetric := m.NewMetricsMiddleware(ttlLog)

	sMetric.Get("randomKey")   // [log] 16:49:43.358290 GET "randomKey" -> miss (key randomKey does not exists)
	ttlMetric.Get("randomKey") // [log] 16:49:43.358290 GET "randomKey" -> miss (key randomKey does not exists)
	fmt.Println("==== Store ====")
	sMetric.Report()
	fmt.Println("==== TtlStore ====")
	ttlMetric.Report()

	valStore, err1 := store.SetWithEncryption(sLog, "1", "1")
	valTtlStore, err2 := store.SetWithEncryption(ttlLog, "2", "nycolas")

	if err1 != nil || err2 != nil {
		fmt.Println(err1)
		fmt.Println(err2)
	}

	fmt.Println("store value: " + valStore)
	fmt.Println("ttlstore value: " + valTtlStore)

	fmt.Println(s)  // pointer to a store -> address
	fmt.Println(*s) // dereference -> value
	fmt.Println(&s) // address itself

	err := s.Set("name", "Ricardo")
	if err != nil {
		return
	}
	fmt.Println(*s) // dereference -> value

	s.Set("lastName", "Albe")
	fmt.Println(*s) // dereference -> value

	value, returned := s.Get("nam")
	fmt.Println(value, returned)

	fmt.Println("Renaming a key does not exists")
	s.Rename("go1", "go2")

	if _, ok := s.Pop("go1"); !ok {
		fmt.Println("Popping a key does not exists")
	}

	fmt.Println(s.GetData())
	s.Delete("name")
	fmt.Println(s.GetData())
	s.Delete("nam")
	fmt.Println(s.GetData())

	fmt.Println(0x62)

	// currying function
	squareFunc := selfMath(multiply)
	doubleFunc := selfMath(add)

	fmt.Println(squareFunc(5)) // prints 25
	fmt.Println(doubleFunc(5)) // prints 10

	msg := "dangshootheck"
	removeProfanity(&msg)
}

func removeProfanity(msg *string) {
	value := *msg
	value = strings.ReplaceAll(value, "dang", "****")
	value = strings.ReplaceAll(value, "shoot", "*****")
	value = strings.ReplaceAll(value, "heck", "+++++")
	*msg = value
	println(*msg)
}

func multiply(x, y int) int {
	return x * y
}
func add(x, y int) int {
	return x + y
}

func selfMath(mathFunc func(int, int) int) func(int) int {
	return func(x int) int {
		return mathFunc(x, x)
	}
}
