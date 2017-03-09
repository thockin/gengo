package test

import (
	"encoding/json"
	"testing"

	"github.com/davecgh/go-spew/spew"
	fuzz "github.com/google/gofuzz"
)

func Test_Roundtrip(t *testing.T) {
	fz := fuzz.New()
	for i := 0; i < 100; i++ {
		var beforeStd Tstd
		fz.Fuzz(&beforeStd)
		beforeTest := Ttest(beforeStd)

		jbStd, err := json.Marshal(beforeStd)
		if err != nil {
			t.Errorf("failed to marshal: %v", err)
		}
		jbTest, err := Marshal_ptr_int32_Ttest(beforeTest)
		if err != nil {
			t.Errorf("failed to marshal: %v", err)
		}
		if string(jbStd) != string(jbTest) {
			t.Errorf("marshal expected %q, got %q", string(jbStd), string(jbTest))
		}

		var afterStd Ttest
		err = json.Unmarshal(jbTest, &afterStd)
		if err != nil {
			t.Errorf("failed to unmarshal: %v", err)
		}
		var afterTest Ttest
		err = Unmarshal_ptr_int32_Ttest(jbTest, &afterTest)
		if err != nil {
			t.Errorf("failed to unmarshal: %v", err)
		}
		if dump(afterStd) != dump(afterTest) {
			t.Errorf("expected %v, got %v via %q", dump(afterStd), dump(afterTest), string(jbTest))
		}
		if dump(beforeTest) != dump(afterTest) {
			t.Errorf("expected %v, got %v via %q", dump(beforeTest), dump(afterTest), string(jbTest))
		}
	}
}

func dump(obj interface{}) string {
	return spew.Sprintf("%#v", obj)
}
