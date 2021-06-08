package configs

import (
	"fmt"
	"testing"
)

func TestLoad(t *testing.T) {
	if err := load(); err != nil {
		t.Error(err)
	}
	fmt.Println(Data)
}
