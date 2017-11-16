package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"text/template"
	// "log"
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
	latestReload  = time.Now()
	previousValue float64

	statusBC  statusBlockChain
	statusCC  statusCryptoCompare
	statusCMC statusCoinMarketCap
	statusCN  statusCryptoNator

	//Color values
	printMagenta   = color.New(color.FgMagenta)
	printRed       = color.New(color.FgRed)
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

	//Api urls
	urlBC  = "https://blockchain.info/ticker"
	urlCC  = "https://min-api.cryptocompare.com/data/price?fsym=BTC&tsyms=BTC,EUR"
	urlCMC = "https://api.coinmarketcap.com/v1/ticker/bitcoin/?convert=EUR"
	urlCN  = "https://api.cryptonator.com/api/full/btc-eur"

	//Iteration time
	refresTime = time.Second * 15

	//Database
	dbName string = "cryptGo"
)

/*--------------
Database Loading
--------------*/

/*
*Function used to load in the database
 */
func loadDatabase(dbName string) {

	//Establish connection
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//Create Database
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	if err != nil {
		panic(err)
	}

	//Switch to database
	_, err = db.Exec("USE " + dbName)
	if err != nil {
		panic(err)
	}

	//Initiate Table creation
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `" + dbName + "`.`users` ( `id` INT NOT NULL AUTO_INCREMENT , `username` VARCHAR(255) NOT NULL , `email` VARCHAR(255) NOT NULL , `password` VARCHAR(255) NOT NULL , `creation_date` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP , PRIMARY KEY (`id`)) ENGINE = InnoDB;")
	if err != nil {
		panic(err)
	}

}

/*--------------
Crypto Loading
--------------*/

/*
*Function used to load in all the data on launch
 */
func loadCrypto() {

	//Set time ticker
	t := time.NewTicker(refresTime)

	//Start looping
	for {
		loadStatusAllCrypto()
		displayFunction()
		//When t.C receives, the loop will continue
		<-t.C
	}
}

/*
*Load all coin stats
 */
func loadStatusAllCrypto() {
	json.Unmarshal(getBytesByUrl(urlBC), &statusBC)
	json.Unmarshal(getBytesByUrl(urlCC), &statusCC)
	json.Unmarshal(getBytesByUrl(urlCMC), &statusCMC)
	json.Unmarshal(getBytesByUrl(urlCN), &statusCN)
}

/*
*Gets the bytes for a json url
 */
func getBytesByUrl(jsonUrl string) []byte {

	//Get respons form url
	resp, err := http.Get(jsonUrl)
	if err != nil {
		fmt.Println(err)
	}

	//Turn json into bytes
	bytes, _ := ioutil.ReadAll(resp.Body)

	return bytes
}

/*--------------
TERMINAL DISPLAY
--------------*/
/*
*Function used to display in the terminal
 */
func displayFunction() {
	fmt.Println("\n")
	printMagenta.Println("\nGetting new Crypto Status")
	printWhite.Printf("Current time:%v", time.Now())
	displayStatusBlockchain()
	displayStatusCryptocompare()
	displayStatusCoinmarketcap()
	displayStatusCryptonator()
	fmt.Println("\n")
}

func displayStatusBlockchain() {
	printHighBlue.Println("\nBlockchain status: ")
	printHighGreen.Print("€")
	printHighGreen.Printf("%v", statusBC.EUR.Recent)
}

func displayStatusCryptocompare() {
	printHighBlue.Println("\nCryptocompare status: ")
	printHighGreen.Print("€")
	printHighGreen.Printf("%v", statusCC.EUR)
}

func displayStatusCoinmarketcap() {
	printHighBlue.Println("\nCoinmarketcap status: ")
	printHighGreen.Print("€")
	printHighGreen.Printf("%v", statusCMC[0].PriceEur)
}

func displayStatusCryptonator() {
	printHighBlue.Println("\nCryptonator status: ")
	printHighGreen.Print("€")
	printHighGreen.Printf("%v", statusCN.Ticker.Price)
}

/*--------------
WEB DISPLAY
--------------*/

/*
*Index page
 */
func indexHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("gtpl/index.gtpl")

	//Printing the error of the template page
	fmt.Println(t.Execute(w, nil))
}

/*
*login page
 */
func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method

	t, _ := template.ParseFiles("gtpl/login.gtpl")

	//Printing the error of the template page
	fmt.Println(t.Execute(w, nil))
}

/*
*Register page
 */
func registerHandler(w http.ResponseWriter, r *http.Request) {
	//get request method
	fmt.Println("method:", r.Method)
	t, _ := template.ParseFiles("gtpl/register.gtpl")

	//Printing the error of the template page
	fmt.Println(t.Execute(w, nil))
}

/*
*MAIN FUNCTION
 */
func main() {

	go loadDatabase(dbName)
	// go loadCrypto()

	// http loading
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/register", registerHandler)
	http.ListenAndServe(":3000", nil)
}
