package main

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

func (ws *WalletServer) Index(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:

		t, _ := template.ParseFiles(path.Join(tempDir, "index.html"))
		t.Execute(w, "")
	default:
		log.Printf("ERROR: Invalid HTTP Method")
	}
}

func (ws *WalletServer) Run() {
	http.HandleFunc("/", ws.Index)
	http.HandleFunc("/style/style.css", ws.Style)
	http.HandleFunc("/wallet", ws.Wallet)
	http.HandleFunc("/walletrecover", ws.WalletRecover)
	http.HandleFunc("/wallet/amount", ws.WalletAmount)
	http.HandleFunc("/transaction", ws.CreateTransaction)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(ws.Port())), nil))

}

func (ws *WalletServer) Style(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		t, _ := template.ParseFiles(path.Join(tempDir, "/style/style.css"))
		t.Execute(w, "")

	default:
		log.Printf("ERROR: Invalid HTTP Method")
	}
}
