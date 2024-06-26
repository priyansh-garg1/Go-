package main

import "fmt"

func main() {
	// fmt.Println("I am Learning Go Languages")

	// // Variables
	// var name string  = "Hello"
	// var version = 1.001
	// version = 2.001
	// var dec float32 = 0.001
	// fmt.Println(dec);/

	// age := 12
	// name := "User"
	// person := 123
	// var agee int = 12
	// fmt.Println("person: ", person, "name: ", name, "age", age)
	// fmt.Printf("age is %d", agee)
	// fmt.Printf("name is %T", name)

	// Input User
	// fmt.Println("Enter your name")
	// var name string
	// Only Take first word
	// fmt.Scan(&name)
	// fmt.Printf("Hello %s", name)
	// // to get full line
	// reader := bufio.NewReader(os.Stdin)
	// name, _ = reader.ReadString('\n')
	// fmt.Println(name)

	// simpleFunction()
	// ans, err := div(5, 0)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	ans, _ := div(5, 2)
	fmt.Println("Division:", ans)
}

// // Function
// func simpleFunction() {
// 	fmt.Print("Function")
// 	fmt.Println("val is:", add(10, 20))
// }

// func add(a, b int) int {
// 	fmt.Println(a + b)
// 	return a + b
// }

func div(a, b float32) (float32, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}
