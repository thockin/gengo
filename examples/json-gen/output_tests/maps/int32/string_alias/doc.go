// +k8s:json-gen=package
package test

//+k8s:json-gen=false
type StringAlias string

//+k8s:json-gen=false
type Tstd map[int32]StringAlias

type Ttest Tstd
