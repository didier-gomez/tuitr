package main
import (
	"log"
	"github.com/didier-gomez/tuitr/handlers"
	"github.com/didier-gomez/tuitr/bd"
)

func main() {
	if bd.CheckConnection() == 0 {
		log.Fatal("sin conexión a bd")
		return
	}
	handlers.Handlers()
}