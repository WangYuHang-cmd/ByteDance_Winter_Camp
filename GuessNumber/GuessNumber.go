package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	maxNum := 1000
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	//fmt.Println("The secret number is ", secretNumber)

	fmt.Println("Please enter your guess")
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occured while reading unput. Please try again", err)
			//return
			continue
		}
		input = strings.TrimSuffix(input, "\n")
		input = strings.TrimSuffix(input, "\r") //防止从终端输入时有\r

		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid Input. Please enter an integer value", err)
			//return
			continue
		}
		fmt.Println("your guess is", guess)

		if guess > secretNumber {
			fmt.Println("Your guess is too big")
		} else if guess < secretNumber {
			fmt.Println("Your guess is too small")
		} else {
			fmt.Println("Correct, you Legend!")
			break
		}
	}
}
