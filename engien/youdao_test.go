package engien

import "testing"

func TestYoudao_Query(t *testing.T) {
	NewYoudao("test").Query()
	NewYoudao("测试").Query()
}
