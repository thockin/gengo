package test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/davecgh/go-spew/spew"
	fuzz "github.com/google/gofuzz"
)

func Test_Roundtrip(t *testing.T) {
	fz := fuzz.New().MaxDepth(10).NilChance(0.3)
	for i := 0; i < 1000; i++ {
		var beforeStd Tstd
		fz.Fuzz(&beforeStd)
		beforeTest := Ttest(beforeStd)

		jbStd, err := json.Marshal(beforeStd)
		if err != nil {
			t.Errorf("failed to marshal Tstd: %v", err)
		}
		jbTest, err := json.Marshal(beforeTest)
		if err != nil {
			t.Errorf("failed to marshal Ttest: %v", err)
		}
		if string(jbStd) != string(jbTest) {
			t.Errorf("marshal expected:\n    %s\ngot:\n    %s\nobj:\n    %s",
				indent(jbStd, "    "), indent(jbTest, "    "), dump(beforeTest))
		}

		var afterStd Tstd
		err = json.Unmarshal(jbTest, &afterStd)
		if err != nil {
			t.Errorf("failed to unmarshal to Tstd: %v", err)
		}
		var afterTest Ttest
		err = json.Unmarshal(jbTest, &afterTest)
		if err != nil {
			t.Errorf("failed to unmarshal to Ttest: %v", err)
		}
		if fingerprint(afterStd) != fingerprint(afterTest) {
			t.Errorf("unmarshal expected:\n    %s\ngot:\n    %s\nvia:\n    %s",
				dump(afterStd), dump(afterTest), indent(jbTest, "    "))
		}
	}
}

const indentStr = ">  "

func fingerprint(obj interface{}) string {
	c := spew.ConfigState{SortKeys: true}
	return c.Sprintf("%v", obj)
}

func dump(obj interface{}) string {
	cfg := spew.ConfigState{
		Indent: indentStr,
	}
	return cfg.Sdump(obj)
}

func indent(src []byte, prefix string) string {
	var buf bytes.Buffer
	json.Indent(&buf, src, prefix, indentStr)
	return buf.String()
}
