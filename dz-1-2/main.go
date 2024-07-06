package main

import "fmt"

func CelsiusToFahrenheit(t float64) float64 { // 11. Конвертер температур
	return t*9/5 + 32
}

func strLen(str string) int { // 13. Длина строки
	size := 0
	for range str {
		size++
	}
	return size
}

func containsElement(arr []string, x string) bool { // 14. Содержит ли массив?
	for _, el := range arr {
		if el == x {
			return true
		}
	}
	return false
}

func countAverage(arr []int) float64 { //15. Содержит ли массив?
	average := 0.0
	for _, el := range arr {
		average += float64(el)
	}
	return average / float64(len(arr))
}

func isPalindrome(s string) bool { // 17. Палиндром
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func minAndMax(arr []int) (int, int) { // 18. Найти минимум и максимум
	maximum, minimum := arr[0], arr[0]
	for _, elem := range arr {
		if elem > maximum {
			maximum = elem
		}
		if elem < minimum {
			minimum = elem
		}
	}
	return minimum, maximum
}

func deleteElement(arr *[]int, idx int) { // 19. Удаление элемента из слайса
	*arr = append((*arr)[0:idx], (*arr)[idx+1:]...)
}

func findElement(arr []int, x int) int { // 20. Линейный поиск
	for idx, el := range arr {
		if el == x {
			return idx
		}
	}
	return -1
}

func main() {
	var t float64
	fmt.Scan(&t)
	fmt.Println(CelsiusToFahrenheit(t))

	var n int // 12. Обратный отсчет
	fmt.Scan(&n)
	for n >= 0 {
		fmt.Print(n, " ")
		n--
	}

	var str1 string
	fmt.Scan(&str1)
	fmt.Println(strLen(str1))

	var sz1 int
	fmt.Scan(&sz1)
	var arr1 []string
	for i := 0; i < sz1; i++ {
		var elem string
		fmt.Scan(&elem)
		arr1 = append(arr1, elem)
	}
	var f string
	fmt.Scan(&f)
	fmt.Println(containsElement(arr1, f))

	var sz2 int
	fmt.Scan(&sz2)
	var arr2 []int
	for i := 0; i < sz2; i++ {
		var elem int
		fmt.Scan(&elem)
		arr2 = append(arr2, elem)
	}
	fmt.Println(countAverage(arr2))

	var str2 string
	fmt.Scan(&str2)
	fmt.Println(isPalindrome(str2))

	fmt.Scan(&n) // 16. Таблица умножения
	for i := 1; i <= 10; i++ {
		fmt.Println(n, "*", i, "=", n*i)
	}

	var sz3 int
	fmt.Scan(&sz3)
	var arr3 []int
	for i := 0; i < sz3; i++ {
		var elem int
		fmt.Scan(&elem)
		arr3 = append(arr3, elem)
	}
	fmt.Println(minAndMax(arr3))

	var sz4 int
	fmt.Scan(&sz4)
	var arr4 []int
	for i := 0; i < sz4; i++ {
		var elem int
		fmt.Scan(&elem)
		arr4 = append(arr4, elem)
	}
	var idx int
	fmt.Scan(&idx)
	deleteElement(&arr4, idx)
	fmt.Println(arr4)

	var sz5 int
	fmt.Scan(&sz5)
	var arr5 []int
	for i := 0; i < sz5; i++ {
		var elem int
		fmt.Scan(&elem)
		arr5 = append(arr5, elem)
	}
	var x int
	fmt.Scan(&x)
	fmt.Println(findElement(arr5, x))
}
