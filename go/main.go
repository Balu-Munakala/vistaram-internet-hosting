package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Product struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Brand    string   `json:"brand"`
	Price    int      `json:"price"`
	Currency string   `json:"currency"`
	Rating   float64  `json:"rating"`
	Reviews  int      `json:"reviews"`
	InStock  bool     `json:"inStock"`
	Image    string   `json:"image"`
	Tags     []string `json:"tags"`
}

var products = []Product{
	{1, "Wireless Headphones", "SoundMax", 1999, "INR", 4.6, 1284, true, "https://picsum.photos/seed/headphones/400/400", []string{"audio", "bluetooth", "bestseller"}},
	{2, "Smart Watch Pro", "TimeX", 3499, "INR", 4.4, 892, true, "https://picsum.photos/seed/watch/400/400", []string{"wearable", "fitness"}},
	{3, "Gaming Mouse RGB", "ClickZ", 899, "INR", 4.8, 2103, false, "https://picsum.photos/seed/mouse/400/400", []string{"gaming", "rgb"}},
	{4, "Mechanical Keyboard", "KeyForge", 2599, "INR", 4.7, 1567, true, "https://picsum.photos/seed/keyboard/400/400", []string{"gaming", "typing"}},
	{5, "4K Webcam", "ClearView", 4299, "INR", 4.3, 654, true, "https://picsum.photos/seed/webcam/400/400", []string{"streaming", "video"}},
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func main() {
	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			var body Product
			json.NewDecoder(r.Body).Decode(&body)
			body.ID = len(products) + 1
			body.Currency = "INR"
			body.InStock = true
			body.Image = "https://picsum.photos/seed/" + strings.ReplaceAll(body.Name, " ", "") + "/400/400"
			products = append(products, body)
			writeJSON(w, http.StatusCreated, map[string]interface{}{"status": "created", "data": body})
			return
		}
		writeJSON(w, http.StatusOK, map[string]interface{}{
			"status": "success", "count": len(products),
			"serverTime": time.Now().Format(time.RFC3339), "data": products,
		})
	})

	http.HandleFunc("/products/", func(w http.ResponseWriter, r *http.Request) {
		id, _ := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/products/"))
		for _, p := range products {
			if p.ID == id {
				writeJSON(w, http.StatusOK, map[string]interface{}{"status": "success", "data": p})
				return
			}
		}
		writeJSON(w, http.StatusNotFound, map[string]string{"status": "error", "message": "not found"})
	})

	http.ListenAndServe(":3002", nil)
}
