// controllers/company_controller.go
package controllers

import (
	"golang-crud/models"
	"golang-crud/service"

	"github.com/gin-gonic/gin"
)

type CompanyController struct {
	companyService service.CompanyService // Not a pointer
}

func NewCompanyController(companyService service.CompanyService) *CompanyController {
	return &CompanyController{companyService: companyService}
}

func (cc *CompanyController) CreateCompany(c *gin.Context) {
	var company models.Company
	if err := c.ShouldBindJSON(&company); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := cc.companyService.CreateCompany(&company); err != nil {
		c.JSON(500, gin.H{"error": "Failed to create company"})
		return
	}

	c.JSON(201, gin.H{"message": "Company created successfully", "company": company})
}

func (cc *CompanyController) GetAllCompanies(c *gin.Context) {
	companies, err := cc.companyService.GetAllCompanies()
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to retrieve companies"})
		return
	}

	c.JSON(200, gin.H{
		"companies": companies,
	})
}

func (cc *CompanyController) DeleteCompany(c *gin.Context) {
	id := c.Param("id")
	if err := cc.companyService.DeleteCompany(id); err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete company"})
		return
	}

	c.JSON(200, gin.H{"message": "Company deleted successfully"})
}
