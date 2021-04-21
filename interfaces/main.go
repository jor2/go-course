package main

import "fmt"

type Bot interface {
	getGreeting() string
}

type EnglishBot struct {}

type SpanishBot struct {}

func (EnglishBot) getGreeting() string {
	return "Hello!"
}

func (SpanishBot) getGreeting() string {
	return "Hola!"
}

func main() {
	eb := EnglishBot{}
	sb := SpanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b Bot) {
	fmt.Println(b.getGreeting())
}