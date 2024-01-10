package main

import (
	"crypto/rsa"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var privateKey *rsa.PrivateKey
var publicKey *rsa.PublicKey
var refreshTokenSecretKey = "your-refresh-secret-key"

var accessExpDuration = time.Hour * 24       // Access token expiration time (1 day)
var refreshExpDuration = time.Hour * 24 * 30 // Refresh token expiration time (30 days)

// var secretKey = "your-secret-key"

type RefreshTokenClaims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

var users = map[string]User{
	"user1": {ID: 1, Username: "user1", Password: "password1"},
	"user2": {ID: 2, Username: "user2", Password: "password2"},
}

func init() {
	// Load RSA private key
	privateBytes, err := os.ReadFile("private.pem")
	if err != nil {
		panic("Failed to load private key")
	}
	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		panic("Failed to parse private key")
	}

	// Load RSA public key
	publicBytes, err := os.ReadFile("public.pem")
	if err != nil {
		panic("Failed to load public key")
	}
	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		panic("Failed to parse public key")
	}
}

func main() {
	e := echo.New()

	e.POST("/login", loginHandler)
	e.GET("/private", privateHandler, middleware)
	e.POST("/refresh", refreshHandler)

	e.Start(":8080")
}

func loginHandler(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}

	// Check if the username and password match
	storedUser, ok := users[u.Username]
	if !ok || storedUser.Password != u.Password {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid credentials"})
	}

	// Generate a JWT access token
	accessToken := generateAccessToken(storedUser)

	// Generate a JWT refresh token
	refreshToken := generateRefreshToken(storedUser)

	// Return both tokens
	return c.JSON(http.StatusOK, map[string]string{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

func generateAccessToken(user User) string {
	token := jwt.New(jwt.SigningMethodRS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(accessExpDuration).Unix() // Access token expiration time

	// Sign the token with the RSA private key
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		// Handle error here
		return ""
	}

	return tokenString
}

func generateRefreshToken(user User) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(refreshExpDuration).Unix() // Refresh token expiration time

	// Sign the token with the refresh token secret key
	tokenString, err := token.SignedString([]byte(refreshTokenSecretKey))
	if err != nil {
		// Handle error here
		return ""
	}

	return tokenString
}

// func generateJWTToken(user User) string {
// 	token := jwt.New(jwt.SigningMethodHS256)
// 	claims := token.Claims.(jwt.MapClaims)
// 	claims["id"] = user.ID
// 	claims["username"] = user.Username
// 	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expiration time (1 day)

// 	// Sign the token with a secret key
// 	tokenString, err := token.SignedString([]byte(secretKey))
// 	if err != nil {
// 		// Handle error here
// 		return ""
// 	}

// 	return tokenString
// }

func middleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        tokenString := c.Request().Header.Get("Authorization")
        if tokenString == "" {
            return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Authorization token missing"})
        }

        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return publicKey, nil // Use the RSA public key for verification
        })

        if err != nil || !token.Valid {
            return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid token"})
        }

        if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
            // You can access the claims here, such as claims["id"] or claims["username"]
            c.Set("user", claims)
            return next(c)
        }

        return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid token"})
    }
}

func privateHandler(c echo.Context) error {
	user := c.Get("user").(jwt.MapClaims)
	userID := int(user["id"].(float64))
	username := user["username"].(string)

	// You can access the user's ID and username here
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "This is a protected route",
		"user_id":  userID,
		"username": username,
	})
}

func refreshHandler(c echo.Context) error {
	refreshToken := c.FormValue("refresh_token")
	if refreshToken == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Refresh token missing"})
	}

	// Verify the refresh token
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(refreshTokenSecretKey), nil
	})

	if err != nil || !token.Valid {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid refresh token"})
	}

	// Generate a new access token
	userID := int(claims["user_id"].(float64))
	storedUser := getUserByID(userID)
	accessToken := generateAccessToken(storedUser)

	// Return the new access token
	return c.JSON(http.StatusOK, map[string]string{"access_token": accessToken})
}

func getUserByID(userID int) User {
	for _, user := range users {
		if user.ID == userID {
			return user
		}
	}
	return User{}
}
