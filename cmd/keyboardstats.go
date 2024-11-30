package main

import (
	"flag"
	"log"

	"github.com/kwarunek/keyboardstats/internal/listener"
	"github.com/kwarunek/keyboardstats/internal/store"
)

func main() {
	dev := flag.String("dev", "/dev/input/event3", "device to listen to")
	verbose := flag.Bool("v", false, "verbose output")
	dumpInterval := flag.Int("dump-interval", 600, "dump interval in seconds")
	dbPath := flag.String("db", "keyboardstats.db", "path to the database")
	flag.Parse()

	store, err := store.NewStore(*dbPath, *dumpInterval)
	if err != nil {
	    log.Fatal(err)
	}
	defer store.Close()

	listener := listener.NewListener(*dev, store, *verbose)
	listener.Listen()
}
