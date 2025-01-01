package benchmark

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/garrettladley/strumbra"
)

var inputLengths = []int{4, 8, 12, 16, 32, 64}

func BenchmarkCmpRandom(b *testing.B) {
	for _, length := range inputLengths {
		b.Run(fmt.Sprintf("String-%d", length), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				a, _b := randomString(length), randomString(length)
				_ = a < _b
			}
		})

		b.Run(fmt.Sprintf("UmbraString-%d", length), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				a, _ := strumbra.New(randomString(length))
				_b, _ := strumbra.New(randomString(length))
				_ = a.Compare(_b)
			}
		})
	}
}

func BenchmarkCmpSame(b *testing.B) {
	for _, length := range inputLengths {
		b.Run(fmt.Sprintf("String-%d", length), func(b *testing.B) {
			s := randomString(length)
			for i := 0; i < b.N; i++ {
				_ = s < s
			}
		})

		b.Run(fmt.Sprintf("UmbraString-%d", length), func(b *testing.B) {
			s, _ := strumbra.New(randomString(length))
			for i := 0; i < b.N; i++ {
				_ = s.Compare(s)
			}
		})
	}
}

func BenchmarkEqRandom(b *testing.B) {
	for _, length := range inputLengths {
		b.Run(fmt.Sprintf("String-%d", length), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				a, _b := randomString(length), randomString(length)
				_ = a == _b
			}
		})

		b.Run(fmt.Sprintf("UmbraString-%d", length), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				a, _ := strumbra.New(randomString(length))
				_b, _ := strumbra.New(randomString(length))
				_ = a == _b
			}
		})
	}
}

func BenchmarkEqSame(b *testing.B) {
	for _, length := range inputLengths {
		b.Run(fmt.Sprintf("String-%d", length), func(b *testing.B) {
			s := randomString(length)
			for i := 0; i < b.N; i++ {
				_ = s == s
			}
		})

		b.Run(fmt.Sprintf("UmbraString-%d", length), func(b *testing.B) {
			s, _ := strumbra.New(randomString(length))
			for i := 0; i < b.N; i++ {
				_ = s == s
			}
		})
	}
}

func BenchmarkEqualsRandom(b *testing.B) {
	for _, length := range inputLengths {
		b.Run(fmt.Sprintf("String-%d", length), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				a, _b := randomString(length), randomString(length)
				_ = a == _b
			}
		})

		b.Run(fmt.Sprintf("UmbraString-%d", length), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				a, _ := strumbra.New(randomString(length))
				_b, _ := strumbra.New(randomString(length))
				_ = a.Equals(_b)
			}
		})
	}
}

func BenchmarkEqualsSame(b *testing.B) {
	for _, length := range inputLengths {
		b.Run(fmt.Sprintf("String-%d", length), func(b *testing.B) {
			s := randomString(length)
			for i := 0; i < b.N; i++ {
				_ = s == s
			}
		})

		b.Run(fmt.Sprintf("UmbraString-%d", length), func(b *testing.B) {
			s, _ := strumbra.New(randomString(length))
			for i := 0; i < b.N; i++ {
				_ = s.Equals(s)
			}
		})
	}
}

func BenchmarkConstructEmpty(b *testing.B) {
	b.Run("UmbraString", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = strumbra.New("")
		}
	})
}

func BenchmarkConstructNonEmpty(b *testing.B) {
	for _, length := range inputLengths {
		b.Run(fmt.Sprintf("UmbraString-%d", length), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = strumbra.New(randomString(length))
			}
		})
	}
}

const charset string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func randomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
