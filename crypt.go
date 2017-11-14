package main

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"net/http"
	"time"
)

type statusBlockChain struct {
	USD struct {
		Recent float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"USD"`
	AUD struct {
		Recent float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"AUD"`
	BRL struct {
		Recent float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"BRL"`
	CAD struct {
		Recent float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"CAD"`
	CHF struct {
		Recent float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"CHF"`
	CLP struct {
		Recent float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"CLP"`
	CNY struct {
		Recent float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"CNY"`
	DKK struct {
		Recent float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"DKK"`
	EUR struct {
		Recent float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"EUR"`
	GBP struct {
		Recent float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"GBP"`
	HKD struct {
		Recent float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"HKD"`
	INR struct {
		Recent float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"INR"`
	ISK struct {
		Recent float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"ISK"`
	JPY struct {
		Recent float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"JPY"`
	KRW struct {
		Recent float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"KRW"`
	NZD struct {
		Recent float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"NZD"`
	PLN struct {
		Recent float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"PLN"`
	RUB struct {
		Recent float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"RUB"`
	SEK struct {
		Recent float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"SEK"`
	SGD struct {
		Recent float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"SGD"`
	THB struct {
		Recent float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"THB"`
	TWD struct {
		Recent float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"TWD"`
}

type statusCryptoCompare struct {
	BTC int     `json:"BTC"`
	EUR float64 `json:"EUR"`
}

type statusCoinMarketCap []struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Symbol           string `json:"symbol"`
	Rank             string `json:"rank"`
	PriceUsd         string `json:"price_usd"`
	PriceBtc         string `json:"price_btc"`
	Two4HVolumeUsd   string `json:"24h_volume_usd"`
	MarketCapUsd     string `json:"market_cap_usd"`
	AvailableSupply  string `json:"available_supply"`
	TotalSupply      string `json:"total_supply"`
	MaxSupply        string `json:"max_supply"`
	PercentChange1H  string `json:"percent_change_1h"`
	PercentChange24H string `json:"percent_change_24h"`
	PercentChange7D  string `json:"percent_change_7d"`
	LastUpdated      string `json:"last_updated"`
	PriceEur         string `json:"price_eur"`
	Two4HVolumeEur   string `json:"24h_volume_eur"`
	MarketCapEur     string `json:"market_cap_eur"`
}

type statusCryptoNator struct {
	Ticker struct {
		Base    string `json:"base"`
		Target  string `json:"target"`
		Price   string `json:"price"`
		Volume  string `json:"volume"`
		Change  string `json:"change"`
		Markets []struct {
			Market string  `json:"market"`
			Price  string  `json:"price"`
			Volume float64 `json:"volume"`
		} `json:"markets"`
	} `json:"ticker"`
	Timestamp int    `json:"timestamp"`
	Success   bool   `json:"success"`
	Error     string `json:"error"`
}

//Timing duration
type Duration int64

var (
	//Set Timing variables
	latestReload   = time.Now()
	previousValue  float64
	sBC            statusBlockChain
	sCC            statusCryptoCompare
	sCMC           statusCoinMarketCap
	sCN            statusCryptoNator
	printMagenta   = color.New(color.FgMagenta)
	printWhite     = color.New(color.FgWhite)
	printHighBlue  = color.New(color.FgHiBlue)
	printHighGreen = color.New(color.FgHiGreen)
)

//Set Constants (Variables that never change)
const (

	//Time
	Nanosecond  Duration = 1
	Microsecond          = 1000 * Nanosecond
	Millisecond          = 1000 * Microsecond
	Second               = 1000 * Millisecond
	Minute               = 60 * Second
	Hour                 = 60 * Minute

	urlBC  = "https://blockchain.info/ticker"
	urlCC  = "https://min-api.cryptocompare.com/data/price?fsym=BTC&tsyms=BTC,EUR"
	urlCMC = "https://api.coinmarketcap.com/v1/ticker/bitcoin/?convert=EUR"
	urlCN  = "https://api.cryptonator.com/api/full/btc-eur"

	refresTime = time.Second * 1
)

/*
*Function used to load in all the data on launch
 */
func loadFunction() {

	//Loop resetting, break if enough time has passed & reset variable
	for {

		//Get time passed from latest reload
		timePassed := time.Now().Sub(latestReload)

		//Reload struct if refreshtime is passed
		if timePassed > refresTime {
			//Reset Latest reload
			latestReload = time.Now()
			loadAllCryptoStatus()
			displayFunction()
		}
	}
}

/*
*Load all coin stats
 */
func loadAllCryptoStatus() {
	json.Unmarshal(getBytesByUrl(urlBC), &sBC)
	json.Unmarshal(getBytesByUrl(urlCC), &sCC)
	json.Unmarshal(getBytesByUrl(urlCMC), &sCMC)
	json.Unmarshal(getBytesByUrl(urlCN), &sCN)
}

/*
*Gets the bytes for a json url
 */
func getBytesByUrl(jsonUrl string) []byte {
	resp, err := http.Get(jsonUrl)
	if err != nil {
		fmt.Println(err)
	}
	bytes, _ := ioutil.ReadAll(resp.Body)

	return bytes
}

/*--------------
DISPLAY
--------------*/
/*
*Function used to display in the terminal
 */
func displayFunction() {
	fmt.Println("\n")
	printMagenta.Println("\nGetting new Crypto Status")
	printWhite.Printf("Current time:%v", time.Now())
	displayBlockchainStatus()
	displayCryptocompareStatus()
	displayCoinmarketcapStatus()
	displayCryptonatorStatus()
	fmt.Println("\n")
}

func displayBlockchainStatus() {
	printHighBlue.Println("\nBlockchain status: ")
	printHighGreen.Print("€")
	printHighGreen.Printf("%v", sBC.EUR.Recent)
}

func displayCryptocompareStatus() {
	printHighBlue.Println("\nCryptocompare status: ")
	printHighGreen.Print("€")
	printHighGreen.Printf("%v", sCC.EUR)
}

func displayCoinmarketcapStatus() {
	printHighBlue.Println("\nCoinmarketcap status: ")
	printHighGreen.Print("€")
	printHighGreen.Printf("%v", sCMC[0].PriceEur)
}

func displayCryptonatorStatus() {
	printHighBlue.Println("\nCryptonator status: ")
	printHighGreen.Print("€")
	printHighGreen.Printf("%v", sCN.Ticker.Price)
}

/*
*MAIN FUNCTION
 */
func main() {
	loadFunction()
}
