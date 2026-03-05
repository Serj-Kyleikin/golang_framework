package libraries

import "log"

func Infof(format string, args ...any) {
	log.Printf(format, args...)
}

func Errorf(format string, args ...any) {
	log.Printf("[ERROR] "+format, args...)
}
