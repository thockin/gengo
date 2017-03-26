// +k8s:json-gen=package
package test

//+k8s:json-gen=false
type Inner struct{}

//+k8s:json-gen=false
type Tstd struct {
	F Inner
}

type Ttest Tstd
