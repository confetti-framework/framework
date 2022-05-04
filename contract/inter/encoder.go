package inter

type Encoder interface {
	IsAble(object interface{}) bool
	EncodeThrough(app App, object interface{}, encoders []Encoder) (string, error)
}
