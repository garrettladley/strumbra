package benchmark

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/garrettladley/stumbra"
)

var inputLengths = []int{4, 8, 12, 16, 32, 64}

func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func BenchmarkCmpRandom(b *testing.B) {
	for _, length := range inputLengths {
		b.Run(fmt.Sprintf("String-%d", length), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				a, c := randomString(length), randomString(length)
				_ = a < c
			}
		})

		b.Run(fmt.Sprintf("UmbraString-%d", length), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				a, _ := stumbra.New(randomString(length))
				c, _ := stumbra.New(randomString(length))
				_ = a.Compare(c)
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
			s, _ := stumbra.New(randomString(length))
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
				a, c := randomString(length), randomString(length)
				_ = a == c
			}
		})

		b.Run(fmt.Sprintf("UmbraString-%d", length), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				a, _ := stumbra.New(randomString(length))
				c, _ := stumbra.New(randomString(length))
				_ = a.Equal(c)
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
			s, _ := stumbra.New(randomString(length))
			for i := 0; i < b.N; i++ {
				_ = s.Equal(s)
			}
		})

	}
}

func BenchmarkConstructEmpty(b *testing.B) {
	b.Run("UmbraString", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_, _ = stumbra.New("")
		}
	})

}

func BenchmarkConstructNonEmpty(b *testing.B) {
	for _, length := range inputLengths {
		b.Run(fmt.Sprintf("UmbraString-%d", length), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, _ = stumbra.New(randomString(length))
			}
		})

	}
}
