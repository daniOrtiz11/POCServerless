package main

import (
	"fmt"
	"log"
	"os"
)

/*
CloseConn is a func to close app
*/
func CloseConn() {
	fmt.Printf(infoColor, "See you soon!")
	fmt.Println()
	fmt.Println()
	os.Exit(0)

}

/*
CurrentState show the current coins in our wallet
*/
func CurrentState() {
	coins := Lambda(arnGet, "")
	if coins != "" {
		fmt.Printf(infoColor, "The wallet have "+coins+" coins")
		fmt.Println()
	}
}

/*
AddCoins add coins to our wallet
*/
func AddCoins(coins string) {
	fmt.Println("Saving " + coins + " coins...")
	resp := Lambda(arnSet, coins)
	if resp == okC {
		AddCoins(coins)
	} else if resp == koC {
		log.Fatal(errorLambda)
	} else {
		fmt.Printf(infoColor, coins+" coins saved in your wallet!")
		fmt.Println()
	}
}
