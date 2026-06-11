package main

import (
	"fmt"
	"encoding/json"
)

type User struct {
	ID int `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`

	internalCode int // this is unexported and will not be included in JSON
}

func main() {
	fmt.Println("Hello, I'm learning JSON in GO!")

	// Struct to JSON ( Marshal )
	user := User{
		ID: 1,
		Username: "abku",
		Email: "abku@abku.dev",
		Password: "Password",
	}

	// Marshal retunrs a byte slice ([]byte) and an error
	jsonData, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	fmt.Println("JsonData ", string(jsonData))

	// JSON to Struct ( Unmarshal )
	jsonString := `{"id":1,"username":"abku","email":"abku@abku.dev","password":"Password"}`
	// Convert JSON string to byte slice
	jsonBytes := []byte(jsonString)
	var user2 User

	err = json.Unmarshal(jsonBytes, &user2)
	if err != nil {
		fmt.Println("Error unmarshaling JSON: ", err)
		return
	}

	fmt.Println(user2.ID, user2.Username, user2.Email, user2.Password)

	// Dealing with unknown JSON fields
	jsonString2 := `{"id":2,"username":"abku2","email":"abku2@abku.dev","password":"Password2","unknownField":"unknownValue"}`

	jsonBytes2 := []byte(jsonString2)

	var result map[string]any // A map where keys are strings and values can be of any type

	json.Unmarshal(jsonBytes2, &result)
	fmt.Println(result)
}

/*
omitempty: When we don't want empty fields to be included in the JSON output, we can use the "omitempty" option in the struct tags. For example"
type User struct {
	ID int `json:"id,omitempty"`
}

type User struct {
	ID int `json:"-"` // This field will be ignored in JSON
}
*/