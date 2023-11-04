package main

import (
	"github.com/Metchain/Metblock/newwallet/metwallet"
	"io"
	"log"
	"net/http"
)

func (ws *WalletServer) WalletRecover(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		w.Header().Add("Content-Type", "application/json")

		seed := req.URL.Query().Get("seed")
		myWallet := metwallet.WalletRecover(seed)
		log.Println(myWallet)
		m, _ := myWallet.MarshalJSON()
		io.WriteString(w, string(m[:]))
	case http.MethodPost:
		w.Header().Add("Content-Type", "application/json")
		myWallet := metwallet.NewWallet()
		m, _ := myWallet.MarshalJSON()
		io.WriteString(w, string(m[:]))
	default:
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error: Invalid HTTP Method")

	}
}
