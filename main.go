package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

const client = 24
const urlTarget = "http://localhost:8000/api/v1/vault/add/voucher"
const pathPayload = "payload/payload.json"

type TaskRequest struct {
	ClientDone chan int
	Stop       chan bool
}

func main() {

	clientDone := 0

	task := &TaskRequest{
		ClientDone: make(chan int),
		Stop:       make(chan bool),
	}

	go task.GetClientDone(client)

	for i := 1; i <= client; i++ {

		go func(t *TaskRequest, pos int) {

			result, err := DoRequest(
				urlTarget,
				Read(pathPayload),
			)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(fmt.Sprintf("client number %d : ", pos), result)

			clientDone++
			t.ClientDone <- clientDone

		}(task, i)

	}

	<-task.Stop

	fmt.Println("all client stop")
}

func Read(path string) []byte {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	return b
}

func (t *TaskRequest) GetClientDone(maxClient int) {
	for {
		select {
		case c := <-t.ClientDone:
			if c == maxClient {

				t.Stop <- true
				break
			}
		default:
		}
	}
}
