package gateway

import "glassgreenhouse.io/plants-service/domain"

type StorePlantPort interface {
	Store(plant *domain.Plant) error
}
