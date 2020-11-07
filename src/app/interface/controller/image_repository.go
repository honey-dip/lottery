package controller

type ImageRepository interface {
	Create(texts []string, fontpath string) (string, error)
}
