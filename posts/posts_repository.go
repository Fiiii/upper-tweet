package posts

import "upper-tweet/posts/entity"

type RepositoryImpl struct {}

type Repository interface {
	GetAll() (*[]entity.Post, error)
	GetByID(id string) (*entity.Post, error)
	Save(note *entity.Post) (*entity.Post, error)
	Delete(id string) (int64, error)
	Update(id string, note *entity.Post) (*entity.Post, error)
}

func (pr *RepositoryImpl) GetAll() (*[]entity.Post, error) {
	return &[]entity.Post{{Content: "Hello world" }}, nil
}

func (pr *RepositoryImpl) GetByID(id string) (*entity.Post, error) {
	return &entity.Post{Content: "Hello world" }, nil
}

func (pr *RepositoryImpl) Save(noteData *entity.Post) (*entity.Post, error) {
	return &entity.Post{Content: "Hello world" }, nil
}

func (pr *RepositoryImpl) Update(id string, noteData *entity.Post) (*entity.Post, error) {
	return &entity.Post{Content: "Hello world" }, nil
}

func (pr *RepositoryImpl) Delete(id string) (int64, error) {
	return 1, nil
}