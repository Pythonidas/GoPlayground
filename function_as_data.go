package main

import "fmt"

const (
	STUDENT = iota
	DOCTOR
	MOM
	GEEK
)

type activity byte

type person struct {
	activities []activity
	speak      func()
}

type people map[string]person

type phrases map[activity]string

func compose(p phrases, a []activity) func() {
	return func() {
		for key, value := range a {
			fmt.Print(p[value])
			if key == len(a)-1 {
				fmt.Println(".")
			} else {
				fmt.Print(" and ")
			}
		}
	}
}

func main() {

	ph := phrases{
		STUDENT: "I read books",
		DOCTOR:  "I fix your head",
		MOM:     "Dinner is ready!",
		GEEK:    "Look ma! My compiler works!"}

	folks := people{
		"Sam":  person{activities: []activity{STUDENT}},
		"Jane": person{activities: []activity{MOM, DOCTOR}},
		"Rob":  person{activities: []activity{GEEK, STUDENT}},
		"Tom":  person{activities: []activity{GEEK, STUDENT, DOCTOR}},
	}

	for name, value := range folks {
		f := compose(ph, value.activities)
		k := value.activities

		folks[name] = person{activities: k, speak: f}
	}

	for name, value := range folks {
		fmt.Printf("%s says: ", name)
		value.speak()
	}
}
