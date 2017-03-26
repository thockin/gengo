// +k8s:json-gen=package
package test

//+k8s:json-gen=false
type Tstd struct {
	F *string
}

type Ttest Tstd
