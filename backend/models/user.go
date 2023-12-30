package models

import (
	"context"
	"errors"
	"fmt"

	"github.com/gdscduzceuniversity/todo-app-1/db"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID       string `bson:"_id,omitempty" json:"id"`
	Username string `bson:"username,omitempty" json:"username"`
	Password []byte `bson:"password,omitempty" json:"password"`
}

type AuthStatus struct {
	IsAuthenticated bool   `json:"isAuthenticated"`
	Message         string `json:"message"`
	Id              string `json:"Id"`
}

type Session struct {
	ID       string `bson:"_id,omitempty" json:"id"`
	Username string `bson:"username,omitempty" json:"username"`
}

func CreateUser(user User) (err error) {

	user.ID = uuid.New().String()
	result, err := db.UserCollection.InsertOne(context.Background(), user)
	if err != nil {
		err = fmt.Errorf("error in InsertOne collection method: %w", err)
		return err
	}

	fmt.Println(result)
	return nil
}

func GetUserByUsername(username string) (user User, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("Error in GetUserByUsername")
		}
	}()

	err = db.UserCollection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			user = User{}
			return user, nil
		}
	}

	return user, err
}

func GetUserByID(id string) (user User, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("Error in GetUserByID")
		}
	}()

	err = db.UserCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			user = User{}
			return user, nil
		}
	}

	return user, err
}

func ValidateUser(c *gin.Context) AuthStatus {
	cookie, err := c.Cookie("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil {
		return AuthStatus{IsAuthenticated: false, Message: "Unauthorized", Id: "0"}
	}

	claims := token.Claims.(*jwt.StandardClaims)
	user, err := GetUserByID(claims.Issuer)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			user = User{}
			return AuthStatus{IsAuthenticated: false, Message: "User not found", Id: user.ID}
		}
	}

	return AuthStatus{IsAuthenticated: true, Message: "Authorized", Id: user.ID}
}
