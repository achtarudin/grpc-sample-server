package main

import "log"

func main() {
	log.SetFlags(0)
	log.SetOutput(&logWriter{})
}
