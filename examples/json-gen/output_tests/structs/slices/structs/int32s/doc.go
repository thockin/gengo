// +k8s:json-gen=package
package test

//+k8s:json-gen=false
type Tstd struct {
	F []struct {
		F1 int32
		F2 int32
		F3 int32
	}
}

type Ttest struct {
	F []struct {
		F1 int32
		F2 int32
		F3 int32
	}
}
