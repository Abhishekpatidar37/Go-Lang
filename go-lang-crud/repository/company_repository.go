package repository

import (
	"golang-crud/models"

	"gorm.io/gorm"
)

type CompanyRepository struct {
	DB *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) *CompanyRepository {
	return &CompanyRepository{DB: db}
}

func (r *CompanyRepository) Create(company *models.Company) error {
	return r.DB.Create(company).Error
}

func (r *CompanyRepository) FindAll() ([]models.Company, error) {
	var companies []models.Company
	err := r.DB.Find(&companies).Error
	return companies, err
}

func (r *CompanyRepository) DeleteById(id string) error {
	return r.DB.Delete(&models.Company{}, id).Error
}
