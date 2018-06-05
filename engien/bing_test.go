package engien

import (
	"fmt"
	"testing"

	"github.com/kr/pretty"
)

func TestBing_Query(t *testing.T) {
	b := NewBing("test", "dark")
	pretty.Println(b)
	fmt.Println(b.Query())
}
