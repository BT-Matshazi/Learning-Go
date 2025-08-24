package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"

	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

// Data represents the data to be passed to the HTML template
type Data struct {
	Title       string
	Description string
	Features    []string
	Examples    []Example
}

// Example represents a code example
type Example struct {
	Title       string
	Description string
	Code        string
}

func main() {
	// Create our data structure
	data := Data{
		Title:       "Comprehensive Guide to Golang",
		Description: "This guide covers the essential concepts of the Go programming language.",
		Features: []string{
			"Simple and concise syntax",
			"Strong static typing with type inference",
			"Memory safety and garbage collection",
			"Built-in concurrency with goroutines and channels",
			"Fast compilation and execution",
			"Rich standard library",
		},
		Examples: []Example{
			{
				Title:       "Hello World",
				Description: "A simple Hello World program in Go",
				Code: `package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
}`,
			},
			{
				Title:       "Goroutines",
				Description: "Using goroutines for concurrent execution",
				Code: `package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}`,
			},
			{
				Title:       "Channels",
				Description: "Communication between goroutines using channels",
				Code: `package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to channel
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	
	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	
	x, y := <-c, <-c // receive from channel
	
	fmt.Println(x, y, x+y)
}`,
			},
		},
	}

	// Parse the HTML template
	tmpl, err := template.ParseFiles("template.html")
	if err != nil {
		log.Fatalf("Error parsing template: %v", err)
	}

	// Execute the template and store the result
	var processed bytes.Buffer
	if err := tmpl.Execute(&processed, data); err != nil {
		log.Fatalf("Error executing template: %v", err)
	}

	// Write the HTML to a temporary file
	htmlFile := "output.html"
	if err := ioutil.WriteFile(htmlFile, processed.Bytes(), 0644); err != nil {
		log.Fatalf("Error writing HTML file: %v", err)
	}
	fmt.Println("HTML file generated successfully:", htmlFile)

	// Convert HTML to PDF
	if err := generatePDF(htmlFile, "golang_guide.pdf"); err != nil {
		log.Fatalf("Error generating PDF: %v", err)
	}
	fmt.Println("PDF file generated successfully: golang_guide.pdf")
}

// generatePDF converts an HTML file to PDF
func generatePDF(htmlFile, pdfFile string) error {
	// Create new PDF generator
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return err
	}

	// Set global options
	pdfg.Dpi.Set(300)
	pdfg.Orientation.Set(wkhtmltopdf.OrientationPortrait)
	pdfg.Grayscale.Set(false)

	// Create page from file
	page := wkhtmltopdf.NewPage(htmlFile)
	page.EnableLocalFileAccess.Set(true)

	// Add page to generator
	pdfg.AddPage(page)

	// Create PDF
	err = pdfg.Create()
	if err != nil {
		return err
	}

	// Write PDF to file
	err = pdfg.WriteFile(pdfFile)
	if err != nil {
		return err
	}

	return nil
}