package task4

import "fmt"

// MainTask4
/* 	Сделать конвейер чисел.

	Даны два канала. В первый пишутся числа. Нужно, чтобы числа читались из первого по мере
	поступления, что-то с ними происходило (допустим, возводились в квадрат) и результат
	записывался во второй канал.

	Довольно частая задача, более подробно можно почитать тут https://blog.golang.org/pipelines.

Итого: Решается довольно прямолинейно — запускаем две горутины. В одной пишем в первый канал.
	Во второй читаем из первого канала и пишем во второй. Главное — не забыть закрыть каналы, чтобы
	ничего нигде не заблокировалось.
*/
func MainTask4() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 0; x <= 10; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	for x := range squares {
		fmt.Println(x)
	}
}
