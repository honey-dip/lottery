package service

type ImageRepository interface {
	Create(winners []string) (string, error)
}
