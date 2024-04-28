package util

import (
	"fmt"
	"go-api/internal/entities"
	"regexp"

	"github.com/go-playground/validator/v10"
)

var uuidRegex *regexp.Regexp

func init() {
	uuidRegex, _ = regexp.Compile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)
}

func RegisterCustomValidator(validate *validator.Validate) {
	err := validate.RegisterValidation("uuid", Uuid)
	if err != nil {
		fmt.Println("Error registering custom validation :", err.Error())
	}

	err = validate.RegisterValidation("deviceStatus", DeviceStatus)
	if err != nil {
		fmt.Println("Error registering custom validation :", err.Error())
	}

	err = validate.RegisterValidation("sensorType", SensorType)
	if err != nil {
		fmt.Println("Error registering custom validation :", err.Error())
	}
}

func ParseValidatorErr(err error) []string {
	errMessages := []string{}

	errs := err.(validator.ValidationErrors)
	for _, e := range errs {
		errMessages = append(errMessages, fmt.Sprintf("%s failed on %s", e.Field(), e.Tag()))
	}

	return errMessages
}

func Uuid(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	if !uuidRegex.MatchString(value) {
		return false
	}
	return true
}

func DeviceStatus(fl validator.FieldLevel) bool {
	value := fl.Field().String()
	if value != string(entities.DEVICE_STATUS_ACTIVE) && value != string(entities.DEVICE_STATUS_INACTIVE) {
		return false
	}
	return true
}

func SensorType(fl validator.FieldLevel) bool {
	value := fl.Field().String()

	for _, v := range entities.SensorTypesName {
		if v["slug"] == value {
			return true
		}
	}

	return false
}
