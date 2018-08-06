package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/apex/gateway"
)

// EnvHandler is a http.HandlerFunc for the / path.
func EnvHandler(w http.ResponseWriter, r *http.Request) {
	// env enivroment variable is set in the template.tml global file and should output "SOMELOCALKEY" Locally and some "SOMESECRETKEY" deployed
	env, ok := os.LookupEnv("ENV_KEY")
	if !ok {
		fmt.Println("NOTHING SET")
	} else {
		fmt.Println(env)
		json.NewEncoder(w).Encode(env)
	}
}

// h wraps a http.HandlerFunc and adds common headers.
func h(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func main() {
	http.Handle("/", h(EnvHandler))
	log.Fatal(gateway.ListenAndServe(":9000", nil))
}
