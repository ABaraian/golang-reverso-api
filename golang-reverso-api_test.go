package reverso

import (
	"fmt"
	"testing"
)

func TestGet_Translation(t *testing.T) {
	result := Get_Translation("english", "french", "Let me see you")
	fmt.Println(result.ContextResults)
	if result.ContextResults == nil {
		fmt.Println("NONE1")
	}
	result = Get_Translation("english", "thai", "how")
	fmt.Println(result)
	if result.ContextResults == nil {
		fmt.Println("NONE")
	}
}
