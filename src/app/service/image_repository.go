package service

type ImageRepository interface {
	Create(texts []string, fontpath string) (string, error)
}
