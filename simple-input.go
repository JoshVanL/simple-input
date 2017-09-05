package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

type Open struct {
	Query    string
	Prompt   string
	Required bool
	Default  string
}

type Select struct {
	Query   string
	Prompt  string
	Choice  *[]string
	Default int
}

func (s *Select) Ask() (responce string) {
	go Catch()

	reader := bufio.NewReader(os.Stdin)
	if s.Query != "" {
		fmt.Printf("\n%s", s.Query)
	}
	if s.Default > 0 {
		fmt.Printf(" (default %s)", s.Default)
	}
	fmt.Print("\n")

	for i, s := range *s.Choice {
		fmt.Printf("%d. %s\n", i+1, s)
	}

	acc := false
	var n int
	for !acc {
		fmt.Print(s.Prompt)

		res, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		res = res[0 : len(res)-1]

		if res == "" {
			if s.Default > 0 {
				//TODO: Make this better
				choices = *s.Choice
				return choices[n-1]
			} else {
				fmt.Printf("Nothing entered.\n")
			}

		} else if n, err = strconv.Atoi(res); err != nil || n < 1 || n > len(*s.Choice) {
			fmt.Printf("Responce must be a number between 1 and %d\n", len(*s.Choice))
		} else {
			acc = true
		}
	}

	//TODO: improve this:
	choices := *s.Choice
	responce = choices[n-1]

	return responce
}

func (o *Open) Ask() (responce string) {
	go Catch()

	reader := bufio.NewReader(os.Stdin)
	if o.Query != "" {
		fmt.Printf("\n%s", o.Query)
	}
	if !o.Required {
		fmt.Printf(" (default %s)", o.Default)
	}
	fmt.Print("\n")

	for {
		if o.Prompt != "" {
			fmt.Print(o.Prompt)
		}

		res, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		res = res[0 : len(res)-1]

		if res == "" {
			if o.Required {
				fmt.Print("Nothing entered\n")
			} else if o.Default != "" {
				return responce
			} else {
				return ""
			}

		} else {
			return res
		}
	}
}

func Catch() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	fmt.Printf("\nHandle exit\n")

	os.Exit(1)
}

func YesNo(query, prompt string, defYes bool) (responce bool) {
	go Catch()

	reader := bufio.NewReader(os.Stdin)
	var option string
	if defYes {
		option = " [Y/n]"
	} else {
		option = " [y/N]"
	}

	fmt.Printf("\n%s%s", query, option)

	acc := false
	for !acc {
		fmt.Printf("\n%s", prompt)

		res, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		res = res[0 : len(res)-1]

		res = strings.ToLower(res)
		if res == "y" || res == "yes" {
			responce = true
			acc = true
		} else if res == "n" || res == "no" {
			responce = false
			acc = true
		} else if res == "" {
			responce = defYes
			acc = true
		} else {
			fmt.Printf("Bad responce. %s", option)
		}
	}

	return responce
}
