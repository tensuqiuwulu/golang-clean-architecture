package repository

import (
	"github.com/tensuqiuwulu/golang-clean-architecture/src/model/entity"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	CreateCustomer(db *gorm.DB, customer *entity.CustomerModel) error
}

type customerRepositoryImpl struct{}

func NewCustomerRepository() CustomerRepository {
	return &customerRepositoryImpl{}
}

func (r *customerRepositoryImpl) CreateCustomer(db *gorm.DB, customer *entity.CustomerModel) error {
	err := db.Table("customers").Create(customer).Error
	if err != nil {
		return err
	}
	return nil
}
