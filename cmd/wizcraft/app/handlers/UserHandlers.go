package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateUser creates a new User
func CreateUser(c *gin.Context) {
	// TODO: Add create logic
	c.String(http.StatusOK, "Create User")
}

// GetUser gets a User by ID
func GetUser(c *gin.Context) {
	id := c.Param("id")
	// TODO: Add get logic
	c.String(http.StatusOK, "Get User with ID: %s", id)
}

// UpdateUser updates an existing User by ID
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	// TODO: Add update logic
	c.String(http.StatusOK, "Update User with ID: %s", id)
}

// DeleteUser deletes an existing User by ID
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	// TODO: Add delete logic
	c.String(http.StatusOK, "Delete User with ID: %s", id)
}

// RegisterUserRoutes registers the CRUD routes for User
func RegisterUserRoutes(router *gin.Engine) {
	router.POST("/user", CreateUser)
	router.GET("/user/:id", GetUser)
	router.PUT("/user/:id", UpdateUser)
	router.DELETE("/user/:id", DeleteUser)
}
