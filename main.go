package main

//kvgo -> key value go project
import "fmt"

func main() {
	fmt.Println("Gokv project")

	store := NewStore(0)
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

	fmt.Println(store.data)
	store.Delete("name")
	fmt.Println(store.data)
	store.Delete("nam")
	fmt.Println(store.data)

	fmt.Println(0x62)
}
