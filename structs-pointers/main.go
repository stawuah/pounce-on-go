package main

import (
	"fmt"
	"time"
)

// --- PATTERN #1: Basic Struct Definition and Usage (Value Types) ---

// UserProfile is a simple struct with value types.
// Use this when the data is small, and you don't need to modify it
// in place within a method or function.
type UserProfile struct {
	ID        int
	Username  string
	CreatedAt time.Time
}

// GetUsername is a method with a value receiver (p UserProfile).
// Signature: func (receiver Type) MethodName()
// This method operates on a COPY of the UserProfile struct.
// It is the correct pattern when the method does not need to
// modify the original struct.
func (p UserProfile) GetUsername() string {
	return p.Username
}

// --- PATTERN #2: Pointers to Structs (Efficiency & Mutability) ---

// LargeData is a complex struct that would be expensive to copy.
// The `Config` field is a pointer because it is a separate,
// complex data structure that we want to manage by reference.
type LargeData struct {
	Name    string
	SizeMB  int
	Config  *ServiceConfig // A pointer to a nested struct
	Metrics *Metrics       // Another pointer to a nested struct
}

// ServiceConfig holds configuration details.
type ServiceConfig struct {
	Host string
	Port int
}

// Metrics holds runtime performance metrics.
type Metrics struct {
	CPUUsage float64
	MemoryMB float64
}

// UpdateMetrics is a method with a pointer receiver (d *LargeData).
// Signature: func (receiver *Type) MethodName()
// This is the correct pattern because the method MODIFIES the
// original `LargeData` struct's `Metrics` field.
// We use a pointer to avoid copying the large `LargeData` struct.
func (d *LargeData) UpdateMetrics(cpu, mem float64) {
	// First, check if the Metrics pointer is nil. This is a crucial pattern
	// for handling optional or lazily-initialized data.
	if d.Metrics == nil {
		d.Metrics = &Metrics{}
	}
	d.Metrics.CPUUsage = cpu
	d.Metrics.MemoryMB = mem
}

// --- PATTERN #3: Pointers for Optionality (Distinguishing `nil` from Zero-Value) ---

// RequestData might have optional fields.
// Pointers to basic types are used here to distinguish
// between an unset value (nil) and a zero value (0 or "").
type RequestData struct {
	UserID  *int    // A pointer to an int. Can be nil.
	Status  *string // A pointer to a string. Can be nil.
	Payload string  // A regular string, which will be empty ("") if not set.
}

// SetUserID sets the UserID field. The value is passed by value,
// and we take its address to set the pointer.
func (rd *RequestData) SetUserID(id int) {
	rd.UserID = &id
}

// IsUserIDSet checks if the UserID was provided.
func (rd *RequestData) IsUserIDSet() bool {
	return rd.UserID != nil
}

// --- PATTERN #4: Pointers to Slices (In-place Modification) ---

// ResourceManager manages a list of resources.
// We use a pointer to a slice of strings because we want to modify
// the list itself (add/remove items) without a full copy.
type ResourceManager struct {
	ID          string
	ResourceIDs *[]string // A pointer to a slice of strings
}

// AddResource adds a new ID to the resource list.
func (rm *ResourceManager) AddResource(id string) {
	// Check if the pointer is nil first. This is a common defensive pattern.
	if rm.ResourceIDs == nil {
		rm.ResourceIDs = &[]string{}
	}
	// The `*` dereferences the pointer to get the slice value,
	// allowing `append` to modify the original slice.
	*rm.ResourceIDs = append(*rm.ResourceIDs, id)
}

// --- PATTERN #5: Pointer to a Slice of Pointers (Advanced) ---

// DataManager is a complex type that holds a dynamic list of large data objects.
// `Items` is a pointer to a slice of pointers.
// This is used for maximum efficiency when managing a dynamic collection of
// large structs, as both the slice and the individual structs are passed by reference.
type DataManager struct {
	Items *[]*LargeData
}

// AddItem adds a new LargeData struct to the manager.
func (dm *DataManager) AddItem(name string, size int) {
	if dm.Items == nil {
		dm.Items = &[]*LargeData{}
	}
	// Create a new LargeData struct and get its address.
	newItem := &LargeData{Name: name, SizeMB: size}
	// Append the pointer to the new item to the slice.
	*dm.Items = append(*dm.Items, newItem)
}

// --- MAIN FUNCTION TO DEMONSTRATE ALL PATTERNS ---

func main() {
	// Pattern #1: Value Struct
	user := UserProfile{ID: 1, Username: "gopher", CreatedAt: time.Now()}
	fmt.Printf("User: %s\n", user.GetUsername())

	fmt.Println("\n--- Demonstrating Pointers ---")

	// Pattern #2: Pointer to Struct
	// Create a new `LargeData` struct. Using `new` returns a pointer.
	largeData := new(LargeData)
	largeData.Name = "MyService"
	largeData.SizeMB = 50
	// The `Config` pointer is currently nil. Let's initialize it.
	largeData.Config = &ServiceConfig{Host: "localhost", Port: 8080}
	fmt.Printf("LargeData Config: %+v\n", *largeData.Config)

	// Call a method with a pointer receiver. This will modify the struct.
	largeData.UpdateMetrics(15.5, 1024.0)
	fmt.Printf("Updated Metrics: %+v\n", *largeData.Metrics)

	// Pattern #3: Pointers for Optionality
	request := new(RequestData)
	fmt.Printf("\nInitial request (UserID is nil): %+v\n", request)
	if request.IsUserIDSet() {
		fmt.Println("This should not be printed.")
	}
	request.SetUserID(99)
	fmt.Printf("Request after setting UserID: %+v\n", request)
	if request.IsUserIDSet() {
		fmt.Printf("UserID is set to: %d\n", *request.UserID)
	}

	// Pattern #4: Pointers to Slices
	resourceMgr := &ResourceManager{ID: "worker-01"}
	// The ResourceIDs pointer is nil. The AddResource method handles this.
	resourceMgr.AddResource("db-conn-123")
	resourceMgr.AddResource("net-conn-456")
	fmt.Printf("\nResource Manager State: %+v\n", resourceMgr)
	fmt.Printf("ResourceIDs: %+v\n", *resourceMgr.ResourceIDs)

	// Pattern #5: Pointer to a Slice of Pointers
	dataManager := new(DataManager)
	dataManager.AddItem("item-1", 100)
	dataManager.AddItem("item-2", 200)

	fmt.Println("\nDataManager Items:")
	// We need to dereference the outer pointer to get the slice,
	// then we can range over the slice of pointers to structs.
	for i, item := range *dataManager.Items {
		fmt.Printf("  Item %d: Name=%s, Size=%d\n", i+1, item.Name, item.SizeMB)
	}
}
