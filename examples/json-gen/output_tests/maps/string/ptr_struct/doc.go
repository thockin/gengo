// +k8s:json-gen=package
package test

//+k8s:json-gen=false
type Struct struct {
	String string
	Int    int32
	Float  float64
	Struct struct {
		X string
	}
	Slice []string
	Map   map[string]string
}

//+k8s:json-gen=false
type Tstd map[string]*Struct

type Ttest Tstd
