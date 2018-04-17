package main

import (
	"log"
	"fmt"
	"sort"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"context"
	"github.com/ethereum/go-ethereum/core/types"
	"net/http"
	"github.com/gorilla/websocket"
	"time"
	"strconv"
	"os"
	"github.com/ethereum/go-ethereum/crypto"
	"bufio"
)

type Bid struct {
	address common.Address ;
	amount int64;
}


var clients = make(map[*websocket.Conn]bool) // connected clients
var broadcast = make(chan Message)           // broadcast channel

var bidsChannelIn = make(chan uint, 1)      // channel to ask for bids
var bidsChannelOut = make(chan []Bid)       // channel to receive bids
var wakeUpSim = make(chan bool, 1)          // channel to wake up simulator after new bids are processed
var cont context.Context

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
	Arg1     string `json:"arg1"`
}

func main() {

	contractInfoFilename := "golang/contractInfo.txt"
	if (len(os.Args) == 2) { //file name is provided
		contractInfoFilename = os.Args[1]
	}

	conn, auth, _, ferris  := ferrisSetup(contractInfoFilename)
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


func ferrisSetup(contractInfoFilename string) (*ethclient.Client, *bind.TransactOpts, *FerrisToken, *Ferris) {

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
	existingFerrisAddress := scanner.Text()

	scanner.Scan()
	scanner.Scan()
	//the private key of the above ferris contract's beneficiary
	privateKeyString := scanner.Text()

	privateKey, _ := crypto.HexToECDSA(privateKeyString)
	auth := bind.NewKeyedTransactor(privateKey)

	ferris, err := NewFerris(common.HexToAddress(existingFerrisAddress), conn)
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

func ferrisEventListeners(ferris *Ferris){
	bids, lastEventId := calculateBids(ferris)
	sendTopEightToClient(bids)
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
			sendTopEightToClient(bids)
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
			//sendTopEightToClient(bids)
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
			sendTopEightToClient(bids)
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
	iter1, err := ferris.FilterNewBid(nil);
	if err != nil {
		log.Fatalf("could not find New Bid iterator: %v", err)
	}
	var lastEventId uint64 = 0
	for iter1.Next() {
		bids = Sum(bids, Bid{iter1.Event.Bidder, iter1.Event.Amount.Int64()})
		lastEventId = setIfGreater(iter1.Event.EventId.Uint64(), lastEventId)
	}
	iter2, err := ferris.FilterAcceptedBid(nil);
	if err != nil {
		log.Fatalf("could not find Accept Bid iterator: %v", err)
	}
	for iter2.Next() {
		bids = Sum(bids, Bid{iter2.Event.Bidder, -iter2.Event.Amount.Int64()})
		lastEventId = setIfGreater(iter2.Event.EventId.Uint64(), lastEventId)
	}
	iter3, err := ferris.FilterWithdrewBid(nil);
	if err != nil {
		log.Fatalf("could not find Withdrew Bid iterator: %v", err)
	}
	for iter3.Next() {
		bids = Sum(bids, Bid{iter3.Event.Bidder, -iter3.Event.Amount.Int64()})
		lastEventId = setIfGreater(iter3.Event.EventId.Uint64(), lastEventId)
	}
	bids = Sort(bids)
	return bids, lastEventId;
}

func sendTopEightToClient(bids []Bid) {
	broadcast <- Message{Method:"clear"}
	numOfRequestedBids := 8
	if numOfRequestedBids > len(bids) {
		numOfRequestedBids = len(bids)
	}
	for _, bid := range bids[:numOfRequestedBids] {
		broadcast <- Message{Method:"ready", Arg0: bid.address.String(), Arg1:strconv.Itoa(int(bid.amount))}
	}
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
					fmt.Println("Ok screw them")
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
					log.Printf("Accept Bid Receipt status is failed %s \n", bids[index].address.String())
					bids = append(bids[:index], bids[index+1:]...)
				} else {
					broadcast <- Message{Method:"load", Arg0: bids[index].address.String(), Arg1:strconv.Itoa(int(bids[index].amount))}
				}
			}
			fmt.Printf("Bids Processed in %d seconds \n", int(time.Now().Sub(start).Seconds()))


			if len(bids) > 0 {
				//spin the wheel thrice
				for i := 0; i <= 360; i++ {
					timer := time.NewTimer(time.Millisecond * 50)
					broadcast <- Message{Method: "spin", Arg0: strconv.Itoa(3 * i)}
					<-timer.C
				}
				timer := time.NewTimer(time.Millisecond * 500)
				<-timer.C
				broadcast <- Message{Method: "clear"}
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