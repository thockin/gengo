// +k8s:json-gen=package
package test

import (
	"encoding"
	"strings"
)

//+k8s:json-gen=false
type KeyType string

func (k KeyType) MarshalText() ([]byte, error) {
	return []byte("MANUAL__" + k), nil
}

func (k *KeyType) UnmarshalText(text []byte) error {
	*k = KeyType(strings.TrimPrefix(string(text), "MANUAL__"))
	return nil
}

var _ encoding.TextMarshaler = KeyType("")
var _ encoding.TextUnmarshaler = new(KeyType)

//+k8s:json-gen=false
type Map map[KeyType]string

//+k8s:json-gen=false
type Tstd Map

type Ttest Tstd
