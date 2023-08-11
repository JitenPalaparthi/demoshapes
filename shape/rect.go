package shape

type Rect struct {
	L, B float32
}

// Method: Method contains a receiver
func (r *Rect) Area() float32 {
	return r.L * r.B
}

// Function : Function does not contain a receiver
func AreaOf(r *Rect) float32 {
	return r.L * r.B
}
