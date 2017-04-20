// +k8s:json-gen=package
package test

import (
	"encoding"
	"strings"
)

//+k8s:json-gen=false
type KeyType struct {
	X string
}

func (k KeyType) MarshalText() ([]byte, error) {
	return []byte("MANUAL__" + k.X), nil
}

func (k *KeyType) UnmarshalText(text []byte) error {
	k.X = strings.TrimPrefix(string(text), "MANUAL__")
	return nil
}

var _ encoding.TextMarshaler = KeyType{}
var _ encoding.TextUnmarshaler = &KeyType{}

//+k8s:json-gen=false
type Map map[KeyType]string

//+k8s:json-gen=false
type Tstd Map

type Ttest Tstd
