package clients

import "fmt"

func PrintMethodSucceed(methodName string, condition bool, prefix string) {
	if prefix == "" {
		prefix = "client"
	}
	if condition {
		fmt.Printf("[%s] %s Success\n", prefix, methodName)
	} else {
		fmt.Printf("[%s] %s Failed\n", prefix, methodName)
	}
}
