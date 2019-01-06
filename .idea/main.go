package main

import (
	"gopkg.in/go-resty/resty.v0";
	"fmt"
	"time"
	//"encoding/json"
	//"log"
	//"io/ioutil"
	"os"
)

type Askbid struct {

	Price	[]float32	`json:price`
	Size	[]float32	`json:size`
}


type Response struct {


	Time		string	`json:time_exchange`
	Exchange	string	`json:symbol_id`
	Asks		Askbid	`json:asks`
	Bids		Askbid	`json:bids`

}
func doSomething(f *os.File) {


	resp, _ := resty.R().
		SetHeader("X-CoinAPI-Key", "1D79C01D-F218-457E-95D9-AD05DD21F3B7").
	//	Get("https://rest.coinapi.io/v1/orderbooks/current?filter_asset_id_base=ETH&filter_asset_id_quote=BTC&limit_levels=3")
		Get("https://rest.coinapi.io/v1/orderbooks/current?filter_symbol_id=COINBASE_SPOT_ETH_BTC,BINANCE_SPOT_ETH_BTC,BITFINEX_SPOT_ETH_BTC,BITFLYER_SPOT_ETH_BTC,BIBOX_SPOT_ETH_BTC,HITBTC_SPOT_ETH_BTC,KUCOIN_SPOT_ETH_BTC,CRYPTOPIA_SPOT_ETH_BTC,POLONIEX_SPOT_ETH_BTC,HUOBIPRO_SPOT_ETH_BTC,UPBIT_SPOT_ETH_BTC,QUOINE_SPOT_ETH_BTC,OKEX_SPOT_ETH_BTC&limit_levels=3")


	//
	//responseData, err := ioutil.ReadAll(resp.Body())
	//
	//if err != nil {
	//	log.Fatal(err)
	//}



	fmt.Printf("\nBody: %v", resp.String())
	fmt.Fprintf(f, resp.String())


}


func startPolling() {




	f, err := os.Create("/Users/gusiri/Desktop/exchangedata")

	if err != nil{
		panic(err)
	}

	defer f.Close()



	for _ = range time.Tick(600 * time.Second) {

	doSomething(f)

	}

}



func main() {

	// exchange list

	/*
	resp, _ := resty.R().
	SetHeader("X-CoinAPI-Key", "73034021-0EBC-493D-8A00-E0F138111F41").
	Get("https://rest.coinapi.io/v1/exchanges")

	fmt.Printf("\nBody: %v", resp.String())
	*/


	/*
		{exchange_id}_SPOT_{asset_id_base}_{asset_id_quote}
		ex: COINBASE_SPOT_BTC_ETH
	*/

	//Get("https://rest.coinapi.io/v1/quotes/current?filter_symbol_id={COINBASE_SPOT_BTC_ETH, BINANCE_SPOT_BTC_ETH, BITFINEX_SPOT_BTC_ETH}&limit_levels={5}")


/*
		resp, _ := resty.R().
			SetHeader("X-CoinAPI-Key", "1D79C01D-F218-457E-95D9-AD05DD21F3B7")
			Get("https://rest.coinapi.io/v1/orderbooks/current?filter_symbol_id=COINBASE_SPOT_ETH_BTC,BINANCE_SPOT_ETH_BTC,BITFINEX_SPOT_ETH_BTC&limit_levels=3")
*/



	//doSomething()
	go startPolling()

	select{}



}