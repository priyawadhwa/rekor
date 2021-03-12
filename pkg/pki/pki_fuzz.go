// +build gofuzz

package pki

import (
	"bytes"
	"fmt"
)

func FuzzNewPublicKey(data []byte) int {
	a := NewArtifactFactory(string(data))
	r := bytes.NewReader([]byte{})
	pk, err := a.NewPublicKey(r)
	if err != nil {
		panic(err)
	}
	fmt.Println("got pk ", pk)
	return 1
}
