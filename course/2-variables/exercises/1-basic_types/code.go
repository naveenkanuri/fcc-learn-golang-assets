package main

import "fmt"

func main() {
	// initialize variables here
	var smsSendingLimit int = 32
	var costPerSMS float32 = 2.3
	var hasPermission bool = true
	var username string = "user"

	fmt.Printf("%v %f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)
}
