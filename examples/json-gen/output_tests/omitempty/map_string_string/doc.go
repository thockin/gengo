// +k8s:json-gen=package
package test

//+k8s:json-gen=false
type Tstd struct {
	F map[string]string `json:"F,omitempty"`
}

type Ttest struct {
	F map[string]string `json:"F,omitempty"`
}
