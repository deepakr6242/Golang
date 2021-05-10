package main

import (
	"fmt"
	"io/ioutil"
	// "log"
)

func main() {
	content, err := ioutil.ReadFile("file.txt")
	if err != nil {
		fmt.Println(err)
	}else{

	fmt.Print(string(content))


// Printf - "Print Formatter" this function allows you to format numbers, variables and strings into the first string parameter you give it
// Print - "Print" This cannot format anything, it simply takes a string and print it
// Println - "Print Line" same thing as Printf() however it will append a newline character\n


}
}