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

}
