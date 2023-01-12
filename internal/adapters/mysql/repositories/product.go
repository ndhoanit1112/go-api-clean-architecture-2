package repositories

import (
	"encoding/json"

	"github.com/ndhoanit1112/go-api-clean-architecture-2/internal/adapters/mysql/dbmodels"
	"github.com/ndhoanit1112/go-api-clean-architecture-2/internal/entities"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func InitProductRepo(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (repo *ProductRepository) CreateProduct(data *entities.Product) error {
	config, err := json.Marshal(data.Config)
	if err != nil {
		return err
	}

	product := dbmodels.Product{
		Name: data.Name,
		BaseEntity: dbmodels.BaseEntity{
			ID:        data.ID,
			CreatedAt: data.Created,
			UpdatedAt: data.Updated,
		},
		MaxFailureCountUser: data.MaxFailureCountUser,
		MaxFailureTimeUser:  data.MaxFailureCountUser,
		Config:              datatypes.JSON(config),
	}
	result := repo.db.Create(&product)
	return result.Error
}

func (repo *ProductRepository) GetProducts() ([]entities.Product, error) {
	var dbProducts []dbmodels.Product
	result := repo.db.Find(&dbProducts)
	if result.Error != nil {
		return nil, result.Error
	}

	var products []entities.Product
	for _, prod := range dbProducts {
		item := entities.Product{
			Name: prod.Name,
			BaseEntity: entities.BaseEntity{
				ID:      prod.ID,
				Created: prod.CreatedAt,
				Updated: prod.UpdatedAt,
			},
			MaxFailureCountUser: prod.MaxFailureCountUser,
			MaxFailureTimeUser:  prod.MaxFailureCountUser,
			//Config:              datatypes.JSON(config),
		}
		products = append(products, item)
	}

	return products, nil
}
