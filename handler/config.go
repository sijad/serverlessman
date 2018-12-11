package handler

// Config defines config file structure
type Config struct {
	Fields []Filed `json:"fields"`
	Output struct {
		Format string `json:"format"`
		Path   string `json:"path"`
	} `json:"output"`
}

var configs map[string]Config

// InitConfigs init provider configs
func InitConfigs() {
	c, err := provider.GetRepoConfigs()
	if err != nil {
		panic("can not get configs from provider: " + err.Error())
	}

	configs = c
}
