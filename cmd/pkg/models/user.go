package models

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

type Car struct {
	Brand   string
	Model   string
	Doors   int
	Mileage int
}

type messageToSend struct {
	phoneNumber string
	messages    string
}