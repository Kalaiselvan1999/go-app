package greetings

import "fmt"

func Greetings(name string) (string) {
	return fmt.Sprintf("Hello %s, welcome to 2nd version of our greetings service", name)
}