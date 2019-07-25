package administrator

import "account/models"

type Repository interface {
	Get(hallID int, account string) (models.Administrator, error)
	GetListByHall(hallID int) ([]models.Administrator, error)
	Update(administratorData *models.Administrator)
}

type SidRepository interface {
	StoreSid(sid string, administratorData string) error
	GetAdministratorDataBySid(sid string) (string, error)
	DeleteSid(sid string) error
}
