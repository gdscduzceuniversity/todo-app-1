package handlers

import (
	"net/http"
	"time"

	"github.com/gdscduzceuniversity/todo-app-1/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"

	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "secret" //ENV file TODO

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Register(c *gin.Context) {
	request := &RegisterRequest{}
	if err := c.Bind(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Bad Request"})
		return
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(request.Password), 14)

	user := models.User{
		Username: request.Username,
		Password: password,
	}

	userExists, err := models.GetUserByUsername(request.Username)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error GetUserByUsername"})
		return
	}

	if userExists.ID != "" {
		// json message with gin
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User already exists",
		})
		return
	}

	if err = models.CreateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error CreateUser"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})

	return
}

func Login(c *gin.Context) {
	request := &LoginRequest{}
	if err := c.ShouldBindJSON(request); err != nil {
		c.JSON(400, gin.H{"message": "Bad Request"})
		return
	}

	auth := models.ValidateUser(c)
	if auth.IsAuthenticated {
		c.JSON(401, gin.H{
			"message": "User already logged in",
			"userId":  auth.Id,
		})
		return
	}

	user, err := models.GetUserByUsername(request.Username)
	if err != nil {
		c.JSON(500, gin.H{"message": "Internal Server Error"})
		return
	}

	if user.ID == "" {
		c.JSON(500, gin.H{"message": "User not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		c.JSON(401, gin.H{"message": "Incorrect password"})
		return
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    user.ID,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := claims.SignedString([]byte(SecretKey))
	if err != nil {
		c.JSON(500, gin.H{"message": "Cannot login"})
		return
	}

	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
	}

	c.SetCookie(cookie.Name, cookie.Value, cookie.Expires.Year(), cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)

	c.JSON(200, gin.H{"message": "success"})
}

func User(c *gin.Context) {
	cookie, err := c.Cookie("jwt")
	if err != nil {
		c.JSON(401, gin.H{"AuthStatus": "Unauthenticated"})
		return
	}

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		c.JSON(401, gin.H{"AuthStatus": "Unauthenticated"})
		return
	}

	claims := token.Claims.(*jwt.StandardClaims)

	user, err := models.GetUserByID(claims.Issuer)

	if err != nil {
		c.JSON(404, gin.H{
			"message": "User not found",
			"userID":  user.ID,
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
		"user":    user,
	})
}

func Logout(c *gin.Context) {
	if jwtCookie, err := c.Cookie("jwt"); err != nil || jwtCookie == "" {
		c.JSON(401, gin.H{"message": "User not logged in"})
		return
	}

	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	}

	c.SetCookie(cookie.Name, cookie.Value, cookie.Expires.Year(), cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)

	c.JSON(200, gin.H{"message": "success"})
}
