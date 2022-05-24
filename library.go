package main

import (
	"fmt"
)

type Book struct {
	Author    string
	Name      string
	Price     int
	Publisher string
}

type Library struct {
	Books []Book
}

func (D Library) srednee() {
	sum := 0
	n := len(D.Books)
	for i := 0; i < n; i++ {
		sum += (D.Books[i].Price)
	}
	avg := (float64(sum)) / (float64(n))
	fmt.Println("Общая оценка библеотеки = ", sum, "Средняя цена = ", avg)
}
func (D Library) BubbleSort(key string) {
	for i := 0; i < len(D.Books); i++ {
		for j := len(D.Books) - 1; j > i; j-- {
			switch key {
			case "Name":
				if D.Books[j-1].Name > D.Books[j].Name {
					swap(D, j-1, j)
				}
			case "Author":
				if D.Books[j-1].Author > D.Books[j].Author {
					swap(D, j-1, j)
				}
			case "Publisher":
				if D.Books[j-1].Publisher > D.Books[j].Publisher {
					swap(D, j-1, j)
				}
			case "Price":
				if D.Books[j-1].Price > D.Books[j].Price {
					swap(D, j-1, j)
				}
			default:
				fmt.Println("Сортировка не выполнена")
			}
		}
	}
}

func swap(D Library, i, j int) {
	tmp := D.Books[i]
	D.Books[i] = D.Books[j]
	D.Books[j] = tmp

}
func main() {
	var key string
	bookstemp := []Book{
		{"Ричард-Мэтисон", "Я - Легенда", 212, "Азбука СБП"},
		{"Э. Ганиев", "Путь к успеху", 500, "Азбука СБП"},
		{"Орловская", "Учебник Английского", 250, "Азбука СбП"},
		{"Орловская", "Учебник Английского для технических вузов", 300, "Азбука СбП"},
		{"Рони-Старший", "Борьба за огонь", 100, "Азбука СбП"},
		{"Стивен Кинг", "Билли Саммерс", 900, "АСТ"},
		{"Казаков Д.", "Удравшие из ада", 80, "АСТ"},
		{"Нора Сакавич", "Элизиум", 250, "Popcorn Books"},
		{"Сэнсом К.", "Мертвая земля: роман", 870, "Азбука"},
	}
	library := Library{bookstemp}
	library.srednee()
	fmt.Println("Введите желаемую сортировку: Name/Author/Publisher/Price")
	fmt.Scanln(&key)
	library.BubbleSort(key)
	fmt.Println(library)
}
