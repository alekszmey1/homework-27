package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
type student struct {
	name string
	age int
	grade int
}
func newStudent(l string) *student{
	s:= strings.Split(l," ")
	p := student{}
	p.name = s[0]
	p.age,_ = strconv.Atoi(s[1])
	p.grade ,_ = strconv.Atoi(s[2])
	return &p
}

type University struct {
	studentByName map[string]*student
}
func newUniversity (stud *student) *University {//нужна для создания экземпляра университете с занесением нового студента в map
		studentNew := map[string]*student{
		stud.name : stud,
	}
		s := University{studentNew}
	return &s
}
/*func (u *University)put(s *student)(){//добавление студента в срез University
	return
}

func (u *University) get()[]student  {
return nil
}*/
func main() {
	//university := NewUniversity()
	//m := make(map[string]student)
	scanner:=bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		l := newStudent(line)
		fmt.Println(l)
		p:=newUniversity(l)
		fmt.Println(p)
	}
	//university.put(l)
}
