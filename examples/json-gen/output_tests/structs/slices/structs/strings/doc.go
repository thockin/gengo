// +k8s:json-gen=package
package test

//+k8s:json-gen=false
type Tstd struct {
	F []struct {
		F1 string
		F2 string
		F3 string
	}
}

type Ttest struct {
	F []struct {
		F1 string
		F2 string
		F3 string
	}
}
