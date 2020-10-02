package main
import (
	"alben.com/greetings"
	"log"
	"math/rand"
	"fmt"
	"time"
)

func main() {
	// Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello(randomName())
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(message)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomName() string {
	names := []string{
		"tito",
		"vic",
		"joey",
	}

	return names[rand.Intn(len(names))]
}