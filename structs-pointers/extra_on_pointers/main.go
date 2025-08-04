package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

// Product is our data model. It represents a single product.
type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

// ProductService is a struct that holds the state of our application.
// In a real application, this would be a database connection.
// It uses a mutex to ensure thread-safe access to the products map.
type ProductService struct {
	products map[int]Product
	mu       sync.RWMutex
	nextID   int
}

// NewProductService is a constructor function that returns a pointer
// to a new ProductService instance.
func NewProductService() *ProductService {
	return &ProductService{
		products: make(map[int]Product),
		nextID:   1,
	}

}

// CreateProduct is a method with a POINTER RECEIVER (*ProductService).
// This is critical because it allows the method to modify the `products` map
// and `nextID` field of the original ProductService instance.
func (ps *ProductService) CreateProduct(newProduct Product) {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	newProduct.ID = ps.nextID
	ps.products[newProduct.ID] = newProduct
	ps.nextID++

	fmt.Printf("Created new product: ID=%d, Name=%s\n", newProduct.ID, newProduct.Name)
}

// GetProducts is a method with a VALUE RECEIVER.
// It only needs to read data, so a copy of the receiver is fine.
func (ps ProductService) GetProducts() []Product {
	ps.mu.RLock()
	defer ps.mu.RUnlock()

	// Convert the map to a slice for the API response.
	products := make([]Product, 0, len(ps.products))
	for _, p := range ps.products {
		products = append(products, p)
	}
	return products
}

// Handler function for POST /products
// This closure "captures" the pointer to our ProductService.
func createProductHandler(service *ProductService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Only allow POST requests
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Decode the JSON from the request body into a new Product struct.
		var p Product
		// The &p gets the memory address of our new Product struct,
		// so the decoder can write the data directly into it.
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Call the method with a POINTER RECEIVER on the service.
		// This modifies the original `ProductService` instance in memory.
		service.CreateProduct(p)

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("Product created successfully"))
	}
}

// Handler function for GET /products
func getProductsHandler(service *ProductService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Call the method with a VALUE RECEIVER.
		// The service variable is a pointer, so the compiler dereferences it
		// for us before making the call.
		products := service.GetProducts()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(products)
	}
}

func main() {
	// Here, we create our SINGLE instance of the ProductService.
	// We get a pointer to it from the constructor function.
	productService := NewProductService()

	// We pass the SAME pointer to all our handlers.
	// This ensures that every handler is working on the same set of data.
	http.HandleFunc("/products", createProductHandler(productService))
	http.HandleFunc("/products", getProductsHandler(productService))

	fmt.Println("Server is running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
