package todo

type Todo struct {
    Id int
    Name string
}

func (t Todo) Open() string {
    return t.Name
}

func (t *Todo) Edit(newName string) {
    t.Name = newName
}
