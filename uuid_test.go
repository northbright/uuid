package uuid_test

import (
	"log"

	"github.com/northbright/uuid"
)

func ExampleNew() {
	for i := 0; i < 5; i++ {
		uuid, _ := uuid.New()
		log.Printf("%v", uuid)
	}
	// Output:
}
