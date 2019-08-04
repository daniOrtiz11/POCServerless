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
	fmt.Println(welcomeMessage)
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
				AddCoins()
			case 3:
				SubtractCoins()
			case 4:
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
	fmt.Println(option1)
	fmt.Println(option2)
	fmt.Println(option3)
	fmt.Println(closeOption)
	fmt.Print("> ")
}

/*
Sumatorio de valores de objetos JSON
*/
