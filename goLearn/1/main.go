package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Grade map[string]int

type Student struct {
	deleted bool
	Name    string `json:"name"`
	Age     int8   `json:"age"`
	Grade   Grade  `json:"grade"`
}

type Stus []Student

func (stus *Stus) add(name string, age int8, grade Grade) {
	*stus = append(*stus, Student{false, name, age, grade})

}
func (stus Stus) edit(id int, s Student) error {
	if id < 0 || len(stus) <= id {
		return errors.New("edit: student id out of range")
	}
	stus[id] = s
	return nil
}
func (stus Stus) delete(id int) error {
	if id < 0 || len(stus) <= id {
		return errors.New("delete: student id out of range")
	}
	stus[id].deleted = true
	return nil
}
func (stus Stus) list() error {
	for _, stu := range stus {
		if stu.deleted {
			continue
		}
		str, err := json.Marshal(stu)
		if err != nil {
			return err
		}

		fmt.Printf("%s\n", str)
	}
	return nil
}

func main() {
	var stus Stus
	stus.add(
		"A",
		15,
		map[string]int{"a": 100, "b": 85},
	)
	stus.add(
		"B",
		13,
		map[string]int{"a": 90, "b": 88},
	)
	stus.list()
	fmt.Println("------------------")
	s := Student{
		false,
		"C",
		19,
		map[string]int{"a": 90, "b": 88},
	}
	stus.edit(1, s)
	stus.list()
	fmt.Println("------------------")
	stus.delete(0)
	stus.list()
}
