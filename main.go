package main

import (
	"fmt"
)

func main() {
	fmt.Print("Running examples..")
	op := &Open{
		Query:    "A query required",
		Prompt:   "$ ",
		Required: true,
	}
	ans := op.Ask()
	fmt.Printf("Result >%s\n", ans)

	op = &Open{
		Query:    "A query not required",
		Prompt:   "$ ",
		Required: false,
		Default:  "default",
	}
	ans = op.Ask()
	fmt.Printf("Result >%s\n", ans)

	//	se := &Select{
	//		Query: "A select no def",
	//

	//ans1 := Open("A query? OPEN ANSWER", "$ ")
	//ans2 := Select("A query? A CHOICE", "» ", []string{"foo", "bla", "bar"})
	//ans3 := YesNo("A query? YES NO", "> ", true)
	//ans4 := NoQuery("No Query» ")

	//fmt.Printf("%s\n", ans1)
	//fmt.Printf("%s\n", ans2)
	//fmt.Printf("%v\n", ans3)
	//fmt.Printf("%v\n", ans4)
}
