package icfg

import "testing"

func TestRead(t *testing.T) {
	type testData struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
	}
	d := &testData{}
	Load(CfgTypeYaml, "test", "./test/test.yaml", d)
	cache := Get("test").(*testData)

	if cache.Name != "icfg" {
		t.Error("read error")
	}

	t.Logf("%+v", cache)
}
