package main

import (
	"fmt"
	"sort"
)

type Student struct {
	name  string
	age   int
	marks int
}

/* ASSIGNMENT 5: SLICES */

func sliceExample() {
	nums := []int{5, 2, 9, 1, 7}

	fmt.Println("Original:", nums)

	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	avg := float64(sum) / float64(len(nums))
	fmt.Println("Sum:", sum)
	fmt.Println("Average:", avg)

	sort.Ints(nums)
	fmt.Println("Sorted:", nums)

	nums = append(nums[:2], nums[3:]...)
	fmt.Println("After removal:", nums)
}

/* ASSIGNMENT 6 & 7 STRUCTS & METHODS */

func (s Student) Display() {
	fmt.Println("Name:", s.name, "Age:", s.age, "Marks:", s.marks)
}

func (s *Student) UpdateMarks(newMarks int) {
	s.marks = newMarks
}

func (s Student) IsPassed() bool {
	return s.marks >= 40
}

func (s Student) Grade() string {
	if s.marks >= 90 {
		return "A"
	} else if s.marks >= 75 {
		return "B"
	} else if s.marks >= 60 {
		return "C"
	} else if s.marks >= 40 {
		return "D"
	}
	return "F"
}

func studentExample() {
	students := []Student{
		{"Sai", 24, 70},
		{"Swaroop", 25, 92},
		{"Manoj", 26, 55},
	}

	for i := 0; i < len(students); i++ {
		students[i].Display()
	}

	top := students[0]
	for i := 1; i < len(students); i++ {
		if students[i].marks > top.marks {
			top = students[i]
		}
	}
	fmt.Println("Top student:")
	top.Display()

	top.UpdateMarks(98)
	fmt.Println("After update:")
	top.Display()
	fmt.Println("Passed?", top.IsPassed())
	fmt.Println("Grade:", top.Grade())
}

/* ASSIGNMENT 8: POINTERS */

func pointerExample() {
	x := 10
	p := &x

	fmt.Println("x:", x)
	fmt.Println("Address:", p)
	fmt.Println("Value via pointer:", *p)

	*p = 50
	fmt.Println("After change:", x)
}

func swap(a *int, b *int) {
	*a, *b = *b, *a
}

func swapExample() {
	x := 10
	y := 20

	fmt.Println("Before swap:", x, y)
	swap(&x, &y)
	fmt.Println("After swap:", x, y)
}

/* MAIN FUNCTION */

func main() {
	fmt.Println("Slice Example")
	sliceExample()

	fmt.Println("\n Student Example")
	studentExample()

	fmt.Println("\n Pointer Example")
	pointerExample()

	fmt.Println("\n Swap Example")
	swapExample()
}
