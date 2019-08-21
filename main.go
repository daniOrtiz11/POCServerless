package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println()
	fmt.Printf(infoColor, "*****************************")
	fmt.Println()
	fmt.Printf(infoColor, "**"+welcomeMessage+"**")
	fmt.Println()
	fmt.Printf(infoColor, "*****************************")
	fmt.Println()
	fmt.Println()
	mainMenu()
}

func mainMenu() {
	stdReader := bufio.NewReader(os.Stdin)
	for {
		showOptions()
		sendData, err := stdReader.ReadString('\n')
		if err != nil {
			fmt.Printf(exceptionReader)
			log.Fatal(err)
		}
		sendData = strings.Replace(sendData, "\n", "", -1)
		option, err := strconv.Atoi(sendData)
		if err == nil {
			switch option {
			case 1:
				CurrentState()
			case 2:
				fmt.Println("Coins to save:")
				stdReader := bufio.NewReader(os.Stdin)
				sendData, err := stdReader.ReadString('\n')
				if err == nil {
					coins := strings.Replace(sendData, "\n", "", -1)
					AddCoins(coins)
				} else {
					log.Fatal(err)
				}
			case 3:
				CloseConn()
			default:
				fmt.Println(badFormatOption)
			}
		} else {
			fmt.Println(badFormatNumber)
		}
		fmt.Println()
	}
}
func showOptions() {
	fmt.Println(optionsTitle)
	fmt.Printf(strongWhiteColor, option1)
	fmt.Println()
	fmt.Printf(strongWhiteColor, option2)
	fmt.Println()
	fmt.Printf(strongWhiteColor, closeOption)
	fmt.Println()
	fmt.Print("> ")
}
