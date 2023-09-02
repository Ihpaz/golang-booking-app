package service

import (
	"errors"
	"math/rand"

	"github.com/Ihpaz/golang-restapi-userfamily/entity"
	"github.com/Ihpaz/golang-restapi-userfamily/repository"
)

type CustomerService interface {
	Validate(customer *entity.Customer) error
	Create(customer *entity.Customer) (*entity.Customer, error)
	FindAll() (*[]entity.Customer, error)
	FindCustomerByCstId(customer *entity.Customer, uid uint64) (*entity.Customer, error)
	UpdateACustomer(customer *entity.Customer, uid uint64) (*entity.Customer, error)
	DeleteACustomer(customer *entity.Customer, uid uint64) (int64, error)
}

type servicecustomer struct{}

var (
	repo repository.CustomerRepository
)

func NewCustomerService(repository repository.CustomerRepository) CustomerService {
	repo = repository
	return &servicecustomer{}
}

func (*servicecustomer) Validate(customer *entity.Customer) error {
	if customer == nil {
		err := errors.New("The customer is empty")
		return err
	}
	if customer.Cst_name == "" {
		err := errors.New("The customer name is empty")
		return err
	}
	if customer.Cst_dob_date.IsZero() {
		err := errors.New("The customer date of birth is empty")
		return err
	}
	return nil
}

func (*servicecustomer) Create(customer *entity.Customer) (*entity.Customer, error) {
	customer.Cst_id = rand.Int63()
	return repo.Save(customer)
}

func (*servicecustomer) FindAll() (*[]entity.Customer, error) {
	return repo.FindAll()
}

func (*servicecustomer) FindCustomerByCstId(customer *entity.Customer, uid uint64) (*entity.Customer, error) {
	return repo.FindCustomerByCstId(customer, uid)
}

func (*servicecustomer) UpdateACustomer(customer *entity.Customer, uid uint64) (*entity.Customer, error) {
	customer.Cst_id = rand.Int63()
	return repo.UpdateACustomer(customer, uid)
}

func (*servicecustomer) DeleteACustomer(customer *entity.Customer, uid uint64) (int64, error) {
	return repo.DeleteACustomer(customer, uid)
}