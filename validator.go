package celeritas

import (
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
)

type Validation struct {
	Data   url.Values
	Errors map[string]string
}

func (c *Celeritas) Validator(data url.Values) *Validation {
	return &Validation{
		Data:   data,
		Errors: make(map[string]string),
	}
}

func (v *Validation) Valid() bool {
	return len(v.Errors) == 0
}

func (v *Validation) AddError(key, message string) {
	if _, exist := v.Errors[key]; !exist {
		v.Errors[key] = message
	}
}

func (v *Validation) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)

	return x != ""

}

func (v *Validation) Required(r *http.Request, fields ...string) {
	for _, field := range fields {
		value := r.Form.Get(field)
		if strings.TrimSpace(value) == "" {
			v.AddError(field, "This field can not be blank")
		}
	}
}

func (v *Validation) Check(ok bool, key, message string) {
	if !ok {
		v.AddError(key, message)
	}
}

func (v *Validation) IsEMail(field, value string) {
	if !govalidator.IsEmail(value) {
		v.AddError(field, "invalid E-Mail Adress")
	}
}

func (v *Validation) IsInt(field, value string) {
	if !govalidator.IsInt(value) {
		v.AddError(field, "This field must be an integer value")
	}
}

func (v *Validation) IsFloat(field, value string) {
	if !govalidator.IsFloat(value) {
		v.AddError(field, "This field must be an float value")
	}
}

func (v *Validation) IsDateISO(field, value string) {
	_, err := time.Parse("2006-01-02", value)
	if err != nil {
		v.AddError(field, "This field must be a date in ISO Format YYYY-MM-DD")
	}
}

func (v *Validation) NoSpaces(field, value string) {
	if !govalidator.HasWhitespace(value) {
		v.AddError(field, "This field did not allow whitespace")
	}
}
