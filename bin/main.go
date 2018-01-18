// WIP
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/abetomo/vsq.go"
	"os"
)

func main() {
	operation := flag.String("operation", "", "operation")
	dbFilePath := flag.String("db", "", "Path of DB file")
	value := flag.String("value", "", "Data to be added (Used with unshift, push and send operations)")
	id := flag.String("id", "", "Id of the data to delete (Only used for delete operation)")
	flag.Parse()

	if *operation == "" {
		flag.Usage()
		os.Exit(1)
	}
	if *dbFilePath == "" {
		flag.Usage()
		os.Exit(1)
	}

	var vsqBasic vsq.VerySimpleQueue
	var vsqSqs vsq.VerySimpleQueueLikeSQS

	switch *operation {
	case "unshift":
		if *value == "" {
			flag.Usage()
			os.Exit(10)
		}
		vsqBasic.Load(*dbFilePath)
		vsqBasic.Unshift(*value)

	case "push":
		if *value == "" {
			flag.Usage()
			os.Exit(10)
		}
		vsqBasic.Load(*dbFilePath)
		vsqBasic.Push(*value)

	case "shift":
		vsqBasic.Load(*dbFilePath)
		v, _ := vsqBasic.Shift()
		fmt.Println(v)

	case "pop":
		vsqBasic.Load(*dbFilePath)
		v, _ := vsqBasic.Pop()
		fmt.Println(v)

	case "send":
		if *value == "" {
			flag.Usage()
			os.Exit(10)
		}
		vsqSqs.Load(*dbFilePath)
		fmt.Println(vsqSqs.Send(*value, vsq.UniqId))

	case "receive":
		vsqSqs.Load(*dbFilePath)
		v, _ := vsqSqs.Receive()
		bytes, _ := json.Marshal(v)
		fmt.Println(string(bytes))

	case "delete":
		if *id == "" {
			flag.Usage()
			os.Exit(10)
		}
		vsqSqs.Load(*dbFilePath)
		fmt.Println(vsqSqs.Delete(*id))

	default:
		flag.Usage()
		os.Exit(20)
	}
}
