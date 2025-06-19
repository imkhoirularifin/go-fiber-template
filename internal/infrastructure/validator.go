package infrastructure

import (
	"go-fiber-template/lib/xvalidator"
)

func setupValidator() {
	var err error
	val, err = xvalidator.NewValidator(
		xvalidator.WithCustomValidator(&xvalidator.DateValidator{}),
	)
	if err != nil {
		panic(err)
	}
}
