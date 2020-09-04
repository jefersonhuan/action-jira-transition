package main

import (
	"fmt"
	"main/jira"
)

func main() {
	fmt.Println("Initialing Jira status transition")

	err := jira.StartTransition()
	if err != nil {
		panic(err)
	}
}
