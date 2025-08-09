package models

type Grind struct {
	Name    string `yaml:"name"`
	Roaster string `yaml:"roaster"`
	Method  string `yaml:"method"`
	Dose    string `yaml:"dose"`
	Grinder string `yaml:"grinder"`
	Setting string `yaml:"setting"`
}
