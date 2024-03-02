package services

import (
	"aria-cloud/models"
	"aria-cloud/util"
)

const (
	pwd_salt = "*#890"
)

func LoginHandler(username, password string) (bool, error) {
	encPassword := util.Sha1([]byte(password + pwd_salt))
	user, err := models.AuthenticateUser(username, encPassword)
	if !user || err != nil {
		return false, err
	}
	return true, nil
}

func Register(username, password string) (bool, error) {
	encPassword := util.Sha1([]byte(password + pwd_salt))
	resp, err := models.AuthRegister(username, encPassword)
	if resp {
		return true, nil
	} else {
		return false, err
	}

}

func UpadteUserToken(username, token string) bool {
	return models.UpdateUserToken(username, token)
}

func GetTokenByUsername(username string) (string, error) {
	return models.GetTokenByUsername(username)
}
