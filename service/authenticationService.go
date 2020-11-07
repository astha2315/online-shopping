package service

import "online-shopping/repository"

type LoginService interface {
	LoginUser(userName string, password string) (bool, int)
}
type loginInformation struct {
	userName string
	password string
}

func StaticLoginService() LoginService {
	return &loginInformation{
		userName: "bikash.dulal@wesionary.team",
		password: "testing",
	}
}
func (info *loginInformation) LoginUser(userName string, password string) (bool, int) {

	userData, err := repository.NewUserRepository().GetUserByUsernameAndPassword(userName, password)

	if err != nil {
		return false, 0
	}

	if userData != nil {
		return true, int(userData.Id)
	} else {
		return false, 0
	}

}
