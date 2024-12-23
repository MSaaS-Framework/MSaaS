package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"MSaaS-Framework/MSaaS/pkg/object"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateProject(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a new project object
	project := object.Project{
		General: &object.GeneralSpec{
			Name:        "Test Project",
			Type:        "Test Type",
			Description: "Test Description",
		},
	}

	// Marshal the project object to JSON
	projectJSON, err := json.Marshal(project)
	assert.NoError(t, err)

	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodPost, "http://localhost:8080/project", bytes.NewBuffer(projectJSON))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	// Create a new HTTP client
	client := &http.Client{}

	// Perform the request
	resp, err := client.Do(req)
	assert.NoError(t, err)
	defer resp.Body.Close()

	// Check the response
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	bodyBytes, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	assert.Equal(t, "Create Project", string(bodyBytes))
}
