package models

import (
	"strconv"

	"github.com/asaskevich/govalidator"
)

//return true if valid
func InitValidator() {
	govalidator.TagMap["name"] = govalidator.Validator(validateName)
	govalidator.TagMap["id"] = govalidator.Validator(validateId)
}

func validateName(s string) bool {

	if !govalidator.StringLength(s, "2", "10") && !govalidator.IsUTFLetter(s) {
		return false
	}
	return true
}

func validateId(s string) bool {

	if !govalidator.IsInt(s) {
		return false
	}
	int, _ := strconv.Atoi(s)
	if int < 1 {
		return false
	}
	return true
}
