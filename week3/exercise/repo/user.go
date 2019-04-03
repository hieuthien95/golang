package repo

import (
	"exercise-3/helper"
	"exercise-3/model"
	"os"

	"math/rand"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

type UserRepo interface {
	Login(string, string) (*model.User, error)
	List(helper.Pagination) ([]model.User, error)

	Get(id float64) (*model.User, error)
	Create(model.User) error
	Update(model.User) error
	Delete(int) error
}

type UserRepoImpl struct {
	DB *gorm.DB
}

func (self *UserRepoImpl) Login(username, password string) (*model.User, error) {

	var userOutput model.User
	self.DB.Where("username = ? AND password = ?", username, password).First(&userOutput)

	secretKey := os.Getenv("SECRET_KEY")
	mySigningKey := []byte(secretKey)

	// type MyCustomClaims struct {
	// 	User model.User `json:"user"`
	// 	jwt.StandardClaims
	// }

	// // Create the Claims
	// claims := MyCustomClaims{
	// 	userOutput,
	// 	jwt.StandardClaims{
	// 		ExpiresAt: 15000,
	// 		Issuer:    "test",
	// 	},
	// }

	// token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// tokenStr, _ := token.SignedString(mySigningKey)

	expired := time.Now().Add(600 * time.Second)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":    userOutput.Username,
		"userid":      userOutput.ID,
		"expired":     expired,
		"createdTime": time.Now(),
		"seed":        rand.Int(),
	})
	tokenStr, err := token.SignedString(mySigningKey)
	if err != nil {
	}

	userOutput.Token = tokenStr

	self.DB.Table("users").Update(&userOutput)

	return &userOutput, nil
}
func (self *UserRepoImpl) List(pagination helper.Pagination) ([]model.User, error) {
	users := []model.User{}
	offset := pagination.GetOffset()
	limit := pagination.GetLimit()
	err := self.DB.Offset(offset).
		Limit(limit).
		Find(&users).
		Error
	return users, err
}

func (self *UserRepoImpl) Get(id float64) (*model.User, error) {
	user := model.User{}
	err := self.DB.Where("id = ?", id).First(&user).Error
	return &user, err
}
func (self *UserRepoImpl) Create(user model.User) error {
	return self.DB.Create(&user).Error
}
func (self *UserRepoImpl) Update(user model.User) error {
	self.DB.Update(user)

	return nil
}
func (self *UserRepoImpl) Delete(id int) error {
	return self.DB.Where("id = ?", id).Delete(&model.User{}).Error
}
