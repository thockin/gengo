// +k8s:json-gen=package
package test

type Inner struct {
	F string
}

//+k8s:json-gen=false
type Tstd []struct {
	F Inner
}

type Ttest Tstd
