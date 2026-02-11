package main
import ("os";"path/filepath";"testing")

func TestLoadConfig(t *testing.T) {
	tmp := t.TempDir()
	os.WriteFile(filepath.Join(tmp,"c.yml"), []byte("watch:\n  - '*.go'\ncommand: go build\nrestart: true\n"), 0644)
	c, e := LoadConfig(filepath.Join(tmp,"c.yml"))
	if e != nil { t.Fatal(e) }
	if len(c.Watch)!=1 { t.Errorf("want 1 watch, got %d",len(c.Watch)) }
	if !c.Restart { t.Error("want restart=true") }
	if c.Debounce!=200 { t.Errorf("want 200, got %d",c.Debounce) }
}

func TestMissing(t *testing.T) {
	_, e := LoadConfig("/no"); if e==nil { t.Error("want error") }
}
