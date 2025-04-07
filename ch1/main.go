package main

import (
	"fmt"
	"log"
	"net/http"
)

// fmt -- для форматирования ввода и вывода
// os -- пакет для работы с операционной системой
// bufio -- буферизованный ввод и вывод обертывает в Reader или Writer

func main() {
	// Тема 1.1 Hello world

	//fmt.Println("Hello World!")

	// Тема 1.2 Аргументы командной строки

	// echo1
	/*
		var s, sep string
		for i := 1; i < len(os.Args); i++ {
			s += sep + os.Args[i]
			sep = " "
		}
		fmt.Println(s)
	*/

	// echo2
	/*
		s, sep := "", ""
		for _, arg := range os.Args[1:] {
			s += sep + arg
			sep = " "
		}
		fmt.Println(s)
	*/

	// echo3
	// fmt.Println(strings.Join(os.Args[1:], " "))

	// ex 1.1
	// fmt.Println(strings.Join(os.Args, " "))

	// ex 1.2
	/*
		for i, arg := range os.Args {
			fmt.Println(i, arg)
		}
	*/

	// ex 1.3
	// -----

	// Тема 1.3 Поиск повторяющихся строк

	// Dup1 - выводит текст каждой строки, которая появляется в STDIN более одного раза,
	// а также количество появлений

	// Ключ в map должен быть сравним с помощью ==
	// Map - неупорядоченная стркутура данных
	// Время операций с map -- константное время

	/*
		counts := make(map[string]int)
		input := bufio.NewScanner(os.Stdin) // -- возвращает сканнер входного потока
		for input.Scan() {                  // -- считывает строку и удаляет символ новой строки, возвращает true или false
			counts[input.Text()]++			// -- input.Text() - возвращает прочитанную строку
		}
		// Примечание: игнорируем потенциальные ошибки из input.Err()
		for line, n := range counts {
			if n > 1 {
				fmt.Printf("%d\t%s\n", n, line)
			}
		}
	*/

	// Dup2 -- читает входной поток или список именовынных файлов и выводит дубликаты

	/*
		counts := make(map[string]int)
		files := os.Args[1:] // - файлы
		if len(files) == 0 {
			countLines(os.Stdin, counts) // если файлов нет, то передаётся входной поток
		} else {
			for _, arg := range files {
				f, err := os.Open(arg) // открытие файла по аргументу
				if err != nil {
					fmt.Fprintf(os.Stderr, "dup2: %v\n", err) // вывод ошибки на поток ошибок
					continue
				}
				countLines(f, counts)
				f.Close()
			}
		}
		for line, n := range counts {
			if n > 1 {
				fmt.Printf("%d\t%s\n", n, line)
			}
		}
	*/

	// dup3 -- считывает данные в память и затем считает дубликаты
	/*
		counts := make(map[string]int)
		for _, filename := range os.Args[1:] {
			data, err := os.ReadFile(filename) // readfile возвращает срез байт и ошибку
			if err != nil {
				fmt.Fprintf(os.Stderr, "dop3: %v/n", err)
				continue
			}
			for _, line := range strings.Split(string(data), "\n") {
				counts[line]++
			}
		}
		for line, n := range counts {
			if n > 1 {
				fmt.Printf("%d\t%s\n", n, line)
			}
		}*/

	// ex 1.4
	/*
		counts := make(map[string]int)
		countFiles := make(map[string]string)
		files := os.Args[1:] // - файлы
		if len(files) == 0 {
			countLines(os.Stdin, counts, countFiles, "") // если файлов нет, то передаётся входной поток
		} else {
			for _, arg := range files {
				f, err := os.Open(arg) // открытие файла по аргументу
				if err != nil {
					fmt.Fprintf(os.Stderr, "dup2: %v\n", err) // вывод ошибки на поток ошибок
					continue
				}
				countLines(f, counts, countFiles, arg)
				f.Close()
			}
		}
		for line, n := range counts {
			if n > 1 {
				fmt.Printf("%d\t%s\t--%s\n", n, line, countFiles[line])
			}
		}
	*/

	// Тема 1.4  Анимированные GIF-изображения
	// --------

	// Тема 1.5 Выборка URL

	//Fetch
	/*
		for _, url := range os.Args[1:] {
			resp, err := http.Get(url) // -- http запрос
			if err != nil {
				fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
				os.Exit(1)
			}
			b, err := io.ReadAll(resp.Body) // -- resp.Body - содержит ответ сервера в виде потока, доступного для чтения
			resp.Body.Close()
			if err != nil {
				fmt.Fprintf(os.Stderr, "fetch: чтение %s: %v/n", url, err)
				os.Exit(1)
			}
			fmt.Printf("%s", b)
		}
	*/

	// ex 1.7
	/*
		for _, url := range os.Args[1:] {
			resp, err := http.Get(url)
			if err != nil {
				fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
				os.Exit(1)
			}
			_, err = io.Copy(os.Stdout, resp.Body) //
			resp.Body.Close()
			if err != nil {
				fmt.Fprintf(os.Stderr, "fetch: чтение %s: %v/n", url, err)
				os.Exit(1)
			}

		}
	*/
	// ex 1.8
	/*
		for _, url := range os.Args[1:] {
			if !strings.HasPrefix(url, "http://") {
				url = "http://" + url
			}
			resp, err := http.Get(url)
			if err != nil {
				fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
				os.Exit(1)
			}
			b, err := io.ReadAll(resp.Body)
			resp.Body.Close()
			if err != nil {
				fmt.Fprintf(os.Stderr, "fetch: чтение %s: %v/n", url, err)
				os.Exit(1)
			}
			fmt.Printf("%s", b)
		}
	*/
	// ex 1.9
	/*
		for _, url := range os.Args[1:] {
			resp, err := http.Get(url)
			if err != nil {
				fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
				os.Exit(1)
			}
			fmt.Println("Status: ", resp.Status)
			b, err := io.ReadAll(resp.Body)
			resp.Body.Close()
			if err != nil {
				fmt.Fprintf(os.Stderr, "fetch: чтение %s: %v/n", url, err)
				os.Exit(1)
			}
			fmt.Printf("%s", b)
		}
	*/

	// Тема 1.6 Параллельная выборка URL

	//Fetchall - выполняет параллельную выборку URL и сообщает о затраченном времени и размере ответа для каждого из них
	/*
		start := time.Now()
		ch := make(chan string) // канал -- механихм связи, который позволяет одной go-подпрограмме
		// передавать значения определенного типа другой go-подпрограмме.
		for _, url := range os.Args[1:] {
			go fetch(url, ch) // Запуск go-программы
			// go-попрограмма --- представляет собой параллельное выполнение функции
		}
		for range os.Args[1:] {
			fmt.Println(<-ch)
		}
		fmt.Printf("%.2fs elapsed/n", time.Since(start).Seconds()
	*/

	// ex 1.10
	// --------
	// ex 1.11
	// -------

	// Тема 1.7 Веб-сервер

	// server1 - минимальный "echo" - server

	http.HandleFunc("/", handler) // Каждый запрос вызывает обработчик
	log.Fatal(http.ListenAndServe("localhost:8000", nil))

}

// обработчик возвращает комронент пути из URL запроса
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.PATH = %q\n", r.URL.Path)
}

/* для fetchall
func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // Отправка в канад ch
		return
	}
	nbytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
*/
// Функция для dup2 -- считает дубликаты
/*
func countLines(f *os.File, counts map[string]int, countFiles map[string]string, arg string) { // *os.File можно использовать как io.Reader
	input := bufio.NewScanner(f)
	//sep := " "
	for input.Scan() {
		counts[input.Text()]++
		для ex 1.4
		countFiles[input.Text()] += sep + arg
		sep = " "
	}
	// Примечание: игнорируем ошибки из input.Err()
}
*/
