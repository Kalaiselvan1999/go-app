package greetings

import "fmt"

func Greetings(name string) (string) {
	return fmt.Sprintf("Hello %s, welcome to 1st version of our greetings service", name)
}