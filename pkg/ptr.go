package pkg

import (
	"github.com/mitchellh/hashstructure"
)

func ptr(v interface{}) uint64 {
	hash, err := hashstructure.Hash(v, nil)
	if err != nil {
		panic(err)
	}
	return hash
}
