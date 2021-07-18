package config

// for repo tool
const (
	repoConfigName = "config"
)

type RepoConfig struct {
	DBName string `yaml:"db_name", omitempty`
}
