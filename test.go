package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func getOtherHalf(i int) string {
	app := "1/2HERE is the message "
	return app
}

func getMessage(i int) {
	fmt.Println("HERE is the message ", i)
	fmt.Println(getOtherHalf(i), i)
	wg.Done()
}

/*func main() {
	for index := 0; index < 5; index++ {
		wg.Add(1)
		go getMessage(index)
		wg.Wait()
	}

}
*/
