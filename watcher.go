package main

import ("log";"os";"os/exec";"path/filepath";"strings";"time")

type Watcher struct { cfg Config; cmd *exec.Cmd; last map[string]time.Time }

func NewWatcher(c Config) *Watcher { return &Watcher{cfg:c, last:map[string]time.Time{}} }

func (w *Watcher) Start() {
	w.run()
	tk := time.NewTicker(time.Duration(w.cfg.Debounce)*time.Millisecond)
	for range tk.C {
		if w.changed() {
			log.Println("[hotreload] change detected")
			if w.cfg.Restart && w.cmd != nil && w.cmd.Process != nil { w.cmd.Process.Kill(); w.cmd.Wait() }
			w.run()
		}
	}
}

func (w *Watcher) run() {
	p := strings.Fields(w.cfg.Command); if len(p)==0 { return }
	w.cmd = exec.Command(p[0],p[1:]...); w.cmd.Stdout = os.Stdout; w.cmd.Stderr = os.Stderr
	if w.cfg.Restart { w.cmd.Start() } else { w.cmd.Run() }
}

func (w *Watcher) changed() bool {
	ch := false
	for _, pat := range w.cfg.Watch {
		ms, _ := filepath.Glob(pat)
		for _, m := range ms {
			i, e := os.Stat(m); if e != nil { continue }
			if l, ok := w.last[m]; ok && i.ModTime().After(l) { ch = true }
			w.last[m] = i.ModTime()
		}
	}
	return ch
}
