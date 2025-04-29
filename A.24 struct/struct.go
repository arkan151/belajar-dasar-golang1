package main

import "fmt"

type student struct {
	name  string
	grade int
}

func main() {
	var s1 student
	s1.name = "john wick"
	s1.grade = 2

	fmt.Println("name  :", s1.name)
	fmt.Println("grade :", s1.grade)


	type person struct {
		name string
		age  int
	}
	
	func main() {
		var student struct {
			person
			grade int
		}
	
		student.person = person{"wick", 21}
		student.grade = 2
	
		fmt.Println(student.name)
		fmt.Println(student.age)
		fmt.Println(student.grade)
	}

	type student struct {
		person struct {
			name string
			age  int
		}
		grade   int
		hobbies []string
	}

	var s1 = student{}
	fmt.Println(s1)


	type person struct { name string; age int; hobbies []string }

	var p1 = struct { name string; age int } { age: 22, name: "wick" }
	var p2 = struct { name string; age int } { "ethan", 23 }
	
	fmt.Println(p1)
	fmt.Println(p2)


	type person struct {
		name string `tag1`
		age  int    `tag2`
	}
	
	func main() {
		var p1 = person{"wick", 21}
		fmt.Println(p1)
	}


	type Person struct {
		name string `tag1`
		age  int    `tag2`
	}
	
	type People = Person
	
	func main() {
		var p1 = Person{"wick", 21}
		fmt.Println(p1)
	
		var p2 = People{"wick", 21}
		fmt.Println(Person(p2))
		fmt.Println(People(p1))
	}


	type People1 struct {
		name string
		age int
	}
	
	type People2 = struct {
		name string
		age int
	}
	
	func main() {
		people := People1{"wick", 21}
		fmt.Println(People1(people))
	
		person := People2{"wick", 21}
		fmt.Println(People2(person))
	}


	type student struct {
		name  string
		grade int
	}
	
	func main() {
		var s1 = student{}
		s1.name = "wick"
		s1.grade = 2
	
		var s2 = student{"ethan", 2}
	
		var s3 = student{name: "jason"}
	
		fmt.Println("student 1 :", s1.name)
		fmt.Println("student 2 :", s2.name)
		fmt.Println("student 3 :", s3.name)
	}


	type student struct {
		name  string
		grade int
	}
	
	func main() {
		var s1 = student{name: "wick", grade: 2}
	
		var s2 *student = &s1
		fmt.Println("student 1, name :", s1.name)
		fmt.Println("student 4, name :", s2.name)
	
		s2.name = "ethan"
		fmt.Println("student 1, name :", s1.name)
		fmt.Println("student 4, name :", s2.name)
	}


	type person struct {
		name string
		age  int
	}
	
	type student struct {
		grade int
		person
	}
	
	func main() {
		var s1 = student{}
		s1.name = "wick"
		s1.age = 21
		s1.grade = 2
	
		fmt.Println("name  :", s1.name)
		fmt.Println("age   :", s1.age)
		fmt.Println("age   :", s1.person.age)
		fmt.Println("grade :", s1.grade)
	}


	type person struct {
		name string
		age  int
	}
	
	type student struct {
		person
		age   int
		grade int
	}
	
	func main() {
		var s1 = student{}
		s1.name = "wick"
		s1.age = 21        // age of student
		s1.person.age = 22 // age of person
	
		fmt.Println(s1.name)
		fmt.Println(s1.age)
		fmt.Println(s1.person.age)
	}


	type person struct {
		name string
		age  int
	}
	
	type student struct {
		person
		age   int
		grade int
	}
	
	func main() {
		var p1 = person{name: "wick", age: 21}
		var s1 = student{person: p1, grade: 2}
	
		fmt.Println("name  :", s1.name)
		fmt.Println("age   :", s1.age)
		fmt.Println("grade :", s1.grade)
	}


	type person struct {
		name string
		age  int
	}
	
	type student struct {
		person
		age   int
		grade int
	}
	
	func main() {
		var p1 = person{name: "wick", age: 21}
		var s1 = student{person: p1, grade: 2}
	
		fmt.Println("name  :", s1.name)
		fmt.Println("age   :", s1.age)
		fmt.Println("grade :", s1.grade)
	}


	type person struct {
		name string
		age  int
	}
	
	type student struct {
		person
		age   int
		grade int
	}
	
	func main() {
		var p1 = person{name: "wick", age: 21}
		var s1 = student{person: p1, grade: 2}
	
		fmt.Println("name  :", s1.name)
		fmt.Println("age   :", s1.age)
		fmt.Println("grade :", s1.grade)
	}


	type person struct {
	name string
	age  int
}

func main() {
	var allStudents = []struct {
		person
		grade int
	}{
		{person: person{"wick", 21}, grade: 2},
		{person: person{"ethan", 22}, grade: 3},
		{person: person{"bond", 21}, grade: 3},
	}

	for _, student := range allStudents {
		fmt.Println(student)
	}
}
}