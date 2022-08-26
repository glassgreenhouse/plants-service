package usecase

import (
	"log"
	"reflect"
	"testing"

	"github.com/bxcodec/faker/v3"
	"glassgreenhouse.io/plants-service/application/gateway"
	"glassgreenhouse.io/plants-service/domain"
)

type PLantFakeData struct {
	Name      string `faker:"name"`
	CreatedAt int64  `faker:"unix_time"`
}

type findPlantFakeRepository struct {
	MockAddFn func(id string) (*domain.Plant, error)
}

type storePlantFakeRepository struct {
	MockAddFn func(*domain.Plant) error
}

func (fake *findPlantFakeRepository) Find(id string) (*domain.Plant, error) {
	return fake.MockAddFn(id)
}

func (fake *storePlantFakeRepository) Store(e *domain.Plant) error {
	return fake.MockAddFn(e)
}

func newLoadPlantPortFakeRepository() *findPlantFakeRepository {
	a := PLantFakeData{}
	err := faker.FakeData(&a)

	if err != nil {
		log.Fatalf(err.Error())
	}

	return &findPlantFakeRepository{
		MockAddFn: func(id string) (*domain.Plant, error) {
			return &domain.Plant{
				Name:      a.Name,
				CreatedAt: a.CreatedAt,
			}, nil
		},
	}
}

func newStorePlantPortFakeRepository() *storePlantFakeRepository {
	return &storePlantFakeRepository{
		MockAddFn: func(e *domain.Plant) error { return nil },
	}
}

func TestSavePlantInMemorySucceed(t *testing.T) {
	lpp := newLoadPlantPortFakeRepository()
	spp := newStorePlantPortFakeRepository()

	service := NewPlantService(lpp, spp)

	plant, err := service.Find("id")

	if err != nil {
		t.Error("Expect error to be nil but got:", err)
	}

	err = service.Store(plant)

	if err != nil {
		t.Error("Expect error to be nil but got:", err)
	}
}

func TestNewPlantService(t *testing.T) {
	type args struct {
		lpp gateway.LoadPlantPort
		spp gateway.StorePlantPort
	}
	tests := []struct {
		name string
		args args
		want PlantService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPlantService(tt.args.lpp, tt.args.spp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPlantService() = %v, want %v", got, tt.want)
			}
		})
	}
}
