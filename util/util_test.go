package util

import "testing"

// 普通的测试
func TestGenShortID(t *testing.T) {
	shortID, err := GenShortID()
	if shortID == "" || err != nil {
		t.Error("GenShortID failed")
	}
}

// 性能测试
func BenchmarkGenShortID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenShortID()
	}
}

func BenchmarkGenShortIDTimeConsuming(b *testing.B) {
	b.StopTimer() // 调用该函数停止压力测试的时间计数

	shortID, err := GenShortID()
	if shortID == "" || err != nil {
		b.Error(err)
	}

	b.StartTimer() // 重新开始时间

	for i := 0; i < b.N; i++ {
		GenShortID()
	}
}
