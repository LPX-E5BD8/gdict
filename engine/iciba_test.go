package engine

import (
	"fmt"
	"testing"

	"github.com/kr/pretty"
)

func TestNewIciba(t *testing.T) {
	i := NewIciba("test", "dark")
	pretty.Println(i.result)
}

func TestIcibaResult_Format(t *testing.T) {
	i := NewIciba("test", "dark")
	fmt.Println(i.Query())
	i = NewIciba("测试", "dark")
	fmt.Println(i.Query())
}
