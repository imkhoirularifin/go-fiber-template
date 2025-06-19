package xvalidator

var XValidator *Validator

func Setup() {
	var err error
	XValidator, err = NewValidator(
		WithCustomValidator(&DateValidator{}),
	)
	if err != nil {
		panic(err)
	}
}
