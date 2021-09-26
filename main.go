package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*Напишите программу, которая считывает ввод с stdin,
создаёт структуру student и записывает указатель на
структуру в хранилище map[studentName] *Student.
Программа должна получать строки в бесконечном цикле, создать структуру Student через функцию newStudent,
далее сохранить указатель на эту структуру в map, а после получения EOF (ctrl + d)
вывести на экран имена всех студентов из хранилища. Также необходимо реализовать методы put, get.
 */

type student struct {
	name string
	age int
	grade int
}

func newStudent(a string)(p student){
	s := strings.Split(a," ")
	for i := 0; i < len(s); i++ {
		if i == 0 {
			p.name = s[i]
			fmt.Println(p.name)
		}
		if i == 1{
			p.age,_ = strconv.Atoi(s[i])
			fmt.Println(p.age)
		}
		if i == 2{
			fmt.Println(s[i])
			p.grade ,_ = strconv.Atoi(s[i])
			fmt.Println(p.grade)
		}
		fmt.Println(p)
	}
	return p
}




func main()  {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	newStudent(text)
}
