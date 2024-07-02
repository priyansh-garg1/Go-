package main

import (
	"fmt"
	"io"
	"net/http"
)

type Todo struct {
	ID          int    `json:"id"`
	UserID       int    `json:"userId"`
	Title       string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	fmt.Println("CRUD!")
	res, err := http.Get("http://jsonplaceholder.typicode.com/todo/1")
	if err != nil {
		fmt.Println("Error getting Get response", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Printf("Error: %d - %s\n", res.StatusCode, http.StatusText(res.StatusCode))
        return     }
	}

	// data,err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("Error reading response body", err)
	// 	return
	// }
	// fmt.Println("Data received!", string(data))


	var todo Todo 
	err := json.NewDecoder(res.body).Decode(&todo);
	if err != nil {
		fmt.Println("Error decoding JSON", err)
        return
    }
	fmt.Println(todo)


}
