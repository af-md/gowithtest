package main

// import "fmt"

// func main() {
// 	fmt.Println(HelloWorld("Afzal", ""))
// }

const enPrefix = "Hello, "
const spPrefix = "Hola, "

func HelloWorld(name string, language string) string {
	if language == "Spanish" {
		return spPrefix + name
	}
	if name == "" {
		name = "world"
	}
	return enPrefix + name
}
