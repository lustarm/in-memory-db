package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"imd/src/db"

	"github.com/gorilla/mux"
)

type ApiRequest struct {
    Key string `json:"key"`
    Value string `json:"value"`
}

func StartApi() error {
    r := mux.NewRouter()
    r.HandleFunc("/", baseHandle)
    r.HandleFunc("/create", createHandle).Methods("POST")
    r.HandleFunc("/read", readHandle)

    srv := &http.Server{
        Addr:         "0.0.0.0:8000",
        // Good practice to set timeouts to avoid Slowloris attacks.
        WriteTimeout: time.Second * 15,
        ReadTimeout:  time.Second * 15,
        IdleTimeout:  time.Second * 60,
        Handler: r, // Pass our instance of gorilla/mux in.
    }

    go func() {
        if err := srv.ListenAndServe(); err != nil {
            log.Println(err)
        }
    }()

    var wait time.Duration

    c := make(chan os.Signal, 1)
    // We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
    // SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
    signal.Notify(c, os.Interrupt)

    // Block until we receive our signal.
    <-c
    // Create a deadline to wait for.
    ctx, cancel := context.WithTimeout(context.Background(), wait)
    defer cancel()
    // Doesn't block if no connections, but will otherwise wait
    // until the timeout deadline.
    srv.Shutdown(ctx)
    // Optionally, you could run srv.Shutdown in a goroutine and block on
    // <-ctx.Done() if your application should wait for other services
    // to finalize based on context cancellation.
    log.Println("Shutting down")

    return nil
}

func baseHandle(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "IMD In Memory Database!")
}

func createHandle(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)

    var t ApiRequest
    
    err := json.NewDecoder(r.Body).Decode(&t)

    if err != nil {
        json.NewEncoder(w).Encode(jsonResp{"error" : true,
            "message" : err.Error()})
        return
    }

    if t.Key == "" || t.Value == ""{
        json.NewEncoder(w).Encode(jsonResp{"error" : true, 
            "message" : "Please provide a key and value to use"})
        return
    }
    err = db.Db.Create(t.Key, t.Value)
    if err != nil {
        json.NewEncoder(w).Encode(jsonResp{"error" : true, "message" : err.Error()})
        return
    }

    json.NewEncoder(w).Encode(jsonResp{"error" : false, 
        "message" : "Created key '" + t.Key + "' with value '" + t.Value + "'"})
}

type jsonResp map[string]interface{}

func readHandle(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    key := "test"
    value, err := db.Db.Read(key)
    if err != nil {
        json.NewEncoder(w).Encode(jsonResp{"error" : true, "message" : err.Error()})
        return
    }

    json.NewEncoder(w).Encode(jsonResp{"key" : key, "value" : value})
}
