package repository

import (
	"encoding/json"
	"io/ioutil"
	"merchantBank/models"
	"os"
)

func LoadCustomers() ([]models.Customer, error) {
    file, err := os.Open("data/customers.json")
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var customers []models.Customer
    byteValue, _ := ioutil.ReadAll(file)
    json.Unmarshal(byteValue, &customers)

    return customers, nil
}

func SaveCustomers(customers []models.Customer) error {
    file, err := json.MarshalIndent(customers, "", " ")
    if err != nil {
        return err
    }
    return ioutil.WriteFile("data/customers.json", file, 0644)
}

func LoadHistory() ([]models.History, error) {
    file, err := os.Open("data/history.json")
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var history []models.History
    byteValue, _ := ioutil.ReadAll(file)
    json.Unmarshal(byteValue, &history)

    return history, nil
}

func SaveHistory(history []models.History) error {
    file, err := json.MarshalIndent(history, "", " ")
    if err != nil {
        return err
    }
    return ioutil.WriteFile("data/history.json", file, 0644)
}
