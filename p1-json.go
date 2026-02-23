package main

import (
    "fmt"
    "encoding/json"
)

type Person struct {
    Name string `json:"name"`
    Age  int `json:"age"`
    Email string `json:"email"`
}

func JSONFunc() {
    jsonData := `{"name": "abku", "age": 22, "email": "abhishek@abhishek.li"}`

    var p Person
    // JSON to struct
    err := json.Unmarshal([]byte(jsonData), &p)
    if err != nil {
        fmt.Println("error: ", err)
        return
    }

    fmt.Println(p)

    // fmt.Println("person name: ", p.Name)
    // fmt.Println("person age: ", p.Age)
    // fmt.Println("person email: ", p.Email)

    // Struct to JSON
    jsonBytes, err := json.Marshal(p)
    if err != nil {
        fmt.Println("error: ", err)
        return
    }

    fmt.Println("json output: ", string(jsonBytes))

}