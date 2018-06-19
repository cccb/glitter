package main

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
	BasePath string
}

func NewShaderRepository(basePath string) (*ShaderRepository, error) {
	return nil, nil
}

func (self *ShaderRepository) NewShader(name string) error {
	return nil
}

func (self *ShaderRepository) NextId() uint64 {
	return 1
}

func (self *ShaderRepository) GetPath(id uint64) string {
	return self.BasePath + "/" + string(id)
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
