package filters

type NameLike struct {
	nl string
}

func NewNameLike(nl string) *NameLike {
	return &NameLike{nl: nl}
}

func (n *NameLike) Where() (string, []any) {
	return "name LIKE ?", []any{n.nl}
}
