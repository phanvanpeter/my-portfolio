package forms

import "net/url"

type Form struct {
	url.Values
	Errors map[string][]string
}

func (f *Form) IsValid() bool {
	return len(f.Errors) == 0
}

func (f *Form) Required(s ...string) bool {
	return false
}

func (f *Form) MinLength(s string, length int) bool {
	return len(s) >= length
}
