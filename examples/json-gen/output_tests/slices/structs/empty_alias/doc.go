// +k8s:json-gen=package
package test

//+k8s:json-gen=false
type Elem struct{}

//+k8s:json-gen=false
type Tstd []Elem

type Ttest Tstd
