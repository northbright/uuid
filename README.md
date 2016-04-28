# uuid

Package uuid creates pseudo uuid by using crypto/rand.

[![Build Status](https://travis-ci.org/northbright/uuid.svg?branch=master)](https://travis-ci.org/northbright/uuid)

#### Generate Random UUIDs

    for i := 0; i < 5; i++ {
        uuid, _ := uuid.New()
        fmt.Printf("%v\n", uuid)
    }

#### Documentation
* [API Reference](http://godoc.org/github.com/northbright/uuid)

#### License
* [MIT License](./LICENSE)
