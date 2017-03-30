// +k8s:json-gen=package
package test

import other "k8s.io/gengo/examples/json-gen/output_tests/cross_pkg/string/a"

//+k8s:json-gen=false
type Tstd other.T

type Ttest Tstd
