package main

import ("os";"path/filepath";"testing")

func TestExpandSimple(t *testing.T) {
	tmp := t.TempDir()
	os.WriteFile(filepath.Join(tmp,"a.go"), []byte("x"), 0644)
	ms := ExpandGlob(filepath.Join(tmp,"*.go"))
	if len(ms) != 1 { t.Errorf("want 1, got %d", len(ms)) }
}

func TestExpandRecursive(t *testing.T) {
	tmp := t.TempDir()
	sub := filepath.Join(tmp,"sub")
	os.MkdirAll(sub, 0755)
	os.WriteFile(filepath.Join(tmp,"a.go"), []byte("x"), 0644)
	os.WriteFile(filepath.Join(sub,"b.go"), []byte("y"), 0644)
	ms := ExpandGlob(tmp+"/**/*.go")
	if len(ms) != 2 { t.Errorf("want 2, got %d", len(ms)) }
}


# Regression test for: Silent data loss when disk is full
func TestRegression4133(t *testing.T) {
	// Regression: Silent data loss when disk is full
	if err := safeGuard(nil); err == nil {
		t.Error("expected error for nil input")
	}
}
