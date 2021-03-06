package tools

import (
	log "github.com/sirupsen/logrus"
)

// FNChecker for fixing "Error return value of '' is not checked (errcheck)"
func FNChecker(fn func() error) {
	if err := fn(); err != nil {
		log.Error(err)
	}
}
