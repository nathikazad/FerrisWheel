package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"os"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"strings"
)

func main() {

	var key string
	var conn *ethclient.Client

	var err error
	var address common.Address
	var contract *Ferris
	switch os.Args[1] {
	case "local":
		key = `{"address":"627306090abab3a6e1400e9345bc60c78a8bef57","crypto":{"cipher":"aes-128-ctr","ciphertext":"c5789188e6009914f45c1d280cc54099e7622e469e59f1e3d4dce83135d57b40","cipherparams":{"iv":"5805aeaa8fa6e167a609c38bdc4e70ae"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"37ebb28452e322aa1976931cfbfda8fb3d3799b2f52e5511ab4a3b595f00aa4d"},"mac":"aa3cb8f0ce2f647d1fdf9cbadf161913804885361e86ffa334f9343b6cda25b1"},"id":"a72cc19d-2f5d-4017-b003-ff42c93bd12c","version":3}`
		conn, err = ethclient.Dial("http://localhost:7545")
		if err != nil {
			log.Fatalf("could not create ipc client: %v", err)
		}
	case "testnet":
		key = `{"address":"f332f55eb6a83ab51a25e610efd03074cb3929e0","crypto":{"cipher":"aes-128-ctr","ciphertext":"79e88f8ec2c5555620791bcceb511384f19cd70294fa3d296e9354f9d148b555","cipherparams":{"iv":"54b3faaf645c2f333633cf69f4e7c7ab"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"ce2734d62716a970900e8e837184c537412b1f93c4ce5c39867cf3c387f09cdb"},"mac":"a4ce8d5fa7e2b97ab1fb5020cc88fde39ed993f24dc944c4b48e14070ba300d0"},"id":"cc8e7ef9-62dc-43a0-ba67-7ca17c6a5ad9","version":3}`
		conn, err = ethclient.Dial("/Users/nathik/Library/Ethereum/geth.ipc")
		if err != nil {
			log.Fatalf("could not create ipc client: %v", err)
		}
	}

	switch os.Args[2] {
	case "new":
		auth, err := bind.NewTransactor(strings.NewReader(key), "speakers")
		if err != nil {
			log.Fatalf("could not create auth: %v", err)
		}
		address, _, contract , err = DeployFerris(auth, conn)
		if err != nil {
			log.Fatalf("could not deploy Ferris contract: %v", err)
		}
		fmt.Printf("address:%s\n" , address.String())
	case "existing":
		contract , err = NewFerris(common.HexToAddress("0xa8dDdf8079a2DC48021E1f67423F8f748e885754"), conn)
		if err != nil {
			log.Fatalf("could not find contract: %v", err)
		}
	}

	beneficiary, _ := contract.Beneficiary(nil)
	fmt.Printf("beneficiary: %s \n", beneficiary.String())

		// ** EVENT START
		//c1 := make(chan *FerrisNewBid)
		//_, err = contract.WatchNewBid(nil,c1);
		//if err != nil {
		//	log.Fatalf("could not watch event: %v", err)
		//}
		//for {
		//	msg := <- c1
		//	fmt.Println(msg.Amount.String())
		//	time.Sleep(time.Second * 1)
		//}
		// ** EVENT END
		//transaction, err := contract.Bet(&bind.TransactOpts{
		//	From:     auth.From,
		//	Signer:   auth.Signer,
		//	GasLimit: 2381623,
		//	Value:    big.NewInt(int64(math.Pow10(17))),
		//}, big.NewInt(1))
		//
		//if err != nil {
		//	log.Fatalf("could not bet: %v", err)
		//}
		//fmt.Printf("total bet: %s\n", transaction.String())
		//
		//total, _ = contract.TotalBet(nil)
		//fmt.Printf("total bet: %s\n", total)
	iter, err := contract.FilterNewBid(nil);
	for iter.Next() {
		fmt.Printf("New Bid:%s %s \n", iter.Event.Bidder.String(), iter.Event.Amount.String());
	}

	c1 := make(chan *FerrisNewBid)
	_, err = contract.WatchNewBid(nil,c1);
	if err != nil {
		log.Fatalf("could not watch event: %v", err)
	}
	for {
		msg := <- c1
		fmt.Printf("New Bid:%s %s \n", msg.Bidder.String(), msg.Amount.String());
	}
}