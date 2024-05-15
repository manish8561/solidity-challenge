package contracts

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	"example.com/demo/contracts/biddingwar"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	// biddingwar "./contracts" // for demo
)

const CONTRACT_ADDRESS = "0x5041c3e9d32bdc3091628aced9d13c992384e320"

func DeclareResultByOwner() error {
	client, err := ethclient.Dial("https://nodeapi.test.energi.network/v1/jsonrpc")
	if err != nil {
		return err
	}

	pk := os.Getenv("PRIVATE_KEY")

	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		return err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	address := common.HexToAddress(CONTRACT_ADDRESS)
	instance, err := biddingwar.NewBiddingwar(address, client)
	if err != nil {
		return err
	}

	tx, err := instance.DeclareResult(auth)
	if err != nil {
		return err
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex())

	return nil
}

// BidEvent ..
type BidEvent struct {
	Bid    *big.Int
	Amount *big.Int
	When   *big.Int
	User   common.Address
}

// WinEvent ..
type WinEvent struct {
	Bid        *big.Int
	HighestBid *big.Int
	WinAmount  *big.Int
	When       *big.Int
	Winner     common.Address
}

/**
 * reading logs from blockchain
 */
func ReadLogs() (*[]BidEvent, *[]WinEvent) {
	bidEvents := []BidEvent{}
	winEvents := []WinEvent{}
	client, err := ethclient.Dial("https://nodeapi.test.energi.network/v1/jsonrpc")
	if err != nil {
		log.Fatal(err)
	}

	// 0x Protocol (ZRX) token address
	contractAddress := common.HexToAddress(CONTRACT_ADDRESS)
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(2222430),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(biddingwar.BiddingwarABI)))
	if err != nil {
		log.Fatal(err)
	}

	BidEventSig := []byte("BidEvent(uint256,uint256,uint256,address)")
	WinEventSig := []byte("WinEvent(uint256,uint256,uint256,uint256,address)")
	BidEventSigHash := crypto.Keccak256Hash(BidEventSig)
	WinEventSigHash := crypto.Keccak256Hash(WinEventSig)

	for _, vLog := range logs {
		// fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
		// fmt.Printf("Log Index: %d\n", vLog.Index)

		switch vLog.Topics[0].Hex() {
		case BidEventSigHash.Hex():
			fmt.Printf("Log Name: BidEvent\n")

			var bidEvent BidEvent

			res, err := contractAbi.Unpack("BidEvent", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}

			// fmt.Println("bidevent", res)
			bidEvent.Bid = (vLog.Topics[1]).Big()
			bidEvent.Amount = res[0].(*big.Int)
			bidEvent.When = res[1].(*big.Int)
			bidEvent.User = common.HexToAddress(vLog.Topics[2].Hex())

			fmt.Println(bidEvent)
			bidEvents = append(bidEvents, bidEvent)

		case WinEventSigHash.Hex():
			fmt.Printf("Log Name: WinEvent\n")

			var winEvent WinEvent

			res, err := contractAbi.Unpack("WinEvent", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}

			// fmt.Println("WinEvent", res)
			winEvent.Bid = (vLog.Topics[1]).Big()
			winEvent.HighestBid = res[0].(*big.Int)
			winEvent.WinAmount = res[1].(*big.Int)
			winEvent.When = res[2].(*big.Int)
			winEvent.Winner = common.HexToAddress(vLog.Topics[2].Hex())

			fmt.Println(winEvent)
			winEvents = append(winEvents, winEvent)

		}
	}
	return &bidEvents, &winEvents
}
func CheckGameEndTime() {
	client, err := ethclient.Dial("https://nodeapi.test.energi.network/v1/jsonrpc")
	if err != nil {
		fmt.Println(err)
		return
	}

	address := common.HexToAddress(CONTRACT_ADDRESS)
	instance, err := biddingwar.NewBiddingwar(address, client)
	if err != nil {
		fmt.Println(err)
	}

	gameEndTime, err := instance.GameEndTime(&bind.CallOpts{})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("game end time: %v\n", gameEndTime) // "decimals: 18"

}
