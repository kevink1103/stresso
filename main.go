package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)


func stressExecutor(c chan int64) {
	controlStress := int64(1)

	for {
		select {
		case stress := <-c:
			controlStress = stress
			continue

		default:
			for i := int64(0); i < controlStress; i++ {
				fmt.Println(i)
			}
			time.Sleep(time.Second)
		}
	}
}

func receiveStress(c chan int64) {
	for {
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("text read error:", err)
			return
		}

		text = strings.Replace(text, "\n", "", -1)

		stress, err := strconv.ParseInt(text, 10, 64)
		if err != nil {
			fmt.Println("stress integer parsing error:", err)
			return
		}

		fmt.Println("stress:", stress)

		c <- stress
	}
}


func main() {
	stressC := make(chan int64)

	go receiveStress(stressC)
	go stressExecutor(stressC)

	select {}
}
