package main

import(
	//"fmt"
	"log"
	//"database/sql"
	"github.com/gorilla/mux"
	//"strconv"
	"net/http"
	"encoding/json"
	//_"github.com/lib/pq"
	//"github.com/go-sql-driver/mysql"
)

func main() {
	
	Routes := mux.NewRouter().StrictSlash(true);
	
		Routes.HandleFunc("/hello", getHello).Methods("GET")
	
		log.Fatal(http.ListenAndServe(":3600", Routes))

}

type Example struct {
	ID uint "json:*id"
	Example string "json:*example"
}

var examples = []Example{
	Example{ID:1, Example:"Lorem Ipsum"},
	Example{ID:2, Example:"Lorem Ipsum"},
}

func getHello(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(examples)

	//respondWithJSON(w, http.StatusOK, p)

}
