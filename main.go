package main

import (
	"bufio"
	"fmt"
	"github.com/go-resty/resty/v2"
	"os"
	"strconv"
	"strings"
	"time"
)

func shoot(i int64) {
	client := resty.New()

	resp, _ := client.R().
		EnableTrace().
		Get("https://httpbin.org/get")

	//fmt.Println(resp)
	fmt.Println(fmt.Sprintf("%d: %d", i, resp.StatusCode()))
}

func stressExecutor(c chan int64) {
	controlStress := int64(0)

	for {
		select {
		case stress := <-c:
			controlStress = stress
			continue

		default:
			for i := int64(0); i < controlStress; i++ {
				//fmt.Println(i)
				go shoot(i)
			}
			time.Sleep(time.Second)
		}
	}
}

func receiveStress(c chan int64) {
	for {
		fmt.Println("Requests per Second: ")
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

		// update stress thorough channel
		c <- stress
	}
}



func main() {
	stressC := make(chan int64)

	go receiveStress(stressC)
	go stressExecutor(stressC)

	select {}
}
