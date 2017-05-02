// +k8s:json-gen=package
package test

//+k8s:json-gen=false
type Tstd struct {
	F uint32 `json:"F,string"`
}

type Ttest struct {
	F uint32 `json:"F,string"`
}
