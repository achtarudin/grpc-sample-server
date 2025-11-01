package console

import "log"

func Log(format string, v ...any) {
	log.Printf(format, v...)
}
