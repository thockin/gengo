// +k8s:json-gen=package
package test

//+k8s:json-gen=false
type Tstd struct {
	F bool `json:"F,omitempty"`
}

type Ttest struct {
	F bool `json:"F,omitempty"`
}
