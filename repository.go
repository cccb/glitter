package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

/*
 Files / Shader repository
*/

type Shader struct {
	Id          uint64 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`

	Author string `json:"author"`

	Token string `json:"token"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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

func (self *ShaderRepository) List() ([]*Shader, error) {
	shaders := []*Shader{}

	f, err := os.Open(self.basePath)
	if err != nil {
		return shaders, err
	}
	defer f.Close()

	items, err := f.Readdir(0)
	if err != nil {
		return shaders, err
	}

	for _, item := range items {
		if item.IsDir() == false {
			continue
		}

		shaderId, err := strconv.ParseUint(item.Name(), 10, 64)
		if err != nil {
			log.Println("Found non numeric entry in shader path.")
			log.Println("Please check if the repository is OK.")
			continue
		}

		shader, err := self.LoadShader(shaderId)
		if err != nil {
			log.Println("Error while loading shader:", err)
			continue
		}

		shaders = append(shaders, shader)
	}

	return shaders, nil
}

func (self *ShaderRepository) NextId() uint64 {
	return 1
}

func (self *ShaderRepository) GetRepositoryIdentifierFilename() string {
	return self.basePath + "/SHADER_REPOSITORY"
}

func (self *ShaderRepository) GetPath(id uint64) string {
	return fmt.Sprintf("%s/%d", self.basePath, id)
}

func (self *ShaderRepository) GetMetaFilename(id uint64) string {
	return self.GetPath(id) + "/meta.json"
}

func (self *ShaderRepository) GetProgramFilename(id uint64) string {
	return self.GetPath(id) + "/program.lua"
}

func (self *ShaderRepository) LoadShader(id uint64) (*Shader, error) {
	metaFilename := self.GetMetaFilename(id)
	data, err := ioutil.ReadFile(metaFilename)
	if err != nil {
		return nil, err
	}

	// Parse metadata
	var shader *Shader
	err = json.Unmarshal(data, &shader)
	if err != nil {
		return nil, err
	}

	// Assert a valid shaderid
	shader.Id = id

	return shader, nil
}

func NewShader(name string) *Shader {
	return nil
}
