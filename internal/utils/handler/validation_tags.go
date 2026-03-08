package handlerutils

type validationTag string

const (
	tagRequired validationTag = "required"
	tagEmail    validationTag = "email"
	tagMin      validationTag = "min"
	tagMax      validationTag = "max"
)
