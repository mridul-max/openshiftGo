package controllers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "myapp/models"
)

func CreateItem(w http.ResponseWriter, r *http.Request) {
    name := r.URL.Query().Get("name")
    priceParam := r.URL.Query().Get("price")

    // Convert price from string to float64
    price, err := strconv.ParseFloat(priceParam, 64)
    if err != nil {
        http.Error(w, "Invalid price value", http.StatusBadRequest)
        return
    }

    // Create a new item and add it to the list
    item := models.Item{
        ID:    len(models.Items) + 1,
        Name:  name,
        Price: price,
    }
    models.Items = append(models.Items, item)

    // Return the newly created item as JSON
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(item)
}
