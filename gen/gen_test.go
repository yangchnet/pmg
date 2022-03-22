package gen

import (
	"pmg/pass"
	"testing"
)

func Test_gen(t *testing.T) {
	any := New(pass.Any)
	t.Log(any.Password())

	easy := New(pass.Easy)
	t.Log(easy.Password())

	normal := New(pass.Normal)
	t.Log(normal.Password())

	hard := New(pass.Hard)
	t.Log(hard.Password())
}
