package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var bc *Blockchain

func respondWithJson(w http.ResponseWriter, r *http.Request, statusCode int, payload interface{}) {
	resp, err := json.MarshalIndent(payload, "", " ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("HTTP 500: Internal Server Error"))
		return
	}

	w.WriteHeader(statusCode)
	w.Write(resp)
}

func handleGetBlockChain(w http.ResponseWriter, r *http.Request) {
	req, err := json.MarshalIndent(bc, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	io.WriteString(w, string(req))
}

func handleWriteBlock(w http.ResponseWriter, r *http.Request) {
	var data string
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&data); err != nil {
		log.Fatal("Error while decoding data", err.Error())
		respondWithJson(w, r, http.StatusInternalServerError, data)
		return
	}
	defer r.Body.Close()

	newBlock := bc.AddBlock(data)
	// TODO: Add block validation logic here
	respondWithJson(w, r, http.StatusCreated, newBlock)
}

func makeMuxRouter() http.Handler {
	muxRouter := mux.NewRouter()
	muxRouter.HandleFunc("/", handleGetBlockChain).Methods("GET")
	muxRouter.HandleFunc("/", handleWriteBlock).Methods("POST")
	return muxRouter
}

func run() error {
	mux := makeMuxRouter()
	httpAddr := os.Getenv("ADDR")
	log.Println("Listening on", os.Getenv("ADDR"))
	s := &http.Server{
		Addr:           ":" + httpAddr,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := s.ListenAndServe(); err != nil {
		return err
	}

	return nil
}

func main() {
	bc = NewBlockchain()

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}
