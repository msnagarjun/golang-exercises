package main

func main() {

	// Connect to Database
	connectToDatabase()

	// Create a Table

	user := User{}
	user.Age = 21
	user.FirstName = "Hasher"
	user.LastName = "Gopher"
	user.Email = "hasher.gopher@hashedin.com"
	addUser(user)

	// Query data from database
	// users := []User{}
	users := getUsers()

	//Query for a specific user in database
	getUser(users[0].ID)

	// Cleanup the table
	deleteAllUsers()

	// load data from the json file
	loadAllUsers("data/users.json")

	// Cleanup the table
	deleteAllUsers()

}
