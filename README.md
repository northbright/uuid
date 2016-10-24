# uuid

[![Build Status](https://travis-ci.org/northbright/uuid.svg?branch=master)](https://travis-ci.org/northbright/uuid)
[![Go Report Card](https://goreportcard.com/badge/github.com/northbright/uuid)](https://goreportcard.com/report/github.com/northbright/uuid)

uuid is a [Golang](http://golang.org) package which creates pseudo uuid by using crypto/rand.

#### Generate Random UUIDs

    for i := 0; i < 5; i++ {
        uuid, _ := uuid.New()
        fmt.Printf("%v\n", uuid)
    }

#### Documentation
* [API Reference](http://godoc.org/github.com/northbright/uuid)

#### License
* [MIT License](./LICENSE)
