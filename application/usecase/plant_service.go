package usecase

import (
	"errors"

	"glassgreenhouse.io/plants-service/application/gateway"
	"glassgreenhouse.io/plants-service/domain"
)

var (
	ErrRedirectNotFound = errors.New("Redirect Not Found")
	ErrRedirectInvalid  = errors.New("Redirect Invalid")
)

type PlantService struct {
	LoadPlantPort  gateway.LoadPlantPort
	StorePlantPort gateway.StorePlantPort
}

func NewPlantService(lpp gateway.LoadPlantPort, spp gateway.StorePlantPort) *PlantService {
	return &PlantService{
		lpp,
		spp,
	}
}

func (ps *PlantService) Find(id string) (*domain.Plant, error) {
	return ps.LoadPlantPort.Find(id)
}

func (ps *PlantService) Store(plant *domain.Plant) error {
	// if err := validate.Validate(redirect); err != nil {
	// 	return errs.Wrap(ErrRedirectInvalid, "service.Redirect.Store")
	// }
	// redirect.Code = shortid.MustGenerate()
	// redirect.CreatedAt = time.Now().UTC().Unix()
	return ps.StorePlantPort.Store(plant)
}
