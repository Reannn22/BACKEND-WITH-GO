package main

import (
    "encoding/json"
    "fmt"
)

func main() {
    var jsonString = `{"Name": "john wick", "Age": 27}`
    var jsonData = []byte(jsonString)
    var data1 map[string]interface{}

    err := json.Unmarshal(jsonData, &data1)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Println("user:", data1["Name"])
    fmt.Println("age:", data1["Age"])
}
