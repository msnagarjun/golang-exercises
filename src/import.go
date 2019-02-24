package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func loadAllUsers(filePath string) []User {

	// Open our jsonFile
	jsonFile, err := os.Open(filePath)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	var users = []User{}

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &users)

	// we iterate through every user within our users array and print user details
	for i := 0; i < len(users); i++ {
		fmt.Println("User Age: " + strconv.Itoa(int(users[i].Age)))
		fmt.Println("User First Name: " + users[i].FirstName)
		fmt.Println("User Last Name: " + users[i].LastName)
	}

	fmt.Println("Imported " + strconv.Itoa(len(users)) + " users from " + filePath)
	return users
}
