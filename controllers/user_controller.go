package controllers

import (
	"compartamos-backend/models"
	"compartamos-backend/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
    UserService services.UserService
}

func (ctrl *UserController) CreateUser(c *gin.Context) {
    var input models.User
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Campos inválidos"})
        return
    }

    if input.Age < 18 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "El usuario debe ser mayor de 18 años"})
        return
    }

    user, err := ctrl.UserService.CreateUser(input)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error en la creación de usuario"})
        return
    }

    c.JSON(http.StatusOK, user)
}

func (ctrl *UserController) GetUsers(c *gin.Context) {
    users, err := ctrl.UserService.GetUsers()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener usuarios"})
        return
    }

    c.JSON(http.StatusOK, users)
}

func (ctrl *UserController) GetUser(c *gin.Context) {
    id := c.Param("id")
    user, err := ctrl.UserService.GetUser(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
        return
    }

    c.JSON(http.StatusOK, user)
}

func (ctrl *UserController) UpdateUser(c *gin.Context) {
    id := c.Param("id")
    var input models.User
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Campos inválidos"})
        return
    }

    updatedUser, err := ctrl.UserService.UpdateUser(id, input)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar datos del usuario"})
        return
    }

    c.JSON(http.StatusOK, updatedUser)
}

func (ctrl *UserController) DeleteUser(c *gin.Context) {
    id := c.Param("id")


    err := ctrl.UserService.DeleteUser(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Usuario eliminado"})
}
