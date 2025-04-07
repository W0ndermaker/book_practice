package main

import (
	"conv/tempconv"
	"fmt"
	"os"
	"strconv"
)

//var n = flag.Bool("n", false, "пропуск символа новой строки")
//var sep = flag.String("s", " ", "разделитель")

func main() {

	// 2.3.2 Указатели
	/*
		flag.Parse()
		fmt.Print(strings.Join(flag.Args(), *sep))
		if !*n {
			fmt.Println()
		}
	*/

	// 2.3.3 Функция new
	/*
		p := new(int) // создаёт неименованную переменую типа int с нулевым значением
		// и возвращает указатель на эту переменную
		fmt.Println(*p)
		*p = 2
		fmt.Println(*p)
	*/

	/*
		v, ok := m[key] // поиск в отображении
		v, ok := x.(T) // утверждение о типе
		v, ok = <-ch // получение из канала
	*/
	/*
		fmt.Printf("%g\n", tempconv.BoilingC)
		boilingF := tempconv.CToF(tempconv.BoilingC)
		fmt.Printf("%g\n", boilingF)
		fmt.Printf("%g\n", tempconv.CToF(tempconv.FreezingC))
		c := tempconv.FToC(212.0)
		fmt.Println(c.String())
		fmt.Printf("%v\n", c)
		fmt.Printf("%s\n", c)
		fmt.Println()
		fmt.Printf("Бррр! %v\n", tempconv.AbsoluteZeroC)
		fmt.Println(tempconv.BoilingC)
		fmt.Println(tempconv.CToF(tempconv.BoilingC))

		kelv := tempconv.Kelvin(0)
		fmt.Println(kelv)
		fmt.Println(tempconv.KToC(kelv))
		fmt.Println(tempconv.KToF(kelv))
	*/
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		k := tempconv.Kelvin(t)
		fmt.Printf("%s = %s\n", f, tempconv.FToC(f))
		fmt.Printf("%s = %s\n", f, tempconv.FToK(f))
		fmt.Printf("%s = %s\n", c, tempconv.CToF(c))
		fmt.Printf("%s = %s\n", c, tempconv.CToK(c))
		fmt.Printf("%s = %s\n", k, tempconv.KToC(k))
		fmt.Printf("%s = %s\n", k, tempconv.KToF(k))
	}

}
