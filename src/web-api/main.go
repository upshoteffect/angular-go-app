package main

import (
  "net/http"
  "os"
  "log"
  "fmt"
  "web-api/utils"
  "github.com/rs/cors"
  "github.com/gorilla/mux"
	"github.com/dgrijalva/jwt-go"
	"encoding/json"
  "time"
)

var jwtKey = []byte("my_super_secret_key")

var users = map[string]string{
  "c137@onecause.com": "#th@nH@rm#y#r!$100%D0p#",
}

// Create a struct that models the structure of a user, both in the request body, and in the DB
type LoginData struct {
  User struct {
    Password string `json:"password"`
    Email string `json:"email"`
    Pin int `json:"pin"`
    HoursAndMinutesAtLogin int `json:"hoursAndMinutesAtLogin"`
  } `json:"user"`
}

// Create a struct  that models the structure of a user that gets returned
type ReturnLoginUser struct {
  User struct {
    Email string `json:"email"`
    Token string `json:"token"`
  } `json:"user"`
}

// Create a struct that is used for creating a JWT Token
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Create a struct that is used for errors
type ErrorMessage struct {
	Error string `json:"error"`
}

func main() {
  r := mux.NewRouter()

  r.HandleFunc("/users/login", login)

  // Solves Cross Origin Access Issue
  c := cors.New(cors.Options{
    AllowedOrigins: []string{"http://localhost:4200"},
  })
  handler := c.Handler(r)

  srv := &http.Server{
    Handler: handler,
    Addr:    ":" + os.Getenv("PORT"),
  }

  log.Fatal(srv.ListenAndServe())
}

func login(w http.ResponseWriter, r *http.Request) {

	var loginData LoginData
	// Get the JSON body and decode into a user
  err := json.NewDecoder(r.Body).Decode(&loginData)

	if err != nil {
    // If the structure of the body is wrong, return an HTTP error
    w.WriteHeader(http.StatusBadRequest)
    
    errorMessage := ErrorMessage{"Bad request data."}
    jsonBytes, err := utils.StructToJson(errorMessage); if err != nil {
      fmt.Print(err)
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonBytes)

		return
	}

	// Get the expected password from our in memory map 
	expectedPassword, ok := users[loginData.User.Email]

	// If a password exists for the given user
	// AND, if it is the same as the password we received, the we can move ahead
	// if NOT, then we return an "Unauthorized" status
	if !ok || expectedPassword != loginData.User.Password {
    w.WriteHeader(http.StatusUnauthorized)

    errorMessage := ErrorMessage{"Email and Password do not match."}
    jsonBytes, err := utils.StructToJson(errorMessage); if err != nil {
      fmt.Print(err)
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonBytes)

		return
  }
  
  // Check to see if the pin matches the hours and minutes at log in
  // if NOT, then we return an "Unauthorized" status
  if loginData.User.HoursAndMinutesAtLogin != loginData.User.Pin {
    w.WriteHeader(http.StatusUnauthorized)

    errorMessage := ErrorMessage{"Invalid token."}
    jsonBytes, err := utils.StructToJson(errorMessage); if err != nil {
      fmt.Print(err)
    }
    w.Header().Set("Content-Type", "application/json")
    w.Write(jsonBytes)

		return
  }

	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Username: loginData.User.Email,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		w.WriteHeader(http.StatusInternalServerError)
		return
  }

  // Create the user model to return
  var userToReturn ReturnLoginUser
  userToReturn.User.Email = loginData.User.Email
  userToReturn.User.Token = tokenString
  
  jsonBytes, err := utils.StructToJson(userToReturn); if err != nil {
    fmt.Print(err)
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write(jsonBytes)
  return
	
}
