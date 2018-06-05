package engien

import "testing"

func TestYoudao_Query(t *testing.T) {
	NewYoudao("test", "dark").Query()
	NewYoudao("测试", "light").Query()
}
