package main

import (
	"fmt"
	"log"
	"os"
)

/*
 Files / Shader repository
*/

type ShaderMeta struct {
	Name        string `json:"name"`
	Description string `json:"description"`

	Author string `json:"author"`

	Token string `json:"token"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type Shader struct {
	Meta    *ShaderMeta
	Program string
}

type ShaderRepository struct {
	basePath string
}

func NewShaderRepository(basePath string) *ShaderRepository {
	repo := &ShaderRepository{
		basePath: basePath,
	}

	log.Println("Using shader repository:", basePath)

	return repo
}

func (self *ShaderRepository) Setup() error {
	if self.IsRepository() {
		return nil // Nothing to do here
	}

	if err := self.CanInitialize(); err != nil {
		return err
	}

	if err := self.Initialize(); err != nil {
		return err
	}

	return nil
}

func (self *ShaderRepository) IsRepository() bool {
	// Check if repository identifier exists
	_, err := os.Stat(self.GetRepositoryIdentifierFilename())
	if err != nil {
		return false
	}

	return true
}

func (self *ShaderRepository) CanInitialize() error {
	// Check if path exists and is empty
	f, err := os.Open(self.basePath)
	if err != nil {
		return err
	}
	defer f.Close()

	items, err := f.Readdir(0)
	if err != nil {
		return err
	}

	if len(items) != 0 {
		return fmt.Errorf("Path not empty")
	}

	return nil
}

func (self *ShaderRepository) Initialize() error {
	f, err := os.OpenFile(
		self.GetRepositoryIdentifierFilename(),
		os.O_RDWR|os.O_CREATE,
		0644)

	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write([]byte{})
	return err
}

func (self *ShaderRepository) NewShader(name string) error {
	return nil
}

func (self *ShaderRepository) NextId() uint64 {
	return 1
}

func (self *ShaderRepository) GetRepositoryIdentifierFilename() string {
	return self.basePath + "/SHADER_REPOSITORY"
}

func (self *ShaderRepository) GetPath(id uint64) string {
	return self.basePath + "/" + string(id)
}

func (self *ShaderRepository) GetMetaFilename(id uint64) string {
	return self.GetPath(id) + "/meta.json"
}

func (self *ShaderRepository) GetProgramFilename(id uint64) string {
	return self.GetPath(id) + "/program.lua"
}

func (self *Shader) Save(path string) error {
	return nil
}
