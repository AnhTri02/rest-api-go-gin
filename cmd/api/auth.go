package main

import (
	"net/http"
	"rest-api-in-gin/internal/database"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)
//register handler


type registerRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
	Name     string `json:"name" binding:"required,min=2"`
}

//handlerfor POST /auth/register

func (app *application) registerUser(c *gin.Context){
	var register registerRequest

	if err := c.ShouldBindJSON(&register); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword , err := bcrypt.GenerateFromPassword([]byte(register.Password), bcrypt.DefaultCost)

	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Something went wrong"})
	}

	register.Password = string(hashedPassword)
	user := database.User{
		Email: register.Email,
		Password: register.Password,
		Name: register.Name,
	}

	err = app.models.Users.Insert(&user)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
	}

	c.JSON(http.StatusCreated, user)
}