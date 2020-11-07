package local

import (
	"strings"
	"testing"
)

func TestCreate(t *testing.T) {
	repo := NewRepository()
	texts := []string{"a", "b", "c"}
	path, err := repo.Create(texts, "../../font.ttf")
	if err != nil {
		t.Fatal("Error occured")
	}
	if !strings.Contains(path, "tmp") {
		t.Fatal("path error")
	}
}
