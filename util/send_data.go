package util

import (
	"encoding/json"
	"net/http"
)

func SendData(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w) // w হচ্ছে কোথায় JSON ডেটা যাবে — সেটা বোঝায়।
	encoder.Encode(data)          // productList কে JSON format এ convert করে
}
