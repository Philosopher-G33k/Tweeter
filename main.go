package main

import (
	"fmt"
)

func main() {
	fmt.Println("First from Tweeter")
	tweeter := tweeterManager{}
	tweeter.readCredentailsFromJSON()
	fmt.Println("main")
	fmt.Println(tweeter)
	tweeter.register()
}
