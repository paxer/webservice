package main

import (
	"fmt"
	"github.com/paxer/webservice/models"
)

func main() {
	u := models.User{
		ID: 2,
		FirstName: "Rob",
		LastName: "McMillan",
	}

	fmt.Println(u)
}