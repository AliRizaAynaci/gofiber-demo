package routes

import (
	"demo/database"
	"demo/models"
	"errors"
	"github.com/gofiber/fiber/v2"
	"time"
)

type Order struct {
	ID        uint      `json:"id"`
	User      User      `json:"user"`
	Product   Product   `json:"product"`
	CreatedAt time.Time `json:"order_date"`
}

func CreateResponseOrder(order models.Order, user User, product Product) Order {
	return Order{
		ID:        order.ID,
		User:      user,
		Product:   product,
		CreatedAt: order.CreatedAt,
	}
}

func CreateOrder(ctx *fiber.Ctx) error {
	var order models.Order

	if err := ctx.BodyParser(&order); err != nil {
		return ctx.Status(400).JSON(err.Error())
	}

	var user models.User
	if err := findUser(order.UserRefer, &user); err != nil {
		return ctx.Status(400).JSON(err.Error())
	}

	var product models.Product
	if err := findProduct(order.ProductRefer, &product); err != nil {
		return ctx.Status(400).JSON(err.Error())
	}
	database.Database.Db.Create(&order)

	responseUser := CreateResponseUser(user)
	responseProduct := CreateResponseProduct(product)
	responseOrder := CreateResponseOrder(order, responseUser, responseProduct)

	return ctx.Status(200).JSON(responseOrder)
}

func GetOrders(ctx *fiber.Ctx) error {
	orders := []models.Order{}
	database.Database.Db.Find(&orders)
	responseOrders := []Order{}

	for _, order := range orders {
		var user models.User
		var product models.Product
		database.Database.Db.Find(&user, "id = ?", order.UserRefer)
		database.Database.Db.Find(&product, "id = ?", order.ProductRefer)
		responseOrder := CreateResponseOrder(order, CreateResponseUser(user), CreateResponseProduct(product))
		responseOrders = append(responseOrders, responseOrder)
	}

	return ctx.Status(200).JSON(responseOrders)
}

func FindOrder(id int, order *models.Order) error {
	database.Database.Db.Find(&order, "id = ?", id)
	if order.ID == 0 {
		return errors.New("Order does not exist")
	}
	return nil
}

func GetOrder(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	var order models.Order
	if err != nil {
		return ctx.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err = FindOrder(id, &order); err != nil {
		return ctx.Status(400).JSON(err.Error())
	}

	var user models.User
	var product models.Product
	database.Database.Db.First(&user, order.UserRefer)
	database.Database.Db.First(&product, order.ProductRefer)
	responseUser := CreateResponseUser(user)
	responseProduct := CreateResponseProduct(product)
	responseOrder := CreateResponseOrder(order, responseUser, responseProduct)

	return ctx.Status(200).JSON(responseOrder)
}
