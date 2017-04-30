// +k8s:json-gen=package
package test

//+k8s:json-gen=false
type Embedded struct {
	F1 int32
	F2 string
}

//+k8s:json-gen=false
type Tstd struct {
	F1 int32
	Embedded
}

type Ttest struct {
	F1 int32
	Embedded
}
