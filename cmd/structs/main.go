package main

import (
	"fmt"
	"learning-go/cmd/pkg/models"
)

func main() {
    fmt.Println("Structs in golang")

    beki := models.User{"beki", "asda", true, 12}

    fmt.Println(beki)
}