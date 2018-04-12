package main

import (
	"log"
	"fmt"
	"sort"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"os"
	"math/big"
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"net/http"
	"github.com/gorilla/websocket"
	"time"
	"strconv"
	"strings"
)

type Bid struct {
	address common.Address ;
	amount int64;
}

//address of the deployed ferris contract
var existingFerrisAddress = common.HexToAddress("0x2328ef76C4c55B317573f176b3C751522e7acFD7")
//the keystore output of the above ferris contract's beneficiary
var key = `{"address":"f332f55eb6a83ab51a25e610efd03074cb3929e0","crypto":{"cipher":"aes-128-ctr","ciphertext":"79e88f8ec2c5555620791bcceb511384f19cd70294fa3d296e9354f9d148b555","cipherparams":{"iv":"54b3faaf645c2f333633cf69f4e7c7ab"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"ce2734d62716a970900e8e837184c537412b1f93c4ce5c39867cf3c387f09cdb"},"mac":"a4ce8d5fa7e2b97ab1fb5020cc88fde39ed993f24dc944c4b48e14070ba300d0"},"id":"cc8e7ef9-62dc-43a0-ba67-7ca17c6a5ad9","version":3}`
//the passphrase used to lock the keystore file
var passphrase = "speakers"

var clients = make(map[*websocket.Conn]bool) // connected clients
var broadcast = make(chan Message)           // broadcast channel

var bidsChannelIn = make(chan uint, 1)      // channel to ask for bids
var bidsChannelOut = make(chan []Bid)       // channel to receive bids
var wakeUpSim = make(chan bool, 1)          // channel to wake up simulator after new bids are processed


// Configure the upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Define our message object
type Message struct {
	Method   string `json:"method"`
	Arg0     string `json:"arg0"`
}

func main() {

	//conn, auth, ferrisToken, ferris, ferrisAddress := ferrisSetup(os.Args[1], os.Args[2])
	conn, auth, _, ferris, _ := ferrisSetup(os.Args[1], os.Args[2])

	//calculate balance for each address and setup listeners to listen for events
	go ferrisEventListeners(ferris)

	//run ferris simulator to wait for new bids, to collect money and spin the wheel through web sockets
	go runFerrisSimulator(conn, auth, ferris)

	// Create a simple file server
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)

	// Configure websocket route
	http.HandleFunc("/ws", handleConnections)

	// Start listening for incoming chat messages
	go handleMessages()

	// Start the server on localhost port 8000 and log any errors
	log.Println("http server started on :8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func ferrisSetup(arg1 string, arg2 string) (*ethclient.Client, *bind.TransactOpts, *FerrisToken, *Ferris, common.Address) {
	var conn *ethclient.Client
	var auth *bind.TransactOpts
	var err error
	var address, ferrisAddress common.Address
	var ferris *Ferris;
	var ferrisToken *FerrisToken;
	switch arg1 {
	case "local":
		key = `{"address":"627306090abab3a6e1400e9345bc60c78a8bef57","crypto":{"cipher":"aes-128-ctr","ciphertext":"c5789188e6009914f45c1d280cc54099e7622e469e59f1e3d4dce83135d57b40","cipherparams":{"iv":"5805aeaa8fa6e167a609c38bdc4e70ae"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"37ebb28452e322aa1976931cfbfda8fb3d3799b2f52e5511ab4a3b595f00aa4d"},"mac":"aa3cb8f0ce2f647d1fdf9cbadf161913804885361e86ffa334f9343b6cda25b1"},"id":"a72cc19d-2f5d-4017-b003-ff42c93bd12c","version":3}`
		conn, err = ethclient.Dial("http://localhost:7545")
		if err != nil {
			log.Fatalf("could not create ipc client: %v", err)
		}
	case "testnet":
		conn, err = ethclient.Dial("/Users/nathik/Library/Ethereum/geth.ipc")
		if err != nil {
			log.Fatalf("could not create ipc client: %v", err)
		}
	}

	auth, err = bind.NewTransactor(strings.NewReader(key), passphrase)

	switch arg2 {
	case "new":
		if err != nil {
			log.Fatalf("could not create auth: %v", err)
		}
		address, _, ferrisToken, err = DeployFerrisToken(auth, conn)
		ferrisAddress, _, ferris, err = DeployFerris(auth, conn, address)
		if err != nil {
			log.Fatalf("could not deploy Ferris ferris: %v", err)
		}
		fmt.Printf("address:%s\n" , ferrisAddress.String())
	case "existing":
		ferris, err = NewFerris(existingFerrisAddress, conn)
		address, _ = ferris.GetFerrisTokenAddress(nil)
		ferrisToken, err = NewFerrisToken(address, conn)
		if err != nil {
			log.Fatalf("could not find ferris: %v", err)
		}
	}
	//beneficiary, _ := ferris.Beneficiary(nil)
	//fmt.Printf("beneficiary: %s \n", beneficiary.String())
	//balance, _ := ferrisToken.BalanceOf(nil, beneficiary)
	//fmt.Printf("balance: %s \n", balance.String())
	return conn, auth, ferrisToken, ferris ,ferrisAddress
}

func ferrisEventListeners(ferris *Ferris){
	bids, lastEventId := calculateBids(ferris)
	for _, bid := range bids {
		fmt.Printf("Bid:%s %d \n", bid.address.String(), bid.amount);
	}

	newBidChannel := make(chan *FerrisNewBid)
	_, err := ferris.WatchNewBid(nil, newBidChannel);
	if err != nil {
		log.Fatalf("could not watch for New Bid event: %v", err)
	}

	acceptedBidChannel := make(chan *FerrisAcceptedBid)
	_, err = ferris.WatchAcceptedBid(nil, acceptedBidChannel);
	if err != nil {
		log.Fatalf("could not watch for accepted Bid event: %v", err)
	}

	withdrewBidChannel := make(chan *FerrisWithdrewBid)
	_, err = ferris.WatchWithdrewBid(nil, withdrewBidChannel);
	if err != nil {
		log.Fatalf("could not watch for withdrew Bid event: %v", err)
	}


	for {
		select {
		case msg := <-newBidChannel: //Watch for new bid events
			if msg.EventId.Uint64() > lastEventId {
				fmt.Printf("New Bid:%s %s \n", msg.Bidder.String(), msg.Amount.String());
				if msg.EventId.Uint64() == (lastEventId + 1) && msg.Amount.Int64() > 0 {
					bids = Sum(bids, Bid{msg.Bidder, msg.Amount.Int64()})
					Sort(bids)
					lastEventId++
				} else {
					bids, lastEventId = calculateBids(ferris)
				}

				select {
				case wakeUpSim <- true:
				default:
					// Simulator is awake already, discard event
				}
			}
		case msg := <-acceptedBidChannel: //Watch for accepted bid events
			if msg.EventId.Uint64() > lastEventId {
				fmt.Printf("Accepted Bid:%s %s \n", msg.Bidder.String(), msg.Amount.String());
				if msg.EventId.Uint64() == (lastEventId + 1) {
					bids = Sum(bids, Bid{msg.Bidder, -msg.Amount.Int64()})
					Sort(bids)
					lastEventId++
				} else {
					bids, lastEventId = calculateBids(ferris)
				}
			}
		case msg := <-withdrewBidChannel: //Watch for withdrew bid events
			if msg.EventId.Uint64() > lastEventId {
				fmt.Printf("Withdrew Bid:%s %s \n", msg.Bidder.String(), msg.Amount.String());
				if msg.EventId.Uint64() == (lastEventId + 1) {
					bids = Sum(bids, Bid{msg.Bidder, -msg.Amount.Int64()})
					Sort(bids)
					lastEventId++
				} else {
					bids, lastEventId = calculateBids(ferris)
				}
			}

		case numOfRequestedBids := <-bidsChannelIn: // Communication channel to ferris simulator to get bids
			if numOfRequestedBids > uint(len(bids)) {
				numOfRequestedBids = uint(len(bids))
			}
			bidsChannelOut <- bids[:numOfRequestedBids]
		}
	}
}

// Go through all the bid events and store accounts that have not expended all their bid money yet
func calculateBids(ferris *Ferris) ([]Bid, uint64) {
	var bids []Bid
	iter1, _ := ferris.FilterNewBid(nil);
	var lastEventId uint64 = 0
	for iter1.Next() {
		bids = Sum(bids, Bid{iter1.Event.Bidder, iter1.Event.Amount.Int64()})
		lastEventId = setIfGreater(iter1.Event.EventId.Uint64(), lastEventId)
	}
	iter2, _ := ferris.FilterAcceptedBid(nil);
	for iter2.Next() {
		bids = Sum(bids, Bid{iter2.Event.Bidder, -iter2.Event.Amount.Int64()})
		lastEventId = setIfGreater(iter2.Event.EventId.Uint64(), lastEventId)
	}
	iter3, _ := ferris.FilterWithdrewBid(nil);
	for iter3.Next() {
		bids = Sum(bids, Bid{iter3.Event.Bidder, -iter3.Event.Amount.Int64()})
		lastEventId = setIfGreater(iter3.Event.EventId.Uint64(), lastEventId)
	}
	bids = Sort(bids)
	return bids, lastEventId;
}

func Sum(bids []Bid, newBid Bid) ([]Bid) {
	if newBid.amount == 0 {
		return bids
	}
	index := -1
	// see if the bid address already exists
	for i , bid := range bids {
		if (newBid.address == bid.address) {
			index = i;
			break;
		}
	}
	if index >= 0 { //bid address already exists
		// Add the new bid's amount to existing amount
		bids[index].amount += newBid.amount
		if bids[index].amount <= 0 { //Remove bid if the amount has totalled up to 0
			bids = append(bids[:index], bids[index+1:]...)
		}
	} else { //bid address doesnt already exist then add bid to bids
		bids = append(bids, newBid)
	}

	return bids;
}

func Sort(transactions []Bid) ([]Bid) {
	sort.Slice(transactions, func(i, j int) bool {
		return transactions[i].amount > transactions[j].amount
	});
	return transactions;
}

func setIfGreater(newEventId uint64, lastEventId uint64)(uint64) {
	if newEventId > lastEventId {
		return newEventId
	} else {
		return lastEventId
	}
}

func runFerrisSimulator(conn *ethclient.Client, auth *bind.TransactOpts, ferris *Ferris) {
	for {
		// Ask event thread for the highest 8 bids
		bidsChannelIn <- 8
		bids := <-bidsChannelOut

		if len(bids) > 0 {
			if len(bids) != 8 { //If the seats arent filled then wait for a few more bids
				fmt.Printf("Only %d bids, waiting for a few more \n", len(bids))
				timeToWait := time.Duration(10 - len(bids))* time.Second //Wait longer for fewer seats
				broadcast <- Message{Method:"Not enough bids", Arg0: timeToWait.String()}
				timer := time.NewTimer(timeToWait) // Wait for longer for few rides
				select {
				case <- wakeUpSim: // Bid arrived, so go and pick it up
					timer.Stop()
					continue
				case <- timer.C: // No bids, so start the ride
					fmt.Println("Timer Expired")
				}
			}
			var transactions []*types.Transaction
			nonce, err := conn.PendingNonceAt(context.Background(), auth.From)
			if err != nil {
				log.Fatalf("Nonce Error %v: ", err)
			}
			// Send requests to accept the chosen bids
			for index, bid := range (bids) {
				transaction, err := ferris.Accept(&bind.TransactOpts{
					From:     auth.From,
					Signer:   auth.Signer,
					GasLimit: 2381623,
					Value:    big.NewInt(0),
					Nonce:    big.NewInt(int64(nonce)),
				}, bid.address, big.NewInt(bid.amount))
				nonce++
				if err != nil {
					log.Fatalf("Accepting Bid Error address:%s %v: ", bid.address.String(), err)
					bids = append(bids[:index], bids[index+1:]...)
				} else {
					transactions = append(transactions, transaction)
					log.Printf("Bid Transaction initiated for address: %s for %d FT\n", bid.address.String(), bid.amount)
				}
			}
			start := time.Now()

			// Wait for the bid accept requests to be fulfilled
			for index, transaction := range (transactions) {
				receipt, err := bind.WaitMined(context.Background(), conn, transaction)
				if err != nil {
					log.Fatalf("Wait for mining error %s %v: ", bids[index].address.String(), err)
				} else if receipt.Status == types.ReceiptStatusFailed {
					log.Println("Accept Bid failed %s %v: ", bids[index].address.String(), err)
					bids = append(bids[:index], bids[index+1:]...)
				} else {
					transaction.ChainId()
				}
				broadcast <- Message{Method:"load", Arg0: strconv.Itoa(index)}
			}
			fmt.Printf("Bids Accepted in %d seconds \n", int(time.Now().Sub(start).Seconds()))

			// Spin the wheel once
			for i := 0; i <= 120; i++ {
				timer := time.NewTimer(time.Millisecond * 50)
				broadcast <- Message{Method: "spin", Arg0: strconv.Itoa(3*i)}
				<- timer.C
			}
		} else { // No more bids left, so just wait
			fmt.Println("Waiting for new bids, sim thread going to sleep")
			<- wakeUpSim
			fmt.Println("New Bid, Sim thread awake")
		}
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()

	// Register our new client
	clients[ws] = true

	for {
		var msg Message
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
	}
}

func handleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast
		// Send it out to every client that is currently connected
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}