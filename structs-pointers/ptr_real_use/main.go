package main

import (
	"errors"
	"fmt"
	"sync"
)

type Item struct {
	Name  string
	Price int
}

type Store struct {
	items map[string]Item
	mu    sync.Mutex
}

func NewStore() *Store {
	return &Store{
		items: make(map[string]Item),
	}
}

func (s *Store) AddItem(item Item) {
	s.mu.Lock()

	defer s.mu.Unlock()

	s.items[item.Name] = item
}

func (s *Store) GetItem(name string) (error, *Item) {
	s.mu.Lock()
	defer s.mu.Unlock()

	item, ok := s.items[name]

	if !ok {
		return errors.New("item not found"), nil
	}

	return nil, &item
}

func main() {
	// Create a new store using the constructor function.
	store := NewStore()

	// Create a few Item instances.
	apple := Item{Name: "Apple", Price: 1}
	banana := Item{Name: "Banana", Price: 2}

	// Add the items to the store using the AddItem method.
	store.AddItem(apple)
	store.AddItem(banana)

	fmt.Println("Items successfully added to the store.")

	// Test the GetItem function for a successful case.
	fmt.Println("\n--- Testing GetItem for 'Apple' ---")
	err, retrievedItem := store.GetItem("Apple")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Found item: Name=%s, Price=%d\n", retrievedItem.Name, retrievedItem.Price)
	}

	// Test the GetItem function for a failure case.
	fmt.Println("\n--- Testing GetItem for 'Orange' ---")
	err, retrievedItem = store.GetItem("Orange")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		// This block should not be reached.
		fmt.Printf("Found item: Name=%s, Price=%d\n", retrievedItem.Name, retrievedItem.Price)
	}
}
