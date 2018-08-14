package wellgo

import (
	"github.com/go-playground/validator"
)

type Validator struct{
	validator.Validate
}


//TODO 根据入参与controller参数匹配验证
func (v *Validator) Vl(s interface{}) error{
	return v.Struct(s)
}