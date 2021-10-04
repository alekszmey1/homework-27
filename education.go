package education


type  University struct {
	studentsByName map[string]*Student
}

func NewUniversity()*University  {
	return &University{
		studentsByName: make(map[string]*Student),
	}
}
func (u *University) EnrollStudent(student *Student)error  {
	if _, found := u.studentsByName[student.name]; !found {
		u.studentsByName[student.name] = student
		return nil
	}
	return newErrStudentAlreadyEnrolled(student.name)
}
func (u *University) EnrolledStudents()[]*Student{
	var students []*Student
	for_, v := range u.studentsByName {
		students = append(students, v)
	}
	return students
}

