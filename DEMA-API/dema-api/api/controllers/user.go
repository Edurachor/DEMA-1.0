package controllers

import (
	"apidema/api/auth"
	"apidema/api/database"
	"apidema/api/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {

	var User models.User

	//Verificando formulário
	c.Request.ParseForm()
	if !c.Request.Form.Has("name") || !c.Request.Form.Has("email") || !c.Request.Form.Has("password") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Preencha todos os dados!"})
		return
	}

	// Preenchendo usuário para enviar para banco de dados
	User.Name = c.Request.FormValue("name")
	User.Email = c.Request.FormValue("email")
	if err := User.HashPassword(c.Request.FormValue("password")); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Criando usuário com os dados preenchidos no banco de dados
	if err := database.Db.Create(&User); err.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
		return
	}

	// Retornando resultado para a requisição
	c.JSON(http.StatusOK, User)
}

func Login(c *gin.Context) {

	var request auth.Authentication
	var user models.User

	// Passando informações do contexto para o modelo de requisição
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	// Buscando email da requisição no banco de dados e, caso haja, passando para o modelo do usuário
	if record := database.Db.Where("email = ?", request.Email).First(&user); record.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		c.Abort()
		return
	}

	// Testando senha
	if checkPassword := user.CheckPassword(request.Password); checkPassword != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "senha incorreta"})
		c.Abort()
		return
	}

	// Gerando o token de acesso (JWT)
	token, err := auth.GenerateToken(fmt.Sprint(user.ID), user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, token)
	c.Redirect(http.StatusOK, "https://youtube.com")
}

func FindAllUsers(c *gin.Context) {

	var users []models.User
	database.Db.Find(&users)
	c.JSON(http.StatusOK, users)
}
