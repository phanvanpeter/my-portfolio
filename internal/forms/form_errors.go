package forms

type formErrors map[string][]string

func (f formErrors) Get(field string) string {
	e := f[field]
	if len(e) == 0 {
		return ""
	}
	return e[0]
}

func (f formErrors) Add(field, msg string) {
	f[field] = append(f[field], msg)
}