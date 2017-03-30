// +k8s:json-gen=package
package test

//+k8s:json-gen=false
type Elem struct {
	F1 int32
	F2 *int32
	F3 float32
	F4 *float32
	F5 string
	F6 *string
	F7 struct {
		F string
	}
	F8 *struct {
		F string
	}
}

//+k8s:json-gen=false
type Tstd []struct {
	F1 Elem
	F2 Elem
}

type Ttest Tstd
