package handler

import (
	"fmt"
	"os"
	"time"

	"exercise-3/repo"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func InitRoutes(engine *gin.Engine, db *gorm.DB) {
	engine.GET("/ping", pingHandler)
	groupRouterNote := engine.Group("/note")
	groupRouterUser := engine.Group("/user")

	groupRouterNote.GET("/:id", func(c *gin.Context) {
		noteRepository := &repo.NoteRepoImpl{
			DB: db,
		}
		result, err := NoteGet(c, noteRepository)
		simpleReturnHandler(c, err, result)
	})
	groupRouterNote.POST("", func(c *gin.Context) {
		// 1. Repo
		repo := &repo.NoteRepoImpl{
			DB: db,
		}
		// 2. Create note
		result, err := NoteCreate(c, repo)
		// 3. Handle result & err
		simpleReturnHandler(c, err, result)
	})
	groupRouterNote.PUT("/:id", func(c *gin.Context) {
		repo := &repo.NoteRepoImpl{
			DB: db,
		}
		err := NoteUpdate(c, repo)
		simpleReturnHandler(c, err, nil)
	})
	groupRouterNote.DELETE("/:id", func(c *gin.Context) {
		repo := &repo.NoteRepoImpl{
			DB: db,
		}
		err := NoteDelete(c, repo)
		simpleReturnHandler(c, err, nil)
	})

	groupRouterUser.POST("/login", func(c *gin.Context) {
		// 1. Repo
		repo := &repo.UserRepoImpl{
			DB: db,
		}
		// 2. Login user
		result, err := UserLogin(c, repo)
		// 3. Handle result & err
		simpleReturnHandler(c, err, result)
	})

	// 1. Authentication // Identity
	// 2. Lam logger/tracking
	// 3. Recovery
	// 4. Add nhieu cai middleware va no chay tuan tu
	groupRouterUser.Use(userMiddleware)
	{
		groupRouterUser.GET("/list-user", func(c *gin.Context) {
			// 1. Repo
			repo := &repo.UserRepoImpl{
				DB: db,
			}
			// 2. List user
			result, err := UserList(c, repo)
			// 3. Handle result & err
			simpleReturnHandler(c, err, result)
		})

		groupRouterUser.GET("/", func(c *gin.Context) {
			repo := &repo.UserRepoImpl{
				DB: db,
			}
			result, err := UserGet(c, repo)
			simpleReturnHandler(c, err, result)
		})

		groupRouterUser.POST("/", func(c *gin.Context) {
			repo := &repo.UserRepoImpl{
				DB: db,
			}
			err := UserCreate(c, repo)
			simpleReturnHandler(c, err, nil)
		})

		groupRouterUser.PUT("/", func(c *gin.Context) {
			fmt.Print("PUT")
			// 3. Handle result & err
			simpleReturnHandler(c, nil, nil)
		})

		groupRouterUser.DELETE("/:id", func(c *gin.Context) {
			repo := &repo.UserRepoImpl{
				DB: db,
			}
			err := UserDelete(c, repo)
			simpleReturnHandler(c, err, nil)
		})
	}
}

func simpleReturnHandler(c *gin.Context, err error, result interface{}) {
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, result)
}

func userMiddleware(c *gin.Context) {
	secretKey := os.Getenv("SECRET_KEY")
	mySigningKey := []byte(secretKey)

	// type MyCustomClaims struct {
	// 	User model.User `json:"user"`
	// 	jwt.StandardClaims
	// }

	tokenString := c.GetHeader("token")

	// check trong DB co token do khong // TODO?
	repo := &repo.UserRepoImpl{
		DB: db,
	}
	repo.DB

	// sample token is expired.  override time so it parses as valid
	// token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		c.Abort()
	}

	// check expired time
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		expired, err := time.Parse(time.RFC3339, claims["expired"].(string))

		if err != nil {
			c.Abort()
		}
		if expired.Before(time.Now()) {
			c.Abort()
		}

		// fmt.Printf("%v", claims["username"])
	}

	// 1. Logic quan trong nhat la, xu ly ngung cai cai request
	// 2. Tao ra cac du lieu de set vao context cho cai handler dung
	// 3. Quyet dinh cho phep di tiep den cai middleware tiep theo hoac handler
	// if c.GetHeader("token") != "202cb962ac59075b964b07152d234b70" {
	// 	c.AbortWithStatus(400)
	// 	return
	// }
	c.Next()
}
