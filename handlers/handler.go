package handlers

import (
	"encoding/json"
	"merchantBank/models"
	"merchantBank/repository"
	"net/http"
	"sync"
)

var (
    loggedInUsers = make(map[int]bool)
    mu            sync.Mutex
)

func Login(w http.ResponseWriter, r *http.Request) {
    var request struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }
    json.NewDecoder(r.Body).Decode(&request)

    customers, _ := repository.LoadCustomers()

    for _, customer := range customers {
        if customer.Username == request.Username && customer.Password == request.Password {
            mu.Lock()
            loggedInUsers[customer.ID] = true
            history, _ := repository.LoadHistory()
            history = append(history, models.History{
                CustomerID: customer.ID,
                Action:     "login",
            })
            repository.SaveHistory(history)
            mu.Unlock()

            w.WriteHeader(http.StatusOK)
            json.NewEncoder(w).Encode("Login Successful")
            return
        }
    }

    w.WriteHeader(http.StatusUnauthorized)
    json.NewEncoder(w).Encode("Invalid Username or Password")
}

func Payment(w http.ResponseWriter, r *http.Request) {
    var request struct {
        CustomerID int     `json:"customer_id"`
        Amount     float64 `json:"amount"`
    }
    json.NewDecoder(r.Body).Decode(&request)

    mu.Lock()
    if !loggedInUsers[request.CustomerID] {
        mu.Unlock()
        w.WriteHeader(http.StatusUnauthorized)
        json.NewEncoder(w).Encode("Customer Not Logged In")
        return
    }
    mu.Unlock()

    customers, _ := repository.LoadCustomers()

    for i, customer := range customers {
        if customer.ID == request.CustomerID {
            customers[i].Balance -= request.Amount
            repository.SaveCustomers(customers)

            history, _ := repository.LoadHistory()
            history = append(history, models.History{
                CustomerID: customer.ID,
                Action:     "payment",
                Amount:     request.Amount,
            })
            repository.SaveHistory(history)

            w.WriteHeader(http.StatusOK)
            json.NewEncoder(w).Encode("Payment Successful")
            return
        }
    }

    w.WriteHeader(http.StatusBadRequest)
    json.NewEncoder(w).Encode("Customer Not Found")
}

func Logout(w http.ResponseWriter, r *http.Request) {
    var request struct {
        CustomerID int `json:"customer_id"`
    }
    json.NewDecoder(r.Body).Decode(&request)

    mu.Lock()
    delete(loggedInUsers, request.CustomerID)
    history, _ := repository.LoadHistory()
    history = append(history, models.History{
        CustomerID: request.CustomerID,
        Action:     "logout",
    })
    repository.SaveHistory(history)
    mu.Unlock()

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode("Logout Successful")
}
