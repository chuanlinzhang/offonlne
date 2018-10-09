package main
import (
	log "github.com/sirupsen/logrus"
)
func main() {
log.WithFields(log.Fields{
	"animl":"walrus",
}).Info("A walrus appears")
}