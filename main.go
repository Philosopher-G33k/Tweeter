package main

import (
	"fmt"
)

func main() {
	fmt.Println("First from Tweeter")
	tweeter := tweeterManager{}
	tweeter.getcredentials()
	tweeter.register()
}
