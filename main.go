package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"github.com/mdh67899/openfalcon-flume-monitor/process"
)

func process_signal(pid int, fn func()) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	log.Println(pid, "register signal notify")
	for {
		s := <-sigs
		log.Println("recv", s)

		switch s {
		case syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			log.Println("gracefull shut down")
			fn()
			log.Println(pid, "exit")
			os.Exit(0)

		default:
			log.Println("couldn't process this signal:", s)
		}
	}
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	cfg := flag.String("c", "cfg.json", "configuration file")
	flag.Parse()

	service := process.NewProgram(*cfg)
	service.Process()

	process_signal(os.Getpid(), service.Stop)
}
