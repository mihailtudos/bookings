package forms

import (
	"net/http"
	"net/url"
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
		f.Errors.Add(field, "This field is required.")
		return false
	}

	return true
}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
