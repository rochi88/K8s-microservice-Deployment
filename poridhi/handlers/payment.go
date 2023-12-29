package handlers

import (
	"log"
	"net/http"
	"io"
    "bytes"
    "encoding/json"
)

type Payment struct {
	l *log.Logger
}

type PaymentData struct {
    Amount int `json:"amount"`
}

func NewPayment(l *log.Logger) *Payment{
	return &Payment{l}
}

func (p *Payment) LoadPayments(rw http.ResponseWriter, r *http.Request) {
    p.l.Println("Handle GET Payments")

    // Fetch the data from fruit-api
    resp, err := http.Get("http://app-dotnet:8080/api/LoadPayments")
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close() // Close the response body

    // Read the response body
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        log.Fatalln(err)
    }

    // Write the body to the client
    rw.WriteHeader(resp.StatusCode)
    rw.Write(body)
}

func (p *Payment) Pay(w http.ResponseWriter, r *http.Request) {

    p.l.Println("Handle POST Request")

    var newPayment PaymentData
    err := json.NewDecoder(r.Body).Decode(&newPayment)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Send the product to the other service as a POST request
    payload, err := json.Marshal(newPayment)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    url := "http://app-dotnet:8080/api/Pay"
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		p.l.Printf("Payment added successfully")
		w.WriteHeader(http.StatusOK)
	} else {
		http.Error(w, "Payment not added", resp.StatusCode)
		return
	}
}



