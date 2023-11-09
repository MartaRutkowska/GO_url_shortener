package main

import (
	"errors"
	"fmt"
)

func main() {

	var userUrl string
	fmt.Println("Hello there, provide the url u wish to shorten.")
	fmt.Scanf("%s", &userUrl)

	if urlValid(userUrl) {
		fmt.Println("your url is", userUrl)
	}

}

func urlValid(user_url string) bool {
	if len(user_url) == 0 {
		errors.New("provide non empty url")
		return false
	}
	return true
}
