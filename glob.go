package main

import ("os";"path/filepath";"strings")

// ExpandGlob handles ** patterns by walking directories recursively.
func ExpandGlob(pattern string) []string {
	if !strings.Contains(pattern, "**") {
		ms, _ := filepath.Glob(pattern)
		return ms
	}
	parts := strings.SplitN(pattern, "**", 2)
	root := parts[0]
	if root == "" { root = "." }
	suffix := strings.TrimPrefix(parts[1], "/")
	var matches []string
	filepath.Walk(root, func(p string, i os.FileInfo, e error) error {
		if e != nil || i.IsDir() { return nil }
		if suffix == "" { matches = append(matches, p); return nil }
		if m, _ := filepath.Match(suffix, filepath.Base(p)); m {
			matches = append(matches, p)
		}
		return nil
	})
	return matches
}
