// Go поддерживает <em><a href="http://en.wikipedia.org/wiki/Pointer_(computer_programming)">указатели</a></em>,
// позволяя вам передавать ссылки на значения и записи
// в вашей программе.

package main

import "fmt"

// Мы покажем, как работают указатели, на примере 2 функций:
// `zeroval` и `zeroptr`. В `zeroval` определен только 1
// аргумент с типом `int`, который передается по значению.
// `zeroval` получает копию `ival` при вызове функции.
func zeroval(ival int) {
	ival = 0
}

// `zeroptr` получает в качестве аргумента параметр `*int`,
// который является указателем на `int`. Запись `*iptr` в
// теле функции _разыменовывает_ указатель с его адреса
// памяти на текущее значение по этому адресу. Присвоение
// значения разыменованному указателю изменяет значение
// по указанному адресу.
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// Запись `&i` получает ссылку на область памяти, в
	// которой хранится `i`, т.е. указатель на `i`.
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// Указатели могут быть выведены на экран
	fmt.Println("pointer:", &i)
}
