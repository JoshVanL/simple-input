package main

import (
	"fmt"
	//simple "github.com/JoshVanL/simple-input"
)

func main() {
	fmt.Print("Running examples..")

	ans1 := Open("A query? OPEN ANSWER", "$ ")
	ans2 := Select("A query? A CHOICE", "» ", []string{"foo", "bla", "bar"})
	ans3 := YesNo("A query? YES NO", "> ", true)

	fmt.Printf("%s\n", ans1)
	fmt.Printf("%s\n", ans2)
	fmt.Printf("%v\n", ans3)
}
