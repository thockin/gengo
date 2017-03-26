// +k8s:json-gen=package
package test

//+k8s:json-gen=false
type Inner struct {
	Inner *Inner
}

//+k8s:json-gen=false
type Tstd struct {
	F Inner
}

type Ttest struct {
	F Inner
}
