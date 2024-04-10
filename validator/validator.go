package validator

import (
	"crud-compartamos-api/types"
	v "github.com/go-ozzo/ozzo-validation/v4"
	"regexp"
	"time"
)

type IValidator interface {
	ValidateRequest(request types.Client) error
}

type Validator struct{}

func NewValidator() IValidator {
	return &Validator{}
}

func (c *Validator) ValidateRequest(request types.Client) error {
	return v.ValidateStruct(&request, c.Rules(&request)...)
}

func (c *Validator) Rules(request *types.Client) []*v.FieldRules {
	var rules []*v.FieldRules
	regexDNI := regexp.MustCompile("^\\d{8}$")
	regexName := regexp.MustCompile("^[a-zA-ZñÑ\\s]+$")

	rules = append(rules, []*v.FieldRules{
		v.Field(&request.DNI, v.Required, v.Match(regexDNI)),
		v.Field(&request.City, v.Required),
		v.Field(&request.Birthdate, v.Required, v.Date(time.DateOnly).Max(time.Now())),
		v.Field(&request.Gender, v.Required, v.In(gender...)),
		v.Field(&request.LastName, v.Required, v.Match(regexName)),
		v.Field(&request.Name, v.Required, v.Match(regexName)),
	}...)
	return rules
}
