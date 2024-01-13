package product

import (
	"sesi_12/helpers"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type (
	ProductHandler interface {
		CreateProduct(c *fiber.Ctx) error
		GetAllProduct(c *fiber.Ctx) error
		GetProductByID(c *fiber.Ctx) error
		UpdateProduct(c *fiber.Ctx) error
		DeleteProduct(c *fiber.Ctx) error
	}

	productHandler struct {
		ProductService ProductService
	}
)

func NewProductHandler(ps ProductService) ProductHandler {
	return &productHandler{
		ProductService: ps,
	}
}

func (ph productHandler) CreateProduct(c *fiber.Ctx) error {
	var product Product
	if err := c.BodyParser(&product); err != nil {
		response := helpers.APIResponse("failed to parse request body", fiber.StatusBadRequest, "error", nil)
		return c.Status(response.Meta.Code).JSON(response)
	}

	created, err := ph.ProductService.CreateProduct(&product)
	if err != nil {
		response := helpers.APIResponse("failed to create product", fiber.StatusBadRequest, "error", nil)
		return c.Status(response.Meta.Code).JSON(response)
	}

	response := helpers.APIResponse("success created", fiber.StatusCreated, "success", created)
	return c.Status(response.Meta.Code).JSON(response)
}

func (ph productHandler) GetAllProduct(c *fiber.Ctx) error {
	products, err := ph.ProductService.GetAllProduct()
	if err != nil {
		response := helpers.APIResponse("failed to get all product", fiber.StatusBadRequest, "error", nil)
		return c.Status(response.Meta.Code).JSON(response)
	}

	response := helpers.APIResponse("success get all", fiber.StatusOK, "success", products)
	return c.Status(response.Meta.Code).JSON(response)
}

func (ph productHandler) GetProductByID(c *fiber.Ctx) error {
	id := c.Params("id")
	value, err := strconv.Atoi(id)
	if err != nil {
		response := helpers.APIResponse("invalid product id", fiber.StatusBadRequest, "error", nil)
		return c.Status(response.Meta.Code).JSON(response)
	}

	product, err := ph.ProductService.GetProductByID(value)
	if err != nil {
		response := helpers.APIResponse("failed to get product by id", fiber.StatusBadRequest, "error", nil)
		return c.Status(response.Meta.Code).JSON(response)
	}

	response := helpers.APIResponse("success get by id", fiber.StatusOK, "success", product)
	return c.Status(response.Meta.Code).JSON(response)
}

func (ph productHandler) UpdateProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	value, err := strconv.Atoi(id)
	if err != nil {
		response := helpers.APIResponse("invalid product id", fiber.StatusBadRequest, "error", nil)
		return c.Status(response.Meta.Code).JSON(response)
	}

	product, err := ph.ProductService.GetProductByID(value)
	if err != nil {
		response := helpers.APIResponse("failed to get product by id", fiber.StatusBadRequest, "error", nil)
		return c.Status(response.Meta.Code).JSON(response)
	}

	if err := c.BodyParser(&product); err != nil {
		response := helpers.APIResponse("failed to parse request body", fiber.StatusBadRequest, "error", nil)
		return c.Status(response.Meta.Code).JSON(response)
	}

	updated, err := ph.ProductService.UpdateProduct(product)
	if err != nil {
		response := helpers.APIResponse("failed to update product", fiber.StatusBadRequest, "error", nil)
		return c.Status(response.Meta.Code).JSON(response)
	}

	response := helpers.APIResponse("success updated", fiber.StatusOK, "success", updated)
	return c.Status(response.Meta.Code).JSON(response)
}

func (ph productHandler) DeleteProduct(c *fiber.Ctx) error {
	id := c.Params("id")
	value, err := strconv.Atoi(id)
	if err != nil {
		response := helpers.APIResponse("invalid product id", fiber.StatusBadRequest, "error", nil)
		return c.Status(response.Meta.Code).JSON(response)
	}

	product, err := ph.ProductService.GetProductByID(value)
	if err != nil {
		response := helpers.APIResponse("failed to get product by id", fiber.StatusBadRequest, "error", nil)
		return c.Status(response.Meta.Code).JSON(response)
	}

	deleted, err := ph.ProductService.DeleteProduct(product, value)
	if err != nil {
		response := helpers.APIResponse("failed to update product", fiber.StatusBadRequest, "error", nil)
		return c.Status(response.Meta.Code).JSON(response)
	}

	response := helpers.APIResponse("success deleted", fiber.StatusOK, "success", deleted)
	return c.Status(response.Meta.Code).JSON(response)
}
