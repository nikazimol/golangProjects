package main

import (
	"fmt"
	"unicode"
)

func plus(a, b int) int { // 2. Сложение чисел
	return a + b
}

func oddOrEven(a int) { // 3. Четное или нечетное
	if a%2 == 0 {
		fmt.Println("Четное")
	} else {
		fmt.Println("Нечетное")
	}
}

func maximum(a, b, c int) int { // 4. Максимум из трех чисел
	if a < b {
		a, b = b, a
	}
	if a < c {
		a, c = c, a
	}
	return a
}

func factorial(n uint64) uint64 { // 5. Факториал числа
	var f, i uint64 = 1, 2
	for i <= n {
		f *= i
		i++
	}
	return f
}

func isVowel(s string) bool { // 6. Проверка символа (на русском и английском)
	rs := []rune(s)
	if unicode.ToLower(rs[0]) == 97 || unicode.ToLower(rs[0]) == 101 || unicode.ToLower(rs[0]) == 105 ||
		unicode.ToLower(rs[0]) == 111 || unicode.ToLower(rs[0]) == 117 || unicode.ToLower(rs[0]) == 121 ||
		unicode.ToLower(rs[0]) == 'а' || unicode.ToLower(rs[0]) == 'е' || unicode.ToLower(rs[0]) == 'ё' ||
		unicode.ToLower(rs[0]) == 'о' || unicode.ToLower(rs[0]) == 'у' || unicode.ToLower(rs[0]) == 'и' ||
		unicode.ToLower(rs[0]) == 'э' || unicode.ToLower(rs[0]) == 'ы' || unicode.ToLower(rs[0]) == 'я' ||
		unicode.ToLower(rs[0]) == 'ю' {
		return true
	}
	return false
}

func primeNumbers(n int) { // 7. Простые числа
	var sieve []int
	p := 2
	for i := 0; i <= n; i++ {
		sieve = append(sieve, i)
	}
	sieve[1] = 0
	for p*p < n {
		for i := p * p; i <= n; i += p {
			sieve[i] = 0
		}
		for j := p; j <= n; j++ {
			if sieve[j] > p {
				p = sieve[j]
				break
			}
		}
	}
	for _, elem := range sieve {
		if elem != 0 {
			fmt.Println(elem)
		}
	}
}

func reversedString(s *string) string { // 8. Строка и ее перевертыш
	rs := []rune(*s)
	var newString []rune
	for i := len(rs) - 1; i >= 0; i-- {
		newString = append(newString, rs[i])
	}
	*s = string(newString)
	return *s
}

func sumElements(array []int) int { // 9. Массив и его сумма
	sum := 0
	for i := 0; i < len(array); i++ {
		sum += array[i]
	}
	return sum
}

type Rectangle struct { // 10. Структуры и методы
	wight  float64
	height float64
}

func (r *Rectangle) area() float64 {
	return r.wight * r.height
}

func main() {
	fmt.Println("Привет, мир!") // 1. Привет, мир!

	var pa, pb int
	fmt.Scan(&pa, &pb)
	fmt.Println(plus(pa, pb))

	var a int
	fmt.Scan(&a)
	oddOrEven(a)

	var ma, mb, mc int
	fmt.Scan(&ma, &mb, &mc)
	fmt.Println(maximum(ma, mb, mc))

	var fn uint64
	fmt.Println(factorial(fn))

	var ch string
	fmt.Scan(&ch)
	fmt.Println(isVowel(ch))

	var pn int
	fmt.Scan(&pn)
	primeNumbers(pn)

	var s string
	fmt.Scan(&s)
	fmt.Println(s)
	reversedString(&s)
	fmt.Println(s)

	arr := []int{1, 2, 5, 10}
	fmt.Println(sumElements(arr))

	r := Rectangle{1.2, 6.8}
	fmt.Println(r.area())
}
