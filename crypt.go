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

type statusCoinMarketCapBTC []struct {
	LastUpdated string `json:"last_updated"`
	PriceEur    string `json:"price_eur"`
}

type statusCoinMarketCapLTC []struct {
	LastUpdated string `json:"last_updated"`
	PriceEur    string `json:"price_eur"`
}

type userCryptoRecord struct {
	Id           int
	Amount       float64
	AddOrDeduct  string
	CreationDate string
}

type pageVariables struct {
	Title     string
	Variables map[string]map[int]userCryptoRecord
}

//Timing duration
type Duration int64

var (
	//Set Timing variables
	latestReload  = time.Now()
	previousValue float64

	statusCMCBTC statusCoinMarketCapBTC
	statusCMCLTC statusCoinMarketCapLTC

	//Color values
	printMagenta   = color.New(color.FgMagenta)
	printRed       = color.New(color.FgRed)
	printWhite     = color.New(color.FgWhite)
	printHighBlue  = color.New(color.FgHiBlue)
	printHighGreen = color.New(color.FgHiGreen)

	userCryptoRecords = map[string]map[int]userCryptoRecord{}
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

	//Urls
	urlCMCBTC = "https://api.coinmarketcap.com/v1/ticker/bitcoin/?convert=EUR"
	urlCMCLTC = "https://api.coinmarketcap.com/v1/ticker/litecoin/?convert=EUR"

	//Iteration time
	refresTime = time.Second * 180

	//Database
	dbName string = "cryptGo"
)

/*--------------
Helper functions
--------------*/

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

/*--------------
Database Loading
--------------*/

/*
*Function used to load in the database
 */
func loadDatabase(dbName string) {

	//Establish connection
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/")
	checkErr(err)
	defer db.Close()

	//Create Database
	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS " + dbName)
	checkErr(err)

	//Switch to database
	_, err = db.Exec("USE " + dbName)
	checkErr(err)

	//Initiate Table creation

	//User Table
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `" + dbName + "`.`users` ( `id` INT NOT NULL AUTO_INCREMENT , `username` VARCHAR(255) NOT NULL , `email` VARCHAR(255) NOT NULL , `password` VARCHAR(255) NOT NULL , `creation_date` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP , PRIMARY KEY (`id`)) ENGINE = InnoDB;")
	checkErr(err)

	//Crypto Table
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `" + dbName + "`.`crypto` ( `id` INT NOT NULL AUTO_INCREMENT , `user_id` INT NOT NULL , `addOrDeduct` VARCHAR(1) NOT NULL,  `currency` VARCHAR(255) NOT NULL , `amount` VARCHAR(255) NOT NULL  , `creation_date` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP , PRIMARY KEY (`id`)) ENGINE = InnoDB;")
	checkErr(err)

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
	json.Unmarshal(getBytesByUrl(urlCMCBTC), &statusCMCBTC)
	json.Unmarshal(getBytesByUrl(urlCMCLTC), &statusCMCLTC)
}

/*
*Gets the bytes for a json url
 */
func getBytesByUrl(jsonUrl string) []byte {

	//Get respons form url
	resp, err := http.Get(jsonUrl)
	checkErr(err)

	//Turn json into bytes
	bytes, _ := ioutil.ReadAll(resp.Body)

	return bytes
}

/*
*Function used to load in a particulary users currency status
 */
func loadUserCrypto(userId int) {

	//Establish connection
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/")
	checkErr(err)
	defer db.Close()

	rows, err := db.Query("SELECT * FROM `"+dbName+"`.`crypto` WHERE user_id = ?", userId)
	checkErr(err)

	for rows.Next() {
		var id int
		var user_id int
		var currency string
		var addOrDeduct string
		var amount float64
		var creation_date string
		err = rows.Scan(&id, &user_id, &addOrDeduct, &currency, &amount, &creation_date)
		checkErr(err)

		cUR := userCryptoRecord{Id: id, Amount: amount, AddOrDeduct: addOrDeduct, CreationDate: creation_date}

		//Reinitialize the map when this isn't declared yet
		if len(userCryptoRecords[currency]) == 0 {
			userCryptoRecords[currency] = map[int]userCryptoRecord{}
		}

		userCryptoRecords[currency][id] = cUR

	}
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
	displayStatusCoinmarketcap()
	fmt.Println("\n")
}

func displayStatusCoinmarketcap() {
	printHighBlue.Println("\nCoinmarketcap status BTC: ")
	printHighGreen.Print("€")
	printHighGreen.Printf("%v", statusCMCBTC[0].PriceEur)

	printHighBlue.Println("\nCoinmarketcap status LTC: ")
	printHighGreen.Print("€")
	printHighGreen.Printf("%v", statusCMCLTC[0].PriceEur)
}

/*--------------
WEB DISPLAY
--------------*/

/*
* Favocin launch handler
 */
func handlerICon(w http.ResponseWriter, r *http.Request) {}

/*
*Index page
 */
func indexPageHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("html/index.html")

	//Printing the error of the template page
	fmt.Println(t.Execute(w, nil))
}

/*
*transactions page
 */
func transactionsPageHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("html/transactions.html")

	loadUserCrypto(1)

	p := pageVariables{Title: "transactions", Variables: userCryptoRecords}

	//Printing the error of the template page
	fmt.Println(t.Execute(w, p))
}

/*
*Crypto page
 */
func transactionsNewPageHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		t, _ := template.ParseFiles("html/new_transaction.html")
		fmt.Println(t.Execute(w, nil))
	} else {
		r.ParseForm()

		//Establish connection
		db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/")
		checkErr(err)
		defer db.Close()

		// insert
		stmt, err := db.Prepare("INSERT `" + dbName + "`.`crypto` SET user_id=?,addOrDeduct=? ,currency=?,amount=?")
		checkErr(err)

		res, err := stmt.Exec(1, r.Form["addOrDeduct"][0], r.Form["currency"][0], r.Form["amount"][0])
		checkErr(err)

		id, err := res.LastInsertId()
		checkErr(err)

		fmt.Println(id)

		http.Redirect(w, r, "/", 301)
	}
}

/*--------------
MAIN FUNCTION
--------------*/
func main() {

	// Function only needed on first launch
	go loadDatabase(dbName)

	// http loading
	http.HandleFunc("/favicon.ico", handlerICon)
	http.HandleFunc("/", indexPageHandler)
	http.HandleFunc("/transactions", transactionsPageHandler)
	http.HandleFunc("/transactions/new", transactionsNewPageHandler)
	http.ListenAndServe(":3000", nil)
}
