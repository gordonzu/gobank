package main

import (
	"log"
	"fmt"
)

func main() {
	store, err := NewPostgresStore()

	if err != nil {
		log.Fatal(err)
		fmt.Printf("Error in PostgresStore()")
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
		fmt.Printf("Error in store.init()")
	}

 	server := NewAPIServer(":3000", store)
	server.Run()
 
}
