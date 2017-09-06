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

	se := &Select{
		Query:  "A select no default",
		Prompt: "> ",
		Choice: &[]string{"foo", "bar", "cat", "dog"},
	}
	ans = se.Ask()
	fmt.Printf("Ans >%s\n", ans)

	se = &Select{
		Query:   "A select with default",
		Prompt:  "$ ",
		Choice:  &[]string{"foo", "bar", "cat", "dog"},
		Default: 2,
	}
	ans = se.Ask()
	fmt.Printf("Ans >%s\n", ans)

	yn := &YesNo{
		Query:   "Default Yes",
		Prompt:  "$ ",
		Default: true,
	}
	ansb := yn.Ask()
	fmt.Printf("Ans >%v\n", ansb)

	yn = &YesNo{
		Query:   "Default No",
		Prompt:  "$ ",
		Default: false,
	}
	ansb = yn.Ask()
	fmt.Printf("Ans >%v\n", ansb)
}
