package controllers

import (
	"github.com/JPSBarbosa/GO-API-REST/Models"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func GetSession(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	return c.Status(http.StatusOK).JSON(id)
}

func Homepage(c *fiber.Ctx) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(`
	<html>
		<head>
			<title>Página Inicial</title>
		</head>
		<body>
			<button onclick="window.location.href='/Teste'">Ir para página teste</button>
		</body>
	</html>`)
}

func Teste(c *fiber.Ctx) error {
	return c.SendString("Teste")
}

func RegisterUser(c *fiber.Ctx) error {
	var newUser Models.User
	if err := c.ParamsParser(&newUser); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Message": "Invalid Request"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"Message": "User Created"})
}

func AuthUser(c *fiber.Ctx) error {
	var User Models.User
	if err := c.ParamsParser(&User); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Message": "Invalid User"})
	}

	var StoredUser Models.User
	err := db.QueryRow("SELECT id, name, password FROM Users WHERE id=$1", User.Id)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"Message": "Invalid User"})
	}

	if User.Password != StoredUser.Password {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"Message": "Invalid Password"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"Message": "User Authenticated"})
}
