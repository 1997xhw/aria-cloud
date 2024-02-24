package Services

import (
	"aria-cloud/Models"
	"aria-cloud/util"
)

const (
	pwd_salt = "*#890"
)

func LoginHandler(username, password string) (bool, error) {
	encPassword := util.Sha1([]byte(password + pwd_salt))
	user, err := Models.AuthenticateUser(username, encPassword)
	if !user || err != nil {
		return false, err
	}
	return true, nil
}

func UpadteUserToken(username, token string) bool {
	return Models.UpdateUserToken(username, token)
}
