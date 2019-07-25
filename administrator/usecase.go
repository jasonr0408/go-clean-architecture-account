package administrator

import "account/models"

type Usecase interface {
	Login(hallID int, account, password string) (string, error)
	GetInfo(sid string) (*models.Administrator, error)
	Logout(sid string) error
	GetList(hallID int) ([]models.Administrator, error)
	Edit(sid string, administratorData *models.Administrator) error
}
