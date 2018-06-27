package main

import (
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

	repository *ShaderRepositroy
}

type ShaderRepositroy struct {
	Base       *gitbase.Gitbase
	Collection *gitbase.Collection
}

func NewShaderRepository(path string) (*ShaderRepository, err) {

	// Initialize gitbase
	base, err := gitbase.NewRepository(config.RepoPath)
	if err != nil {
		return nil, err
	}

	collection, err := base.Use("shaders")

	repository := &ShaderRepositroy{
		Base:       base,
		Collection: collection,
	}

	return repository, nil
}

func (self *ShaderRepositroy) Create(shader *Shader) error {
	// Serialize data
	data, err := json.Marshal(shader)
	if err != nil {
		return err
	}

	// Create new id in collection
	archive, err := self.Collection.NextArchive("created shader archive")
	if err != nil {
		return err
	}

	err = archive.Put("meta.json", data, "created shader metadata")
	return err
}

func (self *ShaderRepositroy) List() ([]*Shader, error) {
	// Get list of shader archives
	archives, err := self.Collection.Archives()
	if err != nil {
		return error
	}

	shaders := []*Shader{}
	for _, archive := range archives {
		// Get shader meta
		metajson, err := archive.Fetch("meta.json")
		if err != nil {
			log.Println("Could not get meta.json from archive:", archive.Id)
			continue
		}

		// Deserialize meta, add to result set
		var shader *Shader
		err = json.Unmarshal(metajson, &shader)
		if err != nil {
			log.Println("Error while deserializing meta.json:", err)
			continue
		}

		shaders = append(shaders, shader)
	}

	return shaders, nil
}

func (self *ShaderRepositroy) Find(uint64 id) (*Shader, error) {
	return nil, nil
}

//
// Active Shader
//
func (self *Shader) Update(meta *Shader) error {

	return nil
}

func (self *Shader) UpdateProgram(program []byte) error {

	return nil
}

func (self *Shader) Program() []byte {
	return []byte{}
}

func (self *Shader) Delete() error {
	return nil
}
