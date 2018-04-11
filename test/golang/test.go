package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"strings"
	"fmt"
	"math/big"
	"github.com/ethereum/go-ethereum/crypto"
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"math/rand"
	"time"
)


var ferrisAddress = common.HexToAddress("0x2328ef76C4c55B317573f176b3C751522e7acFD7")
func main() {

	//conn, auth, ferrisToken, ferris, ferrisAddress := ferrisSetup(os.Args[1], os.Args[2])
	conn, mainAuth, ferrisToken, ferris := ferrisSetup()
	var transactions []*types.Transaction
	var auths []*bind.TransactOpts
	nonce, err := conn.PendingNonceAt(context.Background(), mainAuth.From)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	var randNumbers []int64;
	if err != nil {
		log.Fatalf("Nonce Error %v: ", err)
	}
	for i := 0; i < 10; i++ {
		key, _ := crypto.GenerateKey()
		newAuth := bind.NewKeyedTransactor(key)
		auths = append(auths, newAuth)

		fmt.Printf("New address created %s\n", newAuth.From.String())

		gasPrice, err := conn.SuggestGasPrice(context.Background())
		if err != nil {
			log.Fatalf("Gas estimation Error %v: ", err)
		}

		rawTx := types.NewTransaction(nonce, newAuth.From, big.NewInt(1000000000000000), 2381623, gasPrice, nil)
		nonce++
		signedTx, err := mainAuth.Signer(types.HomesteadSigner{}, mainAuth.From, rawTx)
		if err != nil {
			log.Fatalf("Sign Error %v: ", err)
		}
		err = conn.SendTransaction(context.Background(), signedTx);
		if err != nil {
			log.Fatalf("Send transaction Error %v: ", err)
		} else {
			fmt.Printf("Ether Transfer requested  %s\n", newAuth.From.String())
		}

		ftValue := r1.Int63n(9) + 1
		randNumbers = append(randNumbers, ftValue)
		transaction, err := ferrisToken.Transfer(&bind.TransactOpts{
			From:     mainAuth.From,
			Signer:   mainAuth.Signer,
			GasLimit: 2381623,
			Value:    big.NewInt(0),
			Nonce:    big.NewInt(int64(nonce)),
		}, newAuth.From, big.NewInt(ftValue));
		nonce++
		if err != nil {
			log.Fatalf("Ferris Token Transferring error:%s %v: ", newAuth.From.String(), err)
		} else {
			fmt.Printf("Ferris Token Transfer requested  %s\n", newAuth.From.String())
		}
		transactions = append(transactions, transaction)

	}

	for i, newAuth := range(auths) {
		receipt, err := bind.WaitMined(context.Background(), conn, transactions[i])
		if err != nil {
			log.Fatalf("Wait for mining error %s %v: ", newAuth.From.String(), err)
		} else if receipt.Status == types.ReceiptStatusFailed {
			log.Println("FT Transfer failed %s %v: ", newAuth.From.String(), err)
		} else {
			fmt.Printf("FT Transfer request fulfilled %s\n", newAuth.From.String())
		}

		nonce, err := conn.PendingNonceAt(context.Background(), newAuth.From)
		if err != nil {
			log.Fatalf("Nonce Error %v: ", err)
		}

		transaction, err := ferrisToken.Approve(&bind.TransactOpts{
			From:     newAuth.From,
			Signer:   newAuth.Signer,
			GasLimit: 100000,
			Value:    big.NewInt(0),
			Nonce:	  big.NewInt(int64(nonce)),
		}, ferrisAddress, big.NewInt(randNumbers[i]))
		nonce++
		if err != nil {
			log.Fatalf("Ferris Token Approving error:%s %v: ", newAuth.From.String(), err)
		} else {
			fmt.Printf("Approval requested  %s\n", newAuth.From.String())
		}

		transaction, err = ferris.Bid(&bind.TransactOpts{
			From:     newAuth.From,
			Signer:   newAuth.Signer,
			GasLimit: 100000,
			Value:    big.NewInt(0),
			Nonce:	  big.NewInt(int64(nonce)),
		}, big.NewInt(randNumbers[i]))

		if err != nil {
			log.Fatalf("Ferris Token Wait bidding error error:%s %v: ", newAuth.From.String(), err)
		} else {
			fmt.Printf("Bid requested %s\n", newAuth.From.String())
		}
		transactions[i] = transaction
	}



	for i, newAuth := range(auths) {
		receipt, err := bind.WaitMined(context.Background(), conn, transactions[i])
		if err != nil {
			log.Fatalf("Wait for mining error %s %v: ", newAuth.From.String(), err)
		} else if receipt.Status == types.ReceiptStatusFailed {
			log.Println("Bidding failed %s %v: ", newAuth.From.String(), err)
		} else {
			fmt.Printf("Bid request fulfilled %s %d\n", newAuth.From.String(), randNumbers[i])
		}
	}
}


func ferrisSetup() (*ethclient.Client, *bind.TransactOpts, *FerrisToken, *Ferris) {

	key := `{"address":"f332f55eb6a83ab51a25e610efd03074cb3929e0","crypto":{"cipher":"aes-128-ctr","ciphertext":"79e88f8ec2c5555620791bcceb511384f19cd70294fa3d296e9354f9d148b555","cipherparams":{"iv":"54b3faaf645c2f333633cf69f4e7c7ab"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"ce2734d62716a970900e8e837184c537412b1f93c4ce5c39867cf3c387f09cdb"},"mac":"a4ce8d5fa7e2b97ab1fb5020cc88fde39ed993f24dc944c4b48e14070ba300d0"},"id":"cc8e7ef9-62dc-43a0-ba67-7ca17c6a5ad9","version":3}`
	conn, err := ethclient.Dial("/Users/nathik/Library/Ethereum/geth.ipc")
	if err != nil {
		log.Fatalf("could not create ipc client: %v", err)
	}
	auth, err := bind.NewTransactor(strings.NewReader(key), "speakers")

	ferris, err := NewFerris(ferrisAddress, conn)
	address, _ := ferris.GetFerrisTokenAddress(nil)
	ferrisToken, err := NewFerrisToken(address, conn)
	if err != nil {
		log.Fatalf("could not find ferris: %v", err)
	}

	beneficiary, _ := ferris.Beneficiary(nil)
	fmt.Printf("beneficiary: %s \n", beneficiary.String())
	balance, _ := ferrisToken.BalanceOf(nil, beneficiary)
	fmt.Printf("balance: %s \n", balance.String())
	return conn, auth, ferrisToken, ferris
}



