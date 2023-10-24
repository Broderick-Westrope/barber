package file

type Snippet struct {
	Path string   `yaml:"path"`
	Desc string   `yaml:"description"`
	Tags []string `yaml:"tags"`
}

func (s Snippet) FilterValue() string {
	return s.Path
}
func (s Snippet) Title() string {
	return s.Path
}
func (s Snippet) Description() string {
	return ""
}
