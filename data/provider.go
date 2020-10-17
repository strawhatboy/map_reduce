package data

//Provider ...
// the interface for providers
type Provider interface {
	LoadData() error
	ReadData() interface{}
	SetPath(p string)
}
