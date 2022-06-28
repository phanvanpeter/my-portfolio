package forms

import (
	"fmt"
	"net/url"
	"strings"
)

type Form struct {
	url.Values
	Errors FormErrors
}

func New(data url.Values) *Form{
	return &Form{
		data,
		FormErrors{},
	}
}

func (f *Form) IsValid() bool {
	return len(f.Errors) == 0
}

func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		s := f.Get(field)
		if strings.TrimSpace(s) == "" {
			f.Errors.Add(field, "This field cannot be blank.")
		}
	}
}

func (f *Form) MinLength(field string, length int) bool {
	s := f.Get(field)
	if len(s) < length {
		f.Errors.Add(field, fmt.Sprintf("Minimum length must be at least %d characters.", length))
		return false
	}
	return true
}
