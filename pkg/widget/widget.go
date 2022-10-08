package widget

// A Widget is an example struct
type Widget struct {
	Name  string `json:"name"`
	ID    int    `json:"id"`
	Color Color  `json:"color"`
}

// New creates and returns a new Widget.
func New(name string, id int, color Color) Widget {
	return Widget{name, id, color}
}
