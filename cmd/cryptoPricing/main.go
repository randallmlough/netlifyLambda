package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/apex/gateway"
)

// Coin struct
type Coin struct {
	Name   string
	Symbol string
	Slug   string
	Rank   int
	Quote  Quote
}

type Quote struct {
	Currency string
	Price    float64
}

type Query struct {
	Crypto []string `json:"crypto"`
	Pair   []string `json:"pair"`
}

var m map[string]interface{}

// APIHandler is a http.HandlerFunc for the / path.
func APIHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	fmt.Println(r.Form) // print information on server side.

	query := Query{
		Crypto: r.Form["crypto"],
		Pair:   r.Form["pair"],
	}

	crypto := strings.Join(query.Crypto, "")
	pair := strings.Join(query.Pair, "")

	u := &url.URL{
		Scheme:   "https",
		Host:     "pro-api.coinmarketcap.com",
		Path:     "/v1/cryptocurrency/quotes/latest",
		RawQuery: "symbol=" + crypto + "&convert=" + pair,
	}
	fmt.Println(u.String())

	client := &http.Client{
		CheckRedirect: checkRedirectFunc,
	}

	req, _ := http.NewRequest("GET", u.String(), nil)
	env, _ := os.LookupEnv("COINMARKETCAP")
	req.Header.Add("X-CMC_PRO_API_KEY", env)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(data))

	json.Unmarshal(data, &m)
	if m == nil {
		fmt.Println("No data was fetched")
	}

	coins := m["data"].(map[string]interface{})
	for k := range coins {
		coin := coins[k].(map[string]interface{})

		quote := coin["quote"].(map[string]interface{})
		for k := range quote {

			currency := quote[k].(map[string]interface{})

			coinDetails := Coin{
				Name:   coin["name"].(string),
				Symbol: coin["symbol"].(string),
				Slug:   coin["slug"].(string),
				Rank:   int(coin["cmc_rank"].(float64)),
				Quote: Quote{
					Currency: k,
					Price:    currency["price"].(float64),
				},
			}
			fmt.Println(coinDetails)
			js, err := json.Marshal(coinDetails)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write(js)
		}
	}

}

func checkRedirectFunc(req *http.Request, via []*http.Request) error {
	req.Header.Add("X-CMC_PRO_API_KEY", via[0].Header.Get("X-CMC_PRO_API_KEY"))
	return nil
}

// h wraps a http.HandlerFunc and adds common headers.
func h(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf8")
		next.ServeHTTP(w, r)
	})
}

func main() {
	http.Handle("/", h(APIHandler))
	log.Fatal(gateway.ListenAndServe(":9000", nil))
}
