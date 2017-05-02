// +k8s:json-gen=package
package test

//+k8s:json-gen=false
type Tstd struct {
	F int32 `json:"F,string"`
}

type Ttest struct {
	F int32 `json:"F,string"`
}
