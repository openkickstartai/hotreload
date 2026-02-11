package main

import "testing"

func TestContainerName(t *testing.T) {
	d := NewDockerWatcher("myapp", "8080:8080")
	if d.containerName() != "hotreload-myapp" {
		t.Errorf("want hotreload-myapp, got %s", d.containerName())
	}
}

func TestNewDockerWatcher(t *testing.T) {
	d := NewDockerWatcher("test", "3000:3000")
	if d.image != "test" { t.Error("wrong image") }
	if d.port != "3000:3000" { t.Error("wrong port") }
}
