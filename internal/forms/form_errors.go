package forms

type FormErrors map[string]string

func (f FormErrors) GetAll() FormErrors {
	return f
}

func (f FormErrors) Get(field string) string {
	return f[field]
}

func (f FormErrors) Add(field, msg string) {
	f[field] = msg
}