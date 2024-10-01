package service

import "golang-crud/models"

// CompanyService defines the behavior expected for the company-related operations
type CompanyService interface {
	CreateCompany(company *models.Company) error
	GetAllCompanies() ([]models.Company, error)
	DeleteCompany(id string) error
}
