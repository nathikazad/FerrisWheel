package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"log"
	"fmt"
	"math/big"
	"github.com/ethereum/go-ethereum/crypto"
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"math/rand"
	"time"
	"os"
	"bufio"
	"github.com/ethereum/go-ethereum/common"
)


func main() {
	start := time.Now()
	contractInfoFilename := "golang/contractInfo.txt"
	if (len(os.Args) == 2) { //file name is provided
		contractInfoFilename = os.Args[1]
	}
	conn, mainAuth, ferrisToken, ferris, ferrisAddress  := ferrisSetup(contractInfoFilename)

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

		fmt.Printf("New address created %s %d\n", newAuth.From.String(), int(time.Now().Sub(start).Seconds()))

		gasPrice, err := conn.SuggestGasPrice(context.Background())
		if err != nil {
			log.Fatalf("Gas estimation Error %v: ", err)
		}

		rawTx := types.NewTransaction(nonce, newAuth.From, big.NewInt(10000000000000000), 2381623, gasPrice, nil)
		nonce++
		signedTx, err := mainAuth.Signer(types.HomesteadSigner{}, mainAuth.From, rawTx)
		if err != nil {
			log.Fatalf("Sign Error %v: ", err)
		}
		err = conn.SendTransaction(context.Background(), signedTx);
		if err != nil {
			log.Fatalf("Send transaction Error %v: ", err)
		} else {
			fmt.Printf("Ether Transfer requested  %s %d\n", newAuth.From.String(), int(time.Now().Sub(start).Seconds()))
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
			fmt.Printf("Ferris Token Transfer requested  %s %d\n", newAuth.From.String(), int(time.Now().Sub(start).Seconds()))
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
			fmt.Printf("FT Transfer request fulfilled %s %d\n", newAuth.From.String(), int(time.Now().Sub(start).Seconds()))
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
			fmt.Printf("Approval requested  %s %d\n", newAuth.From.String(), int(time.Now().Sub(start).Seconds()))
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
			fmt.Printf("Bid requested %s %d\n", newAuth.From.String(), int(time.Now().Sub(start).Seconds()))
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
			fmt.Printf("Bid request fulfilled %s %d %d\n", newAuth.From.String(), randNumbers[i], int(time.Now().Sub(start).Seconds()))
		}
	}
}


func ferrisSetup(contractInfoFilename string) (*ethclient.Client, *bind.TransactOpts, *FerrisToken, *Ferris, common.Address) {

	conn, err := ethclient.Dial("ws://127.0.0.1:8546")
	if err != nil {
		log.Fatalf("could not create ipc client: %v", err)
	}

	file, err := os.Open(contractInfoFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	scanner.Scan()
	//address of the deployed ferris contract
	ferrisAddress := common.HexToAddress(scanner.Text())

	scanner.Scan()
	scanner.Scan()
	//the private key of the above ferris contract's beneficiary
	privateKeyString := scanner.Text()

	privateKey, _ := crypto.HexToECDSA(privateKeyString)
	auth := bind.NewKeyedTransactor(privateKey)

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
	return conn, auth, ferrisToken, ferris, ferrisAddress
}



