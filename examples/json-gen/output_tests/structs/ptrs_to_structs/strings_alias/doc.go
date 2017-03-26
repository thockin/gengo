// +k8s:json-gen=package
package test

//+k8s:json-gen=false
type Inner struct {
	F1 *string
	F2 *string
	F3 *string
}

//+k8s:json-gen=false
type Tstd struct {
	F *Inner
}

type Ttest Tstd
