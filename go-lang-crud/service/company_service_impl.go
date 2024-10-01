package service

import (
	"golang-crud/models"
	"golang-crud/repository"
)

// CompanyServiceImpl provides the concrete implementation of the CompanyService interface
type CompanyServiceImpl struct {
	repo *repository.CompanyRepository
}

// NewCompanyServiceImpl creates a new instance of CompanyServiceImpl
func NewCompanyServiceImpl(repo *repository.CompanyRepository) *CompanyServiceImpl {
	return &CompanyServiceImpl{repo: repo}
}

// CreateCompany creates a new company
func (s *CompanyServiceImpl) CreateCompany(company *models.Company) error {
	return s.repo.Create(company)
}

// GetAllCompanies returns a list of all companies
func (s *CompanyServiceImpl) GetAllCompanies() ([]models.Company, error) {
	return s.repo.FindAll()
}

// DeleteCompany deletes a company by ID
func (s *CompanyServiceImpl) DeleteCompany(id string) error {
	return s.repo.DeleteById(id)
}
