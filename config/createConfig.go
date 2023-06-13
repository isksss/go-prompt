package config

import (
	"os"
	"path/filepath"
)

const (
	appName    = "go-prompt"
	configFile = "config.json"
)

// return configFilePath, error
func MakeConfigFile() (string, error) {
	XDG_CONFIG_HOME := os.Getenv("XDG_CONFIG_HOME")
	if XDG_CONFIG_HOME != "" {
		home := getHomeDir()
		XDG_CONFIG_HOME = filepath.Join(home, ".config")
	}
	conf := filepath.Join(XDG_CONFIG_HOME, appName)
	//conf ディレクトリを作成
	if _, err := os.Stat(conf); os.IsNotExist(err) {
		err := os.MkdirAll(conf, 0755)
		if err != nil {
			return "", err
		}
	}
	//config.json ファイルを作成
	configFile := filepath.Join(conf, configFile)
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		file, err := os.Create(configFile)
		if err != nil {
			return "", err
		}
		defer file.Close()
	}

	return configFile, nil
}

func getHomeDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return home
}

func GetConfigFilePath() string {
	XDG_CONFIG_HOME := os.Getenv("XDG_CONFIG_HOME")
	if XDG_CONFIG_HOME != "" {
		home := getHomeDir()
		XDG_CONFIG_HOME = filepath.Join(home, ".config")
	}
	conf := filepath.Join(XDG_CONFIG_HOME, appName)
	configFile := filepath.Join(conf, configFile)
	return configFile
}
