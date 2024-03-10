package main

import "log"

func myLog(format string, args ...any) {
	log.Printf("[my] "+format, args...)
}
