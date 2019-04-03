package handler

import (
	"exercise-3/helper"
	"exercise-3/model"
	"exercise-3/repo"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context, userRepo repo.UserRepo) (*model.User, error) {
	var user model.User
	c.Bind(&user)
	return userRepo.Login(user.Username, user.Password)
}

func UserList(c *gin.Context, userRepo repo.UserRepo) ([]model.User, error) {
	var pagination helper.Pagination
	c.ShouldBindQuery(&pagination)

	return userRepo.List(pagination)
}

func UserGet(c *gin.Context, userRepo repo.UserRepo) (*model.User, error) {
	secretKey := os.Getenv("SECRET_KEY")
	mySigningKey := []byte(secretKey)

	tokenString := c.GetHeader("token")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		c.Abort()
	}

	// check expired time
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		expired, err := time.Parse(time.RFC3339, claims["expired"].(string))

		if err != nil {
		}
		if expired.Before(time.Now()) {
		}
	}

	id := claims["userid"].(float64)
	return userRepo.Get(id)
}

func UserCreate(c *gin.Context, userRepo repo.UserRepo) error {
	user := model.User{}
	if err := c.ShouldBind(&user); err != nil {
		return err
	}

	return userRepo.Create(user)
}

func UserDelete(c *gin.Context, userRepo repo.UserRepo) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return userRepo.Delete(id)
}
