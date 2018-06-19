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

		shader, err := self.GetMeta(shaderId)
		if err != nil {
			log.Println("Error while loading shader:", err)
			continue
		}

		shaders = append(shaders, shader)
	}

	return shaders, nil
}

func (self *ShaderRepository) NextId() (uint64, error) {

	shaders, err := self.List()
	if err != nil {
		return 0, err
	}

	var currentId uint64
	for _, shader := range shaders {
		if shader.Id > currentId {
			currentId = shader.Id
		}
	}

	return currentId + 1, nil
}

func (self *ShaderRepository) GetRepositoryIdentifierFilename() string {
	return self.basePath + "/SHADER_REPOSITORY"
}

func (self *ShaderRepository) GetPath(id uint64) string {
	return fmt.Sprintf("%s/%d", self.basePath, id)
}

func (self *ShaderRepository) Add(shader *Shader) (uint64, error) {
	nextId, err := self.NextId()
	if err != nil {
		return 0, err
	}

	path := self.GetPath(nextId)

	if err := os.MkdirAll(path, 0755); err != nil {
		return 0, err
	}
	metaFilename := path + "/meta.json"

	// Set auto fields
	shader.Id = nextId
	shader.CreatedAt = time.Now()
	shader.UpdatedAt = time.Now()

	// Serialize meta
	data, err := json.Marshal(shader)
	if err != nil {
		return 0, err
	}

	if err := ioutil.WriteFile(metaFilename, data, 0644); err != nil {
		return 0, err
	}

	return nextId, nil
}

func (self *ShaderRepository) GetProgram(id uint64) (string, error) {
	programFilename := self.GetPath(id) + "/program.lua"
	data, err := ioutil.ReadFile(programFilename)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (self *ShaderRepository) UpdateProgram(id uint64, program string) error {
	programFilename := self.GetPath(id) + "/program.lua"
	if err := ioutil.WriteFile(programFilename, []byte(program), 0644); err != nil {
		return err
	}

	return nil
}

func (self *ShaderRepository) Delete(id uint64) error {
	return os.RemoveAll(self.GetPath(id))
}

func (self *ShaderRepository) GetMeta(id uint64) (*Shader, error) {
	path := self.GetPath(id)

	metaFilename := path + "/meta.json"
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

	shader.Id = id

	return shader, nil
}
