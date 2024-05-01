package packOne

type Person struct {
	name   string
	rollno int
}

func (p *Person) Name() string {
	return p.name
}

func (p *Person) RollNo() int {
	return p.rollno
}
