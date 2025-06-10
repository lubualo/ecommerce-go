// auth/auth_validation.go
package auth

import (
	"fmt"
	"strings"
)

// AuthValidation validates whether a request should be allowed based on path, method, and token
func AuthValidation(path, method string, header map[string]string) (bool, int, string) {
	// Public routes
	if (path == "product" && method == "GET") ||
		(path == "category" && method == "GET") {
		return true, 200, ""
	}

	rawAuth := header["authorization"]
	if len(rawAuth) == 0 {
		return false, 401, "Required Token"
	}

	var token string
	if strings.HasPrefix(strings.ToLower(rawAuth), "bearer ") {
		token = rawAuth[len("Bearer "):]
	} else {
		token = rawAuth
	}

	isOk, msg, err := TokenValidation(token)
	if !isOk {
		if err != nil {
			fmt.Println("Token Error " + err.Error())
			return false, 401, err.Error()
		}
		fmt.Println("Token Error " + msg)
		return false, 401, msg
	}

	fmt.Println("Token OK")
	return true, 200, msg
}
