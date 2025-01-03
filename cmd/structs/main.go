package main

import (
	"fmt"
	"learning-go/cmd/pkg/models"
)

func main() {
    fmt.Println("Structs in golang")

    beki := models.User{"beki", "asda", true, 12}  // Use structs.User
    fmt.Println(beki.Name)
}