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

// TokenValidation recibe SOLO la parte del JWT (sin el prefijo "Bearer ")
// Ejemplo de 'token' válido: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImpvaG4iLCJleHAiOjE3MDAwMDAwMDB9.signature"

func TokenValidation(token string) (bool, string, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		fmt.Println("Invalid Token. Lenght != 3")
		return false, "Invalid Token", nil
	}

	// 2) Decodificar la segunda parte (payload) con Base64 Raw URL‐Safe
	//    Esto acepta '-' y '_' y no requiere '=' de padding.
	rawPayload := parts[1]
	payloadBytes, err := base64.RawURLEncoding.DecodeString(rawPayload)
	if err != nil {
		fmt.Println("Decoding token error:", err.Error())
		return false, err.Error(), err
	}
	// 3) Unmarshall del JSON dentro de payloadBytes a nuestra estructura TokenJSON
	var tkj TokenJSON
	if err := json.Unmarshal(payloadBytes, &tkj); err != nil {
		fmt.Println("Cannot be decoded as JSON struct:", err.Error())
		return false, err.Error(), err
	}

	// 4) Verificar expiración (campo Exp asume un timestamp Unix en segundos)
	now := time.Now()
	expiration := time.Unix(int64(tkj.Exp), 0)
	if expiration.Before(now) {
		fmt.Println("Expiration date taken =", expiration.String())
		fmt.Println("Expired Token")
		return false, "Expired Token", nil
	}

	return true, string(tkj.Username), err
}
