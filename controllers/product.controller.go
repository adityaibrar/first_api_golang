package controllers

import (
	"first_api_golang/models"
	"first_api_golang/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ProductController struct {
	DB *gorm.DB
}

var validate = validator.New()

func (productController *ProductController) CreateProduct(c *fiber.Ctx) error {
	var request models.ProductRequest

	if err := c.BodyParser(&request); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "data product invalid")
	}

	if err := validate.Struct(request); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "data product required")
	}

	products := models.Product{
		Name:        request.Name,
		Price:       request.Price,
		Description: request.Description,
	}

	if err := productController.DB.Create(&products).Error; err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to create product")
	}
	return utils.SuccesResponse(c, fiber.StatusCreated, "Product successful insert", products)
}

func (productController *ProductController) DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")

	var product models.Product

	if err := productController.DB.First(&product, id).Error; err != nil {
		return utils.ErrorResponse(c, fiber.StatusNotFound, "Product not found")
	}

	if err := productController.DB.Delete(&product).Error; err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, "Failed to delete product")
	}
	return utils.SuccesResponse(c, fiber.StatusOK, "Product successful deleted", nil)
}

func (productController *ProductController) GetListProduct(c *fiber.Ctx) error {
	var products []models.Product

	if err := productController.DB.Find(&products).Error; err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Products is empty")
	}
	return utils.SuccesResponse(c, fiber.StatusOK, "Successfully get product", products)
}

func (productController *ProductController) DetailProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	var products models.Product

	if err := productController.DB.Find(&products, id).Error; err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Products is empty")
	}
	return utils.SuccesResponse(c, fiber.StatusOK, "Successfully get product", products)
}
