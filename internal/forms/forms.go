package forms

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"net/http"
	"net/url"
	"strings"
)

// Form creates a custom form struct, embeds an url.Values object
type Form struct {
	url.Values
	Errors errors
}

// New initializes a form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		map[string][]string{},
	}
}

// Has checks if the form field is in post and not empty
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if len(x) == 0 {
		return false
	}

	return true
}

func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if len(strings.TrimSpace(value)) == 0 {
			f.Errors.Add(field, "This field is required")
		}
	}
}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

func (f *Form) MinLength(field string, length int, r *http.Request) bool {
	x := r.Form.Get(field)

	if len(x) < length {
		f.Errors.Add(field, fmt.Sprintf("Must be at least %d characters long.", length))
		return false
	}

	return true
}

func (f *Form) MaxLength(field string, length int, r *http.Request) bool {
	x := r.Form.Get(field)

	if len(x) > length {
		f.Errors.Add(field, fmt.Sprintf("Must be max %d characters long.", length))
		return false
	}

	return true
}

func (f *Form) IsEmail(field string) bool {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Invalid email address.")
		return false
	}

	return true
}
