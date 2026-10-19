package main

//kvgo -> key value go project
import (
	"fmt"
	"redis-kvgo/store/kv"
	"strings"
	"sync"
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

var (
	counter int        // shared data
	mu      sync.Mutex // mutual exclusion
)

func main() {
	/*fmt.Println("Gokv project")

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

	if err := s.Set("env", "dev"); err != nil {
		return
	}
	fmt.Println(s)
	copied := s.Clone()
	if err := copied.Set("env", "prod"); err != nil {
		return
	}
	fmt.Println(copied)*/

	//cmds := []Command{
	//{Id: 1, Op: "SET", Key: "env", Value: "prod"},
	//{Id: 2, Op: "SET", Key: "version", Value: "0.0.1"},
	//{Id: 3, Op: "SET", Key: "debug", Value: "true"},
	//{Id: 4, Op: "GET", Key: "env"},
	//{Id: 5, Op: "SET", Key: "region", Value: "south-east-asia"},
	//{Id: 6, Op: "GET", Key: "debug"},
	//{Id: 7, Op: "SET", Key: "content-type", Value: "json"},
	//{Id: 8, Op: "GET", Key: "content-type"},
	//}
	expected := 100
	cmds := make([]Command, expected)
	for i := range expected {
		cmds[i] = Command{Id: i, Op: "INCR", Key: "pageviews"}
	}
	s := kv.NewStore(0)
	CommandOnBoot(s, cmds)

	//fmt.Printf("store size: %d\n", s.Len())
	//fmt.Printf("keys: %v", s.Keys())
	got, _ := s.Get("pageviews")
	fmt.Printf("Expected Incr: %d\n", expected)
	fmt.Printf("got: %s", got)

	// race condition or data race
	raceCondition()
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

func increment() {
	// put the lock where the data/shared data will be read and/or written
	mu.Lock()
	defer mu.Unlock()
	v := counter
	time.Sleep(1 * time.Microsecond)
	counter = v + 1
}

func raceCondition() {
	expected := 1000

	var wg sync.WaitGroup
	for i := 0; i < expected; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment()
		}()
	}

	wg.Wait()

	fmt.Println("\nCounter: ", counter)
	fmt.Println("Expected: ", expected)
}
