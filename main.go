package main

import ("flag";"fmt";"log";"os")

var version = "0.1.0"

func main() {
	w := flag.String("watch","","glob")
	run := flag.String("run","","command")
	restart := flag.Bool("restart",false,"restart on change")
	flag.Parse()
	cfg := Config{Debounce:200,Restart:*restart}
	if *w != "" && *run != "" {
		cfg.Watch = []string{*w}; cfg.Command = *run
	} else {
		var err error
		cfg, err = LoadConfig(".hotreload.yml")
		if err != nil { fmt.Println("No config found"); os.Exit(1) }
	}
	log.Printf("hotreload %s: watching %v\n", version, cfg.Watch)
	NewWatcher(cfg).Start()
}
