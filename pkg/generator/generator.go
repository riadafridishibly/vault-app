package generator

import (
	"crypto/sha256"
	"encoding/binary"
	"errors"
	"io"

	"github.com/riadafridishibly/vault-app/pkg/generator/common"
	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/hkdf"
)

type Generator struct {
	common.Options
	charSet    []byte
	repeatMap  map[byte]bool
	buf        [2]byte
	hkdfReader io.Reader
}

func (g *Generator) init() {
	if g.charSet != nil {
		return
	}
	g.charSet = make([]byte, 0, len(common.AllChars))
	g.charSet = append(g.charSet, []byte(common.LowerAscii)...)
	g.charSet = append(g.charSet, []byte(common.UpperAscii)...)
	g.charSet = append(g.charSet, []byte(common.Numbers)...)

	if g.IncludeSymbols {
		g.charSet = append(g.charSet, []byte(common.Spaces)...)
		g.charSet = append(g.charSet, []byte(common.Symbols)...)
	}

	g.repeatMap = make(map[byte]bool)
	if g.Length <= 0 {
		g.Length = 20
	}

	if !g.AllowRepeat && g.Length >= len(g.charSet) {
		panic("char set length is less than password length: unique not possible")
	}
}

func (g *Generator) Generate(seed, password []byte) (string, error) {
	password = argon2.IDKey(password, []byte("Weird-Password-Seed"), 1, 64*1024, 4, 32)
	pass, err := g.generate(seed, password)
	if err != nil {
		return "", err
	}
	return string(pass), nil
}

func (g *Generator) getPosition() (int, error) {
	_, err := io.ReadFull(g.hkdfReader, g.buf[:])
	if err != nil {
		return -1, err
	}
	return int(binary.BigEndian.Uint16(g.buf[:])), nil
}

func (g *Generator) getChar() (byte, error) {
	pos, err := g.getPosition()
	if err != nil {
		return 0, err
	}
	return g.charSet[int(pos)%len(g.charSet)], nil
}

func (g *Generator) getUniqueChar() (byte, error) {
	ch, err := g.getChar()
	if err != nil {
		return ch, err
	}

	for {
		if len(g.repeatMap) == len(g.charSet) {
			return 0, errors.New("unique not possible: all chars used")
		}
		if _, found := g.repeatMap[ch]; !found {
			g.repeatMap[ch] = true
			return ch, nil
		}

		ch, err = g.getChar()
		if err != nil {
			return ch, err
		}
	}
}

func (g *Generator) generate(seed, password []byte) ([]byte, error) {
	g.init()
	g.hkdfReader = hkdf.New(sha256.New, password, seed, []byte("WPG-info"))

	outputPassword := make([]byte, g.Length)

	var spacePos int = -1
	if g.IncludeSpace {
		pos, err := g.getPosition()
		if err != nil {
			return nil, err
		}
		spacePos = pos%(g.Length-2) + 1
		g.repeatMap[' '] = true
	}

	getCharFn := g.getUniqueChar
	if g.AllowRepeat {
		getCharFn = g.getChar
	}

	for i := 0; i < g.Length; i++ {
		char, err := getCharFn()
		if err != nil {
			return nil, err
		}
		outputPassword[i] = char
		if spacePos == i {
			outputPassword[i] = ' '
		}
	}
	return outputPassword, nil
}
