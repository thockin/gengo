// +k8s:json-gen=package
package test

//+k8s:json-gen=false
type StringAlias string

//+k8s:json-gen=false
type Tstd map[StringAlias]uint8

type Ttest Tstd