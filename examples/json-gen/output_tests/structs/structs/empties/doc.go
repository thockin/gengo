// +k8s:json-gen=package
package test

//+k8s:json-gen=false
type Tstd struct {
	F1 struct{}
	F2 struct{}
	F3 struct{}
}

type Ttest struct {
	F1 struct{}
	F2 struct{}
	F3 struct{}
}
