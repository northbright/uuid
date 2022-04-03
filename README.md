# uuid

uuid is a [Golang](http://golang.org) package which creates pseudo uuid by using crypto/rand.

## Generate Random UUIDs

    for i := 0; i < 5; i++ {
        uuid, _ := uuid.New()
        fmt.Printf("%v\n", uuid)
    }

## Documentation
* [API Reference](http://pkg.go.dev/github.com/northbright/uuid)

## License
* [MIT License](./LICENSE)
