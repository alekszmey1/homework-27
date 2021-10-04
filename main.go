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
var students []string


func (s student) allStudents(stud student)(students[]string){
	var v [3] string
	v[0] = s.name
	v[1] = strconv.Itoa(s.age)
	v[2] = strconv.Itoa(s.grade)
	students = append(students,v[0],v[1],v[2],"\n")
	fmt.Println(students)
	return
}

func main() {
	scanner:=bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	//m := make(map[int]*Student)
	for scanner.Scan() {
		line := scanner.Text()
		s:= strings.Split(line," ")
		p:= student{}
		for i := 0; i < len(s); i++ {
			if i == 0 {
				p.name = s[i]
			}
			if i == 1{
				p.age,_ = strconv.Atoi(s[i])
			}
			if i == 2{
				p.grade ,_ = strconv.Atoi(s[i])
			}
		}
		fmt.Println("-------------------")
		p.allStudents(p)
	}
}