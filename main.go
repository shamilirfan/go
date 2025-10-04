// go run main.go
// cls
// exit
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var PORT string = ":8080"

// home handler
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Home Page")
}

// Struct define
type Product struct {
	ID          int     `json:"id"` // it is called tag
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	ImageURL    string  `json:"imageUrl"`
}

// slice define
var productList []Product

// product handler
func productHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		fmt.Println(w, "Please give a get request.", 400)
		return
	}

	//
	encoder := json.NewEncoder(w) // w হচ্ছে কোথায় JSON ডেটা যাবে — সেটা বোঝায়।
	encoder.Encode(productList)   // productList কে JSON format এ convert করে
}

func main() {
	mux := http.NewServeMux()

	// router/endpoint
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/products", productHandler)

	// server listening
	fmt.Println("Server is running on http://localhost" + PORT)
	err := http.ListenAndServe(PORT, mux)
	if err != nil {
		fmt.Println("Error starting the server", err)
	}
}

func init() {
	// Struct instance/object create
	product1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "Orange is very jucy. I don't like it.",
		Price:       120,
		ImageURL:    "https://www.croq-kilos.com/sites/default/files/styles/1920px/public/2025-03/26505-3228-tout-savoir-sur-lorange-un-tresor-de-vitamines-et-bien-plus-encore.jpg",
	}
	product2 := Product{
		ID:          2,
		Title:       "Apple",
		Description: "Apple is red. I feel boring eating apple.",
		Price:       210,
		ImageURL:    "https://cdn.britannica.com/60/5760-050-FCD7CDA2/Apples-Malus-domestica.jpg",
	}
	product3 := Product{
		ID:          3,
		Title:       "Banana",
		Description: "Banana is yellow. I like it's smell.",
		Price:       40,
		ImageURL:    "https://images.contentstack.io/v3/assets/bltcedd8dbd5891265b/blt66b8f1f45d70de9e/6665ec0febbbaa7d99a9f533/4440854-Banana-hero.jpg?q=70&width=3840&auto=webp",
	}
	product4 := Product{
		ID:          4,
		Title:       "Mango",
		Description: "Mango is so yummy. I love mango.",
		Price:       145,
		ImageURL:    "https://blog2.pittmandavis.com/wp-content/uploads/2023/07/MangoFinal.jpg",
	}
	product5 := Product{
		ID:          5,
		Title:       "Grapes",
		Description: "Grapes is testy. I like it.",
		Price:       370,
		ImageURL:    "https://img.ksl.com/slc/2513/251331/25133147.jpg?filter=kslv2/responsive_story_lg",
	}

	// append instance/object
	productList = append(productList, product1)
	productList = append(productList, product2)
	productList = append(productList, product3)
	productList = append(productList, product4)
	productList = append(productList, product5)
}
