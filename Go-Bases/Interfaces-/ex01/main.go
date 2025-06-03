package main

import "fmt"

type RegisterStudents interface {
	showInfos() string
}

type Students struct {
	name    string
	surname string
	id      int
	date    string
}

func (s *Students) showInfos() string {
	return fmt.Sprint(s.name, s.surname, s.id, s.date)
}

func details(r RegisterStudents) {
	fmt.Println(r.showInfos())
}

func main() {
	student := Students{name: "Ryan", surname: "Lima", id: 1, date: "07/10/2004"}
	student02 := Students{name: "Marcos", surname: "Oliveira", id: 2, date: "09/12/2000"}
	details(&student)
	details(&student02)
}
