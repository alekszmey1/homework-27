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
	if len(s) !=3{
		fmt.Println("введены не корректные данные")
		return nil
	}
	p := student{}
	p.name = s[0]
	p.age,_ = strconv.Atoi(s[1])
	p.grade ,_ = strconv.Atoi(s[2])
	return &p
}

type University struct {
	studentByName map[string]*student
}

//нужна для создания экземпляра университете с занесением нового студента в map
func newUniversity () *University {
	return &University{studentByName: make(map[string]*student)}
}
//добавление студента в срез University
func (u *University)put(s *student)(){
	u.studentByName[s.name] = s
}
func (u *University) getAll()[]*student{
	var students []*student
	for _, v := range u.studentByName {
		students = append(students,v)
	}
	return students
}
func main() {
	university := newUniversity()
	scanner:=bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		l := newStudent(line)
		fmt.Println(l)
		university.put(l)
	}

	for _, v := range university.getAll() {
		fmt.Println(v )
	}

}