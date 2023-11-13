package main

import (
	"fmt"
	"net/url"
	"slices"
)

func main() {

	var userUrl = displayOptions()

	if userUrl != "" {
		fmt.Printf("userUrl: %s", userUrl)
	}

	return
}

func displayOptions() string {

	var userUrl string
	var userOption int
	options := []int{1, 2, 3}

	fmt.Println("Hello, please choose one of the options:")
	fmt.Println("1 - shorten the url")
	fmt.Println("2 - return the full url from shortened form")
	fmt.Println("3 - escape program")

	_, err := fmt.Scanf("%d \n", &userOption)
	if err != nil || !slices.Contains(options, userOption) {
		fmt.Println("No valid option chosen, escaping")
	}

	switch userOption {
	case 1:
		userUrl = readUserUrl()
	case 2:
		userUrl = readUserUrl()
	default:
	}

	return userUrl
}

func readUserUrl() string {

	var userUrl string
	fmt.Println("Provide the url u wish to shorten/get full version of.")

	for {
		_, err := fmt.Scanf("%s \n", &userUrl)
		if err == nil && urlValid(userUrl) {
			break
		} else if err != nil {
			fmt.Println("U need to provide one argument")
		} else if err == nil && !urlValid(userUrl) {
			fmt.Println("your url is invalid! provide a valid one")
		}
	}

	return userUrl
}

func urlValid(userUrl string) bool {
	_, err := url.ParseRequestURI(userUrl)
	return err == nil
}
