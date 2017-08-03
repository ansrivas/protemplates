package internal

type Project interface {
	Create(appname string) error
}
