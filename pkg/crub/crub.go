package crub

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

// This package provides CRUD operations that used by wizcraft or wand cli tools.
// Create, Read, Update, Delete for a given resource.

// For example, the following code snippet shows how to create a new resource:
// $ wand crud -r User -f handlers

// Resource represents the resource metadata (can include fields, etc.)
type Resource struct {
	Name string // Name of the resource (e.g., "User", "Product")
}

// Generate generates CRUD operation template files for the given resource
func Generate(resource Resource, path string) error {
	// Template for the CRUD operations using Gin-Gonic
	const crudTemplate = `package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Create{{.Name}} creates a new {{.Name}}
func Create{{.Name}}(c *gin.Context) {
	// TODO: Add create logic
	c.String(http.StatusOK, "Create {{.Name}}")
}

// Get{{.Name}} gets a {{.Name}} by ID
func Get{{.Name}}(c *gin.Context) {
	id := c.Param("id")
	// TODO: Add get logic
	c.String(http.StatusOK, "Get {{.Name}} with ID: %s", id)
}

// Update{{.Name}} updates an existing {{.Name}} by ID
func Update{{.Name}}(c *gin.Context) {
	id := c.Param("id")
	// TODO: Add update logic
	c.String(http.StatusOK, "Update {{.Name}} with ID: %s", id)
}

// Delete{{.Name}} deletes an existing {{.Name}} by ID
func Delete{{.Name}}(c *gin.Context) {
	id := c.Param("id")
	// TODO: Add delete logic
	c.String(http.StatusOK, "Delete {{.Name}} with ID: %s", id)
}

// Register{{.Name}}Routes registers the CRUD routes for {{.Name}}
func Register{{.Name}}Routes(router *gin.Engine) {
	router.POST("/{{.Name | ToLower}}", Create{{.Name}})
	router.GET("/{{.Name | ToLower}}/:id", Get{{.Name}})
	router.PUT("/{{.Name | ToLower}}/:id", Update{{.Name}})
	router.DELETE("/{{.Name | ToLower}}/:id", Delete{{.Name}})
}
`

	// Create the directory if it doesn't exist
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}

	// Create the file to save the template
	filePath := filepath.Join(path, fmt.Sprintf("%sHandlers.go", resource.Name))
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	// Parse the template and execute it with the resource data
	tmpl, err := template.New("crud").Funcs(template.FuncMap{
		"ToLower": func(s string) string {
			return string([]rune(s)[0]|' ') + s[1:]
		},
	}).Parse(crudTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse template: %v", err)
	}

	err = tmpl.Execute(file, resource)
	if err != nil {
		return fmt.Errorf("failed to execute template: %v", err)
	}

	fmt.Printf("Generated CRUD handlers for resource: %s\n", resource.Name)
	return nil
}

func Run(resourceName string, filePath string) {

	// Check if both flags are provided
	if resourceName == "" || filePath == "" {
		fmt.Println("Both -r (resource) and -f (file path) flags are required.")
		flag.Usage()
		return
	}

	// Generate CRUD handlers
	resource := Resource{Name: resourceName}
	err := Generate(resource, filePath)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
