package model

import (
	"gali.com/config"
)

type Customers struct {
	CustomerCode string `json:"customer_code" form:"customer_code" gorm:"primaryKey"`
	CustmerName  string `json:"custmer_name" form:"custmer_name"`
	CustomerType string `json:"customer_type" form:"customer_type"`
}

func (customer *Customers) CreateCustomer() error {
	if err := config.DB.Create(customer).Error; err != nil {
		return err
	}
	return nil
}

func (customer *Customers) UpdateCustomer(customer_code string) error {
	// if err := config.DB.Model(&Customers{}).Where("customer_code = ?", customer_code).Updates(customer).Error; err != nil {
	if err := config.DB.Model(&Customers{}).Where("customer_code = ?", customer_code).Updates(&customer).Error; err != nil {
		return err
	}
	return nil
}

func (customer *Customers) DeleteByCode(customer_code string) error {
	if err := config.DB.Model(&Customers{}).Where("customer_code = ?", customer_code).Delete(&customer).Error; err != nil {
		return err
	}
	return nil
}

// func (customer *Customers) DeleteCustomer() error {
// 	if err := config.DB.Delete(&customer).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

func GetOneByCode(customer_code string) (Customers, error) {
	var customer Customers
	result := config.DB.Where("customer_code = ? ", customer_code).First(&customer)
	return customer, result.Error
}

func GetAll(keywords string) ([]Customers, error) {
	var customers []Customers
	result := config.DB.Where("customer_code LIKE ? OR custmer_name LIKE ?", "%"+keywords+"%", "%"+keywords+"%").Find(&customers)
	return customers, result.Error
}
