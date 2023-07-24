package main

import (
	"encoding/json"
	"fmt"
)

type Dogs struct {
	DogName string `json:"name"`
}

type Person struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	HasDogs   bool
	Dogs      []Dogs `json:"dogs"`
	HasCats   bool
}

func main() {
	var me Person
	me.ID = 1
	me.FirstName = "Jon"
	me.HasDogs = true
	me.Dogs = []Dogs{{DogName: "Rex"}, {DogName: "Peter"}}

	// create a varibable called 'out' and '_' to leave out the error
	// json.MarshalIndet accepts the object, prefix if any so we left it empty, then indent using one tab
	out, _ := json.MarshalIndent(me, "", "\t")

	fmt.Println(string(out))

}
