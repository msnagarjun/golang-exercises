package main

// User package
type User struct {
	ID        int32
	Age       int32  `json:"age"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
