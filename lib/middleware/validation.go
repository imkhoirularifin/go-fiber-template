package middleware

import (
	"go-fiber-template/internal/domain/dto"
	"go-fiber-template/pkg/xvalidator"

	"github.com/gofiber/fiber/v2"
	"github.com/ryanbekhen/di"
)

type Placement string

const (
	PlacementBody   Placement = "body"
	PlacementQuery  Placement = "query"
	PlacementParam  Placement = "param"
	PlacementCookie Placement = "cookie"
)

func Validate[V any](placement Placement) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var v V
		var err error

		switch placement {
		case PlacementBody:
			err = c.BodyParser(&v)
		case PlacementQuery:
			err = c.QueryParser(&v)
		case PlacementParam:
			err = c.ParamsParser(&v)
		case PlacementCookie:
			err = c.CookieParser(&v)
		}
		if err != nil {
			return err
		}

		// get validator client from global dependency, this will not create new dependency each time this function is called
		xvalidatorClient := di.MustResolve[*xvalidator.Client]()

		if validationErrors := xvalidatorClient.ValidateStructWithLang(v, c.Get("Accept-Language", "id")); validationErrors != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(dto.ResponseDto{
				Message: "Invalid Input",
				Errors:  validationErrors,
			})
		}

		c.Locals("parser", &v)
		return c.Next()
	}
}
