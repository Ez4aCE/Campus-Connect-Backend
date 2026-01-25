package handlers

import (
	
	"net/http"

	"campus-connect-backend/internal/db"
	"campus-connect-backend/internal/models"
	"campus-connect-backend/internal/services"

	"github.com/gin-gonic/gin"
)

type RegisterRequest struct{
	Name      string   `json:"name" binding:"required"`
	Email     string   `json:"email" binding:"required,email"`
	Password  string   `json:"password" binding:"required,min=6"`
}

func Register(c *gin.Context){
	var req RegisterRequest

	if err:=c.ShouldBindJSON(&req); err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	hash, err:= services.HashPassword(req.Password)
	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":"password hashing failed"})
		return
	}

	user:=models.User{
		Name: req.Name,
		Email:  req.Email,
		PasswordHash: hash,
		Role: "participant",
	}

	if err:=db.DB.Create(&user).Error; err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"email already exists"})
		return
	}
	c.JSON(http.StatusCreated,gin.H{"message":"user registered successfully"})
}


type LoginRequest struct{
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
}

func Login(c *gin.Context){
	var req LoginRequest

	if err:=c.ShouldBindJSON(&req); err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}

	var user models.User

	if err:=db.DB.Where("email = ?", req.Email).First(&user).Error;err!=nil{
		c.JSON(http.StatusUnauthorized,gin.H{"error":"Invalid credentials"})
		return
	}
	

	if !services.CheckPassword(user.PasswordHash,req.Password){
		c.JSON(http.StatusUnauthorized,gin.H{"error":"Invalid credentials"})
		return
	}

	token, err:= services.GenerateJWT(user.ID.String(),user.Role)

	if err!=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":"token generation failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token":token})
}