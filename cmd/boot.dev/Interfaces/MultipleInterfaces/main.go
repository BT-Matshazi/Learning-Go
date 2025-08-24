package main

import (
	"fmt"
)

// Multiple Interfaces
// A type can implement any number of interfaces in Go. For example, the empty interface, interface{}, is always implemented by every type because it has no requirements.

// Assignment
// Complete the required methods so that the email type implements both the expense and formatter interfaces.

// Complete the cost() method:

// If the email is not "subscribed", then the cost is 5 cents for each character in the body.
// If it is, then the cost is 2 cents per character.
// Return the total cost of the entire email in cents.
// Complete the format() method.

// It should return a string in this format:
// 'CONTENT' | Subscribed

// If the email is not subscribed, change the second part to "Not Subscribed":
// 'CONTENT' | Not Subscribed

// The single quotes are included in the string, and CONTENT is the email's body. For example:

// 'Hello, World!' | Subscribed

func (e email) cost() int {
	//to get the length of the email
	length := len(e.body)

	//check is subbed 
	if e.isSubscribed {
		return length * 2
	}
	//if not subbed
	return length * 5
}

// format() returns the email content in the specified format
func (e email) format() string {
    // Determine subscription status text
    status := "Not Subscribed"
    if e.isSubscribed {
        status = "Subscribed"
    }
    // Return formatted string with single quotes around body
    return fmt.Sprintf("'%s' | %s", e.body, status)
}

type expense interface {
	cost() int
}

type formatter interface {
	format() string
}

type email struct {
	isSubscribed bool
	body         string
}
