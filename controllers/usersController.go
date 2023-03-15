package controllers

import (
	"os"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/ismetbayandur/goapi/initializers"
	"github.com/ismetbayandur/goapi/models"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)


func SignUp(c *gin.Context){

	var body struct {
		Email string
		Password string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : "failed to read body",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password),10)

	if err != nil {
                c.JSON(http.StatusBadRequest, gin.H{
                        "error" : "failed to hash password",
                })
                return
        }

	user := models.User{Email:body.Email, Password : string(hash)}
	result := initializers.DB.Create(&user)
	if result.Error != nil {
                c.JSON(http.StatusBadRequest, gin.H{
                        "error" : "failed to create user",
                })
                return
        }

	c.JSON(http.StatusOK, gin.H{})
}

func Login(c *gin.Context){
	var body struct {
                Email string
                Password string
        }
	if c.Bind(&body) != nil {
                c.JSON(http.StatusBadRequest, gin.H{
                        "error" : "failed to read body",
                })
                return
        }

	var user models.User
	initializers.DB.First(&user , "email = ?",body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
                        "error" : "invalid email or password",
                })
                return
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(body.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
                        "error" : "invalid email or password",
                })
                return
        }

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"aub": user.ID,
		"exp":time.Now().Add(time.Hour *24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

        if err != nil {
                c.JSON(http.StatusBadRequest, gin.H{
                        "error" : "failed to create token",
			"msg"   : err,
                })
                return
        }

	c.JSON(http.StatusOK, gin.H{
		"token" : tokenString,
	})

}
