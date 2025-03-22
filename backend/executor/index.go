package executor

import "fmt"

func ExecuteWorkflow(key string) error {
	fmt.Println("Executing workflow with key:", key)
	return nil
}
