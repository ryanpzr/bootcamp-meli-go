package repository

type Repository struct {
	values []string
}

func (r *Repository) SaveValue(s string) {
	r.values = append(r.values, s)
}
