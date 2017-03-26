// +k8s:json-gen=package
package test

//+k8s:json-gen=false
type Tstd struct {
	F1 *float32
	F2 *float32
	F3 *float32
}

type Ttest Tstd
