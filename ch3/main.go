package main

// Типы данных:
// 1) Фундоментальные - числа, строки и булевы значения
// 2) Составные - массивы и структуры
// 3) Ссылочные - указатели, срезы, отображения, функции и каналы
// 4) интерфейсы
func main() {
	// Фундоментальные
	// 3.1 Целые числа
	// int, uint
	// 8, 16, 32, 64
	// rune = int32 (Unicode)
	// byte = uint8
	// int8 -- диапозон значений от -128 до 127
	// uint8 -- диапозон значений от 0 до 255

	// Бинарные операторы:
	// * / + - >> << & &^ ^ | == != < <= > >= && ||

	// %d - 10
	// %o - 8
	// %x - 16
	/*
		o := 0666
		fmt.Printf("%T %[1]d %[1]o %#[1]o\n", o)
		x := int64(0xdeadbeef)
		fmt.Printf("%T %[1]d %[1]x %#[1]X\n", x)
	*/

	// руны выводятся с помощью %c или %q - с кавычками

	// 3.2 числа с плавающей точкой
	// float32 float64
	// сравнение с NaN всегда false
	/*
		var z float64
		fmt.Println(z, -z, 1/z, -1/z, z/z)
	*/

	// 3.5 Строки
	// Строка - неизменяемая последовательность байтов
	// len - кол-во байтов
	/*
		s := "hello, world"
		fmt.Println(len(s))
		fmt.Println(s[1])
		fmt.Println(string(65))
		fmt.Println(string(0x4eaccadadc))
	*/
	// unicode/utf8 -- пакет с функциями для
	// кодирования и декодирования рун в виде байтов (UTF-8)

	// fmt.Println(basename("a/b/c.go"))
	// fmt.Println(basename("c.d.go"))
	// fmt.Println(basename("abc"))
	// fmt.Println(basename2("a/b/c.go"))
	// fmt.Println(basename2("c.d.go"))
	// fmt.Println(basename2("abc"))

	// fmt.Println(comma("1022"))
	// fmt.Println(comma("102205468498479498498490"))
	// fmt.Println(comma("100000"))

	// fmt.Println(comma("1022"))
	// fmt.Println(comma("102.2"))
	// fmt.Println(comma("1.022"))
	// fmt.Println(comma("10.22"))
	// fmt.Println(comma("100000"))

	// 3.5.5 преобразование между строками и числами
	// x := 123
	// y := fmt.Sprintf("%d", x)
	// fmt.Println(y, strconv.Itoa(x))

	// fmt.Println(y, strconv.FormatInt(int64(x), 2))
	// s := fmt.Sprintf("x = %b", x)
	// fmt.Println(s)

	// a, _ := strconv.Atoi("123")
	// b, _ := strconv.ParseInt("123", 10, 64)
	// fmt.Println(a, b)

	// Константы

	// const (
	// 	a = 1
	// 	b
	// 	c = 2
	// 	d
	// )

	// fmt.Println(a, b, c, d)

	// Генератор констант
	// type Weekday int
	// const (
	// 	Sunday Weekday = iota + 1
	// 	Monday
	// 	Tuesday
	// 	Wednesday
	// 	Friday
	// 	Saturday
	// )
	// fmt.Println(Friday)

	// const (
	// 	_ = 1 << (10 * iota)
	// 	kb
	// 	mb
	// 	gb
	// 	tb
	// 	pb
	// 	eb
	// 	zb
	// 	yb
	// )

	// fmt.Println(tb)
	// fmt.Println(pb)
	// fmt.Println(eb)
	// fmt.Println(zb)

}

// 3.10
// func comma(s string) string {
// 	n := len(s)
// 	if n <= 3 {
// 		return s
// 	}

// 	var buf bytes.Buffer
// 	first := n % 3
// 	if first == 0 {
// 		first = 3
// 	}
// 	buf.WriteString(s[:first])
// 	for i := first; i < n; i += 3 {
// 		buf.WriteByte(',')
// 		buf.WriteString(s[i : i+3])
// 	}

// 	return buf.String()
// }

// func comma(s string) string {
// 	n := len(s)
// 	if n <= 3 {
// 		return s
// 	}
// 	return comma(s[:n-3]) + ", " + s[n-3:]
// }

// убирает компоненты каталога и суффикс типа файла
/*
func basename(s string) string {
	// Отбрасываем / и все перед ним
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	// Отбрасываем . и все перед ним
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}*/

/*
func basename2(s string) string {
	slash := strings.LastIndex(s, "/") // -1 если / не найден
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}
*/
