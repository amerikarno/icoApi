package adminRepository

import "gorm.io/gorm"

func CustomerInfoJoinCustomerAddresses(db *gorm.DB) *gorm.DB {
	return db.Joins("customer")
}
