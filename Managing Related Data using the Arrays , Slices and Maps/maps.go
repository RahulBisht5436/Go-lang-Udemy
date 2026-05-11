package main

import "fmt"

func main() {
	//here when we initate a  Map it is like map[key_type]valueType
	website := map[string]string{
		"Google": "https/www.google.com",
		"Yahoo":  "https//www.yahoo.com",
	}
	// This is how we Mutate Value inside the Map
	website["Google"] = "Amazon Web Services"

	//delete the key inside the Map
	delete(website, "Google")

	//This is how we access the Map Value
	fmt.Println(website)

}
