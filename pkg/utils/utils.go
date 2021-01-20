package utils

import (
	log "github.com/sirupsen/logrus"
)

// GetHello returns hello message
func GetHello() string {
	log.Debug("Message from package")
	return "hello from package function"
}
