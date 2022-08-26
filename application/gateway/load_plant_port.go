package gateway

import "glassgreenhouse.io/plants-service/domain"

type LoadPlantPort interface {
	Find(id string) (*domain.Plant, error)
}
