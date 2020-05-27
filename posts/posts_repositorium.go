package posts

type NoteEntity struct {
	Content string
}

type Repository struct {}

type INotesRepository interface {
	GetAll() []NoteEntity
	GetByID(id string) NoteEntity
}

func (ns *Repository) GetAll() (notes []NoteEntity) {
	return []NoteEntity{{Content: "Hello world" }}
}

func (ns *Repository) GetByID(id string) (note NoteEntity) {
	return NoteEntity{Content: "Hello world" }
}