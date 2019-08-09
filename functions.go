package main

import (
	"fmt"
	"log"
)

/*
CloseConn is a func to close app
*/
func CloseConn() {
	log.Fatal(closeMessage)
}

/*
CurrentState show the current coins in our wallet
*/
func CurrentState() {
	coins := Lambda(arnGet, "")
	if coins != "" {
		fmt.Println("The wallet have:" + coins)
	}
}

/*
AddCoins add coins to our wallet
*/
func AddCoins(coins string) {
	resp := Lambda(arnAdd, coins)
	if resp == okC {
		AddCoins(coins)
	} else if resp == koC {
		log.Fatal(errorLambda)
	} else {
		fmt.Println(coins + "inserted in your wallet!")
	}
}

/*
SubtractCoins subtract coins to our wallet
*/
func SubtractCoins(coins string) {
	resp := Lambda(arnSub, coins)
	if resp == okC {
		SubtractCoins(coins)
	} else if resp == koC {
		log.Fatal(errorLambda)
	} else {
		fmt.Println(coins + "subtracted in your wallet!")
	}
}
