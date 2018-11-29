package main

import (
	"fmt"
	"log"
	"time"

	"encoding/json"

	"github.com/mhannig/gitbase"
)

type Shader struct {
	Id          uint64 `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Author      string `json:"author"`

	CreatedAt time.Time `json:"created_at"`

	Token string `json:"token"`

	archive *gitbase.Archive
}

type ShaderRepository struct {
	Repository *gitbase.Repository
	Collection *gitbase.Collection
}

func NewShaderRepository(path string) (*ShaderRepository, error) {

	// Initialize gitbase
	repository, err := gitbase.NewRepository(path)
	if err != nil {
		return nil, err
	}

	collection, err := repository.Use("shaders")

	repo := &ShaderRepository{
		Repository: repository,
		Collection: collection,
	}

	return repo, nil
}

func (self *ShaderRepository) Create(shader *Shader) (uint64, error) {
	// Serialize data
	data, err := json.Marshal(shader)
	if err != nil {
		return 0, err
	}

	// Create new id in collection
	archive, err := self.Collection.NextArchive("created shader archive")
	if err != nil {
		return 0, err
	}

	err = archive.Put("meta.json", data, "created shader metadata")

	return archive.Id, err
}

func (self *ShaderRepository) List() ([]*Shader, error) {
	// Get list of shader archives
	archives, err := self.Collection.Archives()
	if err != nil {
		return nil, err
	}

	shaders := []*Shader{}
	for _, archive := range archives {
		// Get shader meta
		metajson, err := archive.Fetch("meta.json")
		if err != nil {
			log.Println("Could not get meta.json from archive:", archive.Id)
			continue
		}

		shader, err := UnmarshalShader(metajson)
		if err != nil {
			log.Println("Could not deserialize meta:", err)
			continue
		}

		shaders = append(shaders, shader)
	}

	return shaders, nil
}

func (self *ShaderRepository) Find(id uint64) (*Shader, error) {
	archive, err := self.Collection.Find(id)
	if err != nil {
		return nil, err
	}

	metajson, err := archive.Fetch("meta.json")
	if err != nil {
		return nil, err
	}

	shader, err := UnmarshalShader(metajson)
	if err != nil {
		shader.archive = archive
	}

	return shader, err
}

func (self *ShaderRepository) Update(id uint64, shader *Shader) error {
	current, err := self.Find(id)
	if err != nil {
		return err
	}

	return current.Update(shader)
}

//
// Helper
//
func UnmarshalShader(metajson []byte) (*Shader, error) {
	// Deserialize meta
	var shader *Shader
	err := json.Unmarshal(metajson, &shader)
	return shader, err
}

func MarshalShader(shader *Shader) ([]byte, error) {
	return json.Marshal(shader)
}

func (self *Shader) Update(next *Shader) error {
	if self.archive == nil {
		return fmt.Errorf("Shader is not persisted")
	}

	payload, err := MarshalShader(next)
	if err != nil {
		return err
	}

	return self.archive.Put("meta.json", payload, "updated shader")
}

func (self *Shader) UpdateProgram(program []byte) error {
	if self.archive == nil {
		return fmt.Errorf("Shader is not persisted")
	}

	return self.archive.Put("program", program, "updated shader program")
}

func (self *Shader) Program() []byte {
	if self.archive == nil {
		return []byte{} // not persisted, nothing to do
	}

	program, _ := self.archive.Fetch("program")

	return program
}

func (self *Shader) Destroy() error {
	if self.archive == nil {
		return fmt.Errorf("Shader is not persisted")
	}

	return self.archive.Destroy("removed shader")
}
