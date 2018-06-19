package main

import (
	"os"
	"testing"
)

func TestRepositoryInitialization(t *testing.T) {
	testRepoPath := os.TempDir() + "/shader-test-repo"
	os.MkdirAll(testRepoPath, 0755)
	defer os.RemoveAll(testRepoPath)

	repo := NewShaderRepository(testRepoPath)

	if repo.IsRepository() != false {
		t.Error("Repository should be uninitialized")
	}

	if err := repo.CanInitialize(); err != nil {
		t.Error("Repository should be initializable:", err)
	}

	if err := repo.Initialize(); err != nil {
		t.Error("Repository initiailization failed:", err)
	}

	t.Log(repo)

}

func TestRepositorySetup(t *testing.T) {
	testRepoPath := os.TempDir() + "/shader-test-repo"
	os.MkdirAll(testRepoPath, 0755)
	defer os.RemoveAll(testRepoPath)

	repo := NewShaderRepository(testRepoPath)
	err := repo.Setup()
	if err != nil {
		t.Error("Repository setup failed:", err)
	}

	// This should work again
	err = repo.Setup()
	if err != nil {
		t.Error("Expected no error, because is initialized. Error:", err)
	}
}

func TestListShader(t *testing.T) {
	testRepoPath := os.TempDir() + "/shader-test-repo"
	os.MkdirAll(testRepoPath, 0755)
	defer os.RemoveAll(testRepoPath)

	repo := NewShaderRepository(testRepoPath)
	if err := repo.Setup(); err != nil {
		t.Error("Repository setup failed:", err)
	}

	shaders, err := repo.List()
	if err != nil {
		t.Error(err)
	}

	if len(shaders) != 0 {
		t.Error("Expected repo to be empty")
	}

	// Create a shader
	shader := &Shader{
		Name:   "bunt3",
		Author: "Ben Utzer",
	}

	_, err = repo.Add(shader)
	if err != nil {
		t.Error(err)
	}

	shaders, err = repo.List()
	if err != nil {
		t.Error(err)
	}

	if len(shaders) != 1 {
		t.Error("Expected shaders list to of length 1")
	}
}

func TestAddGetDeleteShader(t *testing.T) {
	testRepoPath := os.TempDir() + "/shader-test-repo"
	os.MkdirAll(testRepoPath, 0755)
	defer os.RemoveAll(testRepoPath)

	repo := NewShaderRepository(testRepoPath)
	if err := repo.Setup(); err != nil {
		t.Error("Repository setup failed:", err)
	}

	shader := &Shader{
		Name:   "bunt3",
		Author: "Ben Utzer",
	}

	shaderId, err := repo.Add(shader)
	if err != nil {
		t.Error(err)
	}

	if shaderId != 1 {
		t.Error("Expected shaderId to be 1")
	}

	shader.Name = "bunt4"
	shaderId, err = repo.Add(shader)
	if err != nil {
		t.Log(err)
	}

	if shaderId != 2 {
		t.Error("Expected shaderId to be 2")
	}

	// Retrieve shaders
	shader, err = repo.GetMeta(1)
	if err != nil {
		t.Error(err)
	}

	if shader.Name != "bunt3" {
		t.Error("Shader#1.Name should be bunt3")
	}

	_, err = repo.GetMeta(3)
	if err == nil {
		t.Error("Unknown Id should yield an error")
	}

	if err = repo.Delete(1); err != nil {
		t.Error(err)
	}
	_, err = repo.GetMeta(1)
	if err == nil {
		t.Error("Shader#1 should have been deleted!")
	}
}

func TestUpdateGetProgram(t *testing.T) {
	testRepoPath := os.TempDir() + "/shader-test-repo"
	os.MkdirAll(testRepoPath, 0755)
	defer os.RemoveAll(testRepoPath)

	repo := NewShaderRepository(testRepoPath)
	if err := repo.Setup(); err != nil {
		t.Error("Repository setup failed:", err)
	}

	shader := &Shader{
		Name:   "bunt3",
		Author: "Ben Utzer",
	}

	shaderId, err := repo.Add(shader)
	if err != nil {
		t.Error(err)
	}

	_, err = repo.GetProgram(shaderId)
	if err == nil {
		t.Error("A program should not have been found for this shader")
	}

	program := "This\nIs\nMy\nShader\nCode!"
	err = repo.UpdateProgram(shaderId, program)
	if err != nil {
		t.Error(err)
	}

	retrieved, err := repo.GetProgram(shaderId)
	if err != nil {
		t.Error(err)
	}

	if retrieved != program {
		t.Error("Expected", program, " == ", retrieved)
	}
}
