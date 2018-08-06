package wellgo

import (
	"github.com/go-playground/validator"
)

type Validator struct{
	validator.Validate
}

func (v *Validator) Vl(s interface{}) error{
	return v.Struct(s)
}