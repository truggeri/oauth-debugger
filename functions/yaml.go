package oauthdebugger

import (
	"fmt"
	"io/ioutil"
	"path"

	"gopkg.in/yaml.v2"
)

type UserInfo struct {
	Breed    string `yaml:"breed"`
	ImageUrl string `yaml:"image_url"`
	Name     string `yaml:"name"`
	Username string `yaml:"username"`
}

const LOCAL_USERS_FILE = "users.yaml"

func localUser(username string) UserInfo {
	users := loadLocalUsers(LOCAL_USERS_FILE)
	for _, u := range users {
		if u.Username == username {
			return u
		}
	}
	return UserInfo{}
}

func loadLocalUsers(file string) []UserInfo {
	filePath := path.Join(GCP_SOURCE_DIR, file)
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("cannot load file")
		return []UserInfo{}
	}

	var users []UserInfo
	err = yaml.Unmarshal(yamlFile, &users)
	if err != nil {
		fmt.Printf("yaml issue")
		return []UserInfo{}
	}

	return users
}
