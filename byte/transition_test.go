package byte

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	fakeString = "test"
	fakeBytes  = []byte(fakeString)
)

func TestBytes2str(t *testing.T) {
	if !assert.Equal(t, Bytes2str(fakeBytes), fakeString) {
		t.FailNow()
	}
}

func BenchmarkBytes2str(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		Bytes2str(fakeBytes)
	}
}

func BenchmarkSystemStr2bytes(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		_ = []byte(fakeString)
	}
}
func BenchmarkSystemBytes2Str(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		_ = string(fakeBytes)
	}
}
