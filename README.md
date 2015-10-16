# uuid

Package uuid creates pseudo uuid by using crypto/rand.

#### Generate Random UUIDs

    for i := 0; i < 5; i++ {
        uuid, _ := uuid.New()
        fmt.Printf("%v\n", uuid)
    }

#### Documentation
* [API Reference](http://godoc.org/github.com/northbright/uuid)

#### License
* [MIT License](./LICENSE)
