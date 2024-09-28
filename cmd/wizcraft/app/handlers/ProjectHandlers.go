package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateProject creates a new Project
func CreateProject(c *gin.Context) {
	// TODO: Add create logic
	c.String(http.StatusOK, "Create Project")
}

// GetProject gets a Project by ID
func GetProject(c *gin.Context) {
	id := c.Param("id")
	// TODO: Add get logic
	c.String(http.StatusOK, "Get Project with ID: %s", id)
}

// UpdateProject updates an existing Project by ID
func UpdateProject(c *gin.Context) {
	id := c.Param("id")
	// TODO: Add update logic
	c.String(http.StatusOK, "Update Project with ID: %s", id)
}

// DeleteProject deletes an existing Project by ID
func DeleteProject(c *gin.Context) {
	id := c.Param("id")
	// TODO: Add delete logic
	c.String(http.StatusOK, "Delete Project with ID: %s", id)
}

// RegisterProjectRoutes registers the CRUD routes for Project
func RegisterProjectRoutes(router *gin.Engine) {
	router.POST("/project", CreateProject)
	router.GET("/project/:id", GetProject)
	router.PUT("/project/:id", UpdateProject)
	router.DELETE("/project/:id", DeleteProject)
}
