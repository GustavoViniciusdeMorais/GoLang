package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("GoLangGo!")

	states := make(map[string]string)
	fmt.Println(states)

	states["GO"] = "Goias"
	states["MG"] = "Minas Gerais"
	states["CE"] = "Ceara"
	fmt.Println(states)

	fmt.Printf("Get state of %s\n", states["MG"])

	delete(states, "GO")
	fmt.Println(states)

	// first foreach loop
	for k, v := range states {
		fmt.Printf("%s=>%s\n", k, v)
	}

	// order the valuesby key
	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Println("\nOrdered states:")
	for i := range keys {
		fmt.Println(states[keys[i]])
	}
}
