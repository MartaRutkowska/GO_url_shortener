package main

import (
	"fmt"
	"net/url"
)

func main() {

	var userUrl string
	fmt.Println("Hello there, provide the url u wish to shorten.")

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

	return
}

func urlValid(userUrl string) bool {
	_, err := url.ParseRequestURI(userUrl)
	return err == nil
}
