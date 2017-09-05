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

func Catch() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	fmt.Printf("\nHandle exit\n")

	os.Exit(1)
}

func YesNo(query, promt string, defYes bool) (responce bool) {
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
		fmt.Printf("\n%s", promt)

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

func Select(query, promt string, choice []string) (responce string) {
	go Catch()

	reader := bufio.NewReader(os.Stdin)
	var n int
	acc := false
	fmt.Printf("\n%s\n", query)
	for i, s := range choice {
		fmt.Printf("%d. %s\n", i+1, s)
	}

	for !acc {
		fmt.Printf("%s", promt)

		responce, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		responce = responce[0 : len(responce)-1]

		n, err = strconv.Atoi(responce)
		if err != nil || n < 1 || n > len(choice) {
			fmt.Printf("Responce must be a number between 1 and %d\n", len(choice))
		} else {
			acc = true
		}
	}

	responce = choice[n-1]

	return responce
}

func Open(query, promt string) (responce string) {
	go Catch()

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("\n%s\n", query)
	acc := false

	for !acc {
		fmt.Print(promt)

		res, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		res = res[0 : len(res)-1]

		if res == "" {
			fmt.Print("Nothing entered\n")
		} else {
			responce = res
			acc = true
		}
	}

	return responce
}

func NoQuery(promt string) (responce string) {
	go Catch()

	reader := bufio.NewReader(os.Stdin)
	acc := false

	for !acc {
		fmt.Print(promt)

		res, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		res = res[0 : len(res)-1]

		if res == "" {
			fmt.Print("Nothing entered\n")
		} else {
			responce = res
			acc = true
		}
	}

	return responce
}
