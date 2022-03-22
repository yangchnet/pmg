package gen

import (
	"bytes"
	"crypto/rand"
	"math/big"
	"pmg/pass"
)

// Generator generates password
type Generator interface {
	// Password generates password string
	Password() string
}

var (
	_ Generator = (*AnyGen)(nil)
	_ Generator = (*EasyGen)(nil)
	_ Generator = (*NormalGen)(nil)
	_ Generator = (*HardGen)(nil)
)

// AnyGen generates password with strength Any
type AnyGen struct {
	characters string
}

func (g *AnyGen) Password() string {
	return generate(g.characters, 3)
}

// EasyGen generates password with strength Easy
type EasyGen struct {
	characters string
}

func (g *EasyGen) Password() string {
	return generate(g.characters, 6)
}

// NormalGen generates password with strength Normal
type NormalGen struct {
	characters string
}

func (g *NormalGen) Password() string {
	return generate(g.characters, 8)
}

// HardGen generates password with strength Hard
type HardGen struct {
	characters string
}

func (g *HardGen) Password() string {
	return generate(g.characters, 16)
}

// New return a Generator
func New(s pass.Strength) Generator {
	switch s {
	case pass.Any:
		return &AnyGen{
			characters: LowerLetters,
		}
	case pass.Easy:
		return &EasyGen{
			characters: LowerLetters + UpperLetters,
		}
	case pass.Normal:
		return &NormalGen{
			characters: LowerLetters + UpperLetters + Digits,
		}
	case pass.Hard:
		return &HardGen{
			characters: LowerLetters + UpperLetters + Digits + Symbols,
		}
	default:
		return &HardGen{
			characters: LowerLetters + UpperLetters + Digits + Symbols,
		}
	}
}

func generate(seed string, length uint32) string {
	bb := bytes.Buffer{}
	max := big.NewInt(int64(len(seed)))
	for i := 0; i < int(length); i++ {
		n, _ := rand.Int(rand.Reader, max)
		bb.WriteByte(seed[n.Int64()])
	}

	return bb.String()
}
