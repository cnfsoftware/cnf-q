package main

import (
	"flag"
	"fmt"
	"os"

	"cnf-q/pkg/queueclient"
)

const baseURL = "http://localhost:8080"

func main() {
	name := flag.String("name", "", "Queue name")
	message := flag.String("message", "", "Message to push")
	file := flag.String("file", "", "File to read message from")
	url := flag.String("url", baseURL, "Base URL of the queue service")
	token := flag.String("token", "", "Authorization token")
	help := flag.Bool("help", false, "Show available commands")
	flag.Parse()

	if *help || flag.NArg() < 1 {
		printHelp()
		return
	}

	if flag.NArg() < 1 {
		fmt.Println("Error: Missing command")
		printHelp()
		os.Exit(1)
	}

	command := flag.Arg(0)

	switch command {
	case "push":
		if *name == "" || (*message == "" && *file == "") {
			fmt.Println("Usage: cli push --name <queue_name> --message <data> OR --file <filepath>")
			os.Exit(1)
		}
		var data []byte
		if *file != "" {
			var err error
			data, err = os.ReadFile(*file)
			if err != nil {
				fmt.Println("Error reading file:", err)
				return
			}
		} else {
			data = []byte(*message)
		}
		push(*url, *name, data, *token)

	case "pop":
		if *name == "" {
			fmt.Println("Usage: cli pop --name <queue_name>")
			os.Exit(1)
		}
		pop(*url, *name, *token)

	case "peek":
		if *name == "" {
			fmt.Println("Usage: cli peek --name <queue_name>")
			os.Exit(1)
		}
		peek(*url, *name, *token)

	case "list":
		listQueues(*url, *token)

	default:
		fmt.Println("Unknown command")
		os.Exit(1)
	}
}

func printHelp() {
	fmt.Println(`Usage: cnf-q-cli <command> [options]
Commands:
  push    --name <queue_name> --message <data> OR --file <filepath>  Add data to a queue
  pop     --name <queue_name>                                        Remove and return the first element from a queue
  peek    --name <queue_name>                                        View the last element in a queue without removing it
  list                                                               List all available queues

Options:
  --name    Queue name
  --message Message to push
  --file    File path to read message from
  --url     Base URL of the queue service (default: http://localhost:8080)
  --token   Authorization token
  --help    Show this help message
`)
	os.Exit(1)
}

func push(url, queueName string, data []byte, token string) {
	client := queueclient.NewClient(url, token)
	err := client.Push(queueName, data)
	if err != nil {
		fmt.Println("Error pushing message:", err)
	}
}

func pop(url, queueName, token string) {
	client := queueclient.NewClient(url, token)
	bytes, err := client.Pop(queueName)
	if err != nil {
		fmt.Println("Error pushing message:", err)
		return
	}

	fmt.Println(string(bytes))
}

func peek(url, queueName, token string) {
	client := queueclient.NewClient(url, token)
	bytes, err := client.Peek(queueName)
	if err != nil {
		fmt.Println("Error pushing message:", err)
		return
	}

	fmt.Println(string(bytes))
}

func listQueues(url, token string) {
	client := queueclient.NewClient(url, token)
	queues, err := client.ListQueues()
	if err != nil {
		fmt.Println("Error pushing message:", err)
		return
	}

	fmt.Println("Queues:", queues)
}
