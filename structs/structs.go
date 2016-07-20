package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	fmt.Println(Person{"Bob", 20})
	fmt.Println(Person{name: "Alice", age: 30})
	fmt.Println(Person{name: "Fred"})
	fmt.Println(&Person{name: "Ann", age: 30})

	s := Person{name: "Sean", age: 50}

	fmt.Println(s.name)

	sp := &s

	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(s.age)

	fmt.Println(sp)

	fmt.Println("--------------------")

	persons := []Person{
		Person{"Panqd", 20},
		Person{"Panqd", 21},
	}
	for _, person := range persons {
		(&person).age = 22 // 无法修改值为22
	}
	fmt.Println(persons)

	for i := 0; i < len(persons); i++ {
		persons[i].age = 22
	}
	fmt.Println(persons)
}
