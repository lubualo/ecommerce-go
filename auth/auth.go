package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type TokenJSON struct {
	Sub       string
	Event_Id  string
	Token_use string
	Scope     string
	Auth_time int
	Iss       string
	Exp       int
	Iat       int
	Client_id string
	Username  string
}

func ValidateToken(token string) (bool, string, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		fmt.Println("Invalid token")
		return false, "Invalid token", nil
	}

	userInfo, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		fmt.Println("Unable to decode the token: ", err.Error())
		return false, err.Error(), err
	}

	var tkj TokenJSON
	err = json.Unmarshal(userInfo, &tkj)
	if err != nil {
		fmt.Println("Unable to decode into JSON: ", err.Error())
		return false, err.Error(), err
	}

	now := time.Now()
	tm := time.Unix(int64(tkj.Exp), 0)
	if tm.Before(now) {
		fmt.Println("Expired token: " + tm.String())
		return false, "Expired token: " + tm.String(), nil
	}

	return true, string(tkj.Username), nil
}
