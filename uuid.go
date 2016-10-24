// Package uuid creates pseudo uuid by using crypto/rand.
package uuid

import (
	"crypto/rand"
	"fmt"
)

var (
	// DEBUG is debug mode option.
	// Enable / disable debug messages from this package.
	DEBUG = false
)

// New creates pseudo uuid.
func New() (uuid string, err error) {
	b := make([]byte, 16)

	_, err = rand.Read(b)
	if err != nil {
		if DEBUG {
			fmt.Printf("rand.Read() err: %v", err)
		}
		return "", err
	}

	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid, nil
}
