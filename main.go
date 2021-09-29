package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
type Student struct {
	name string
	age int
	grade int
}

func main() {
	scanner:=bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	m := make(map[int]*Student)
	k:=1
	for scanner.Scan() {
		line := scanner.Text()
		s:= strings.Split(line," ")
		p:= Student{}
		for i := 0; i < len(s); i++ {
			if i == 0 {
				p.name = s[i]
			}
			if i == 1{
				p.age,_ = strconv.Atoi(s[i])
			}
			if i == 2{
				fmt.Println(s[i])
				p.grade ,_ = strconv.Atoi(s[i])
			}

		}
		fmt.Println("-------------------")
		fmt.Println(p)
		m = map[int]*Student{
			k: *p,
		}
		k++
	}
	fmt.Println(m)
}
