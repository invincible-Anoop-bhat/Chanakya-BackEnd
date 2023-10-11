package interceptor

import (
	"net/http"
	"os"
)

func validateUser(authTokenString string) bool {
	var AUTHUSER = os.Getenv("AUTH_USERNAME")
	var AUTHPASS = os.Getenv("AUTH_PASSWORD")
	// log.Print(AUTHUSER + ":" + AUTHPASS)
	if string(authTokenString) == (AUTHUSER + ":" + AUTHPASS) {
		return true
	}
	return false
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// authInput := r.Header.Get("Authorization")
		// key := strings.Split(authInput, " ")
		// log.Println("Auth Input   ---------- : ", authInput)
		// decoded, err := base64.StdEncoding.DecodeString(key[1])
		// if err != nil {
		// 	log.Println("Auth Erorr: error decoding , ", err)
		// 	http.Error(w, "Error Authenticating", http.StatusInternalServerError)
		// }
		// log.Print(decoded)

		// log.Print(string(decoded))
		// if validateUser(string(decoded)) == false {
		// 	http.Error(w, "Authentication failed, Wrong credentials ", http.StatusForbidden)
		// }

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
