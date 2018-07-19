package main

import (
	"fmt"
	"log"

	"github.com/petuhovskiy/bitgo"
)

func main() {
	var accessToken string
	fmt.Print("Enter access token: ")
	fmt.Scanf("%s\n", &accessToken)
	session := bitgo.NewSession(accessToken)

	info, err := session.GetSessionInfo()
	fmt.Printf("SessionInfo = %#v\n", info)
	fmt.Printf("Error = %s\n", err)

	btc := session.GetCoin("tbtc")
	wallets, err := btc.GetWalletsPage(0, "", false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Wallets:")
	for _, w := range wallets.Wallets {
		fmt.Printf("%#v\n", w)
	}
}
