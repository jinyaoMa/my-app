package option

import "github.com/gofiber/fiber/v2"

func Apply(r fiber.Router) fiber.Router {
	option := r.Group("/option")
	{

	}
	return option
}
