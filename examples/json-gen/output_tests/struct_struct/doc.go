// +k8s:json-gen=package
package test

type Inner struct {
	F string
}

//+k8s:json-gen=false
type Tstd X
type X struct { //FIXME: remove intermediate T here
	F Inner
}

type Ttest Tstd
