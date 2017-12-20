package uuid_test

import (
	"fmt"

	"github.com/northbright/uuid"
)

func ExampleNew() {
	for i := 0; i < 5; i++ {
		uuid, _ := uuid.New()
		fmt.Printf("%v\n", uuid)
	}
	// Output:
}
