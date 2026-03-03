package main

import "fmt"

/*  ASSIGNMENT 1 */

func evenOdd() {
	var num int
	fmt.Print("Enter number: ")
	fmt.Scan(&num)

	if num%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}

func largestThree() {
	var a, b, c int
	fmt.Print("Enter numbers: ")
	fmt.Scan(&a, &b, &c)

	if a > b && a > c {
		fmt.Println("Largest:", a)
	} else if b > c {
		fmt.Println("Largest:", b)
	} else {
		fmt.Println("Largest:", c)
	}
}

func leapYear() {
	var year int
	fmt.Print("Enter year: ")
	fmt.Scan(&year)

	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		fmt.Println("Leap Year")
	} else {
		fmt.Println("Not Leap Year")
	}
}

/* ASSIGNMENT 2 */

func print1to100() {
	for i := 1; i <= 100; i++ {
		fmt.Println(i)
	}
}

func even1to50() {
	for i := 2; i <= 50; i += 2 {
		fmt.Println(i)
	}
}

func table() {
	var n int
	fmt.Scan(&n)

	for i := 1; i <= 10; i++ {
		fmt.Println(n, "x", i, "=", n*i)
	}
}

func factorial() {
	var n int
	fmt.Scan(&n)

	fact := 1
	for i := 1; i <= n; i++ {
		fact *= i
	}
	fmt.Println("Factorial:", fact)
}

func fibonacci() {
	var n int
	fmt.Print("Enter fibonacci number: ")
	fmt.Scan(&n)

	a, b := 0, 1
	for i := 1; i <= n; i++ {
		fmt.Print(a, " ")
		a, b = b, a+b
	}
	fmt.Println()
}

/* ASSIGNMENT 3 */

func add(a, b int) int {
	return a + b
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func reverse(n int) int {
	rev := 0
	for n > 0 {
		rev = rev*10 + n%10
		n /= 10
	}
	return rev
}

/* ASSIGNMENT 4 */

func arrayOperations() {
	arr := [5]int{10, 20, 30, 40, 50}

	sum := 0
	max := arr[0]
	min := arr[0]

	for i := 0; i < 5; i++ {
		sum += arr[i]

		if arr[i] > max {
			max = arr[i]
		}
		if arr[i] < min {
			min = arr[i]
		}
	}

	fmt.Println("Array:", arr)
	fmt.Println("Sum:", sum)
	fmt.Println("Average:", sum/5)
	fmt.Println("Max:", max)
	fmt.Println("Min:", min)

	for i, j := 0, 4; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	fmt.Println("Reversed:", arr)
}

/* MAIN MENU */

func main() {
	var choice int

	fmt.Println("1 Even or Odd")
	fmt.Println("2 Largest of Three")
	fmt.Println("3 Leap Year")
	fmt.Println("4 Print 1 to 100")
	fmt.Println("5 Even 1 to 50")
	fmt.Println("6 Multiplication Table")
	fmt.Println("7 Factorial")
	fmt.Println("8 Fibonacci")
	fmt.Println("9 Array Operations")
	fmt.Print("Choose option: ")

	fmt.Scan(&choice)

	switch choice {
	case 1:
		evenOdd()
	case 2:
		largestThree()
	case 3:
		leapYear()
	case 4:
		print1to100()
	case 5:
		even1to50()
	case 6:
		table()
	case 7:
		factorial()
	case 8:
		fibonacci()
	case 9:
		arrayOperations()
	default:
		fmt.Println("Invalid choice")
	}
}
