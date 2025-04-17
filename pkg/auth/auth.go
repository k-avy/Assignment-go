package auth

import "fmt"

const expectedUsername = "admin"
const expectedPassword = "secret"

func Authenticate(username string, password string) bool {
	if username != expectedUsername || password != expectedPassword {
		fmt.Println("Authentication failed.")
		return false
	}
	return true
}
