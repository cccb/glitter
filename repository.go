package main

import (
	"log"
	"time"

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
	Base *gitbase.Gitbase
}

func NewShaderRepository(path string) (*ShaderRepository, err) {

	// Initialize gitbase
	base, err := gitbase.NewRepository(config.RepoPath)
	if err != nil {
		return nil, err
	}

	repository := &ShaderRepositroy{
		Base: base,
	}

	return repository, nil
}

func (self *ShaderRepositroy) Create(shader *Shader) error {

	return nil
}

func (self *ShaderRepositroy) List() ([]*Shader, error) {
	return nil, nil
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
