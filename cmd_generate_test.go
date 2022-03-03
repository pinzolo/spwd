package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

const N = 100

func TestCmdGenerateWithDefaultLength(t *testing.T) {
	got, err := execCmdGenerate()
	if err != nil {
		t.Error(err)
	}
	if len(got) != defaultPasswordLength {
		t.Errorf("%s has invalid length: %d", got, len(got))
	}
}

func TestCmdGenerateWithValidLength(t *testing.T) {
	genOpt.length = 32
	defer func() {
		genOpt.length = defaultPasswordLength
	}()

	got, err := execCmdGenerate()
	if err != nil {
		t.Error(err)
	}
	if len(got) != genOpt.length {
		t.Errorf("%s has invalid length: %d", got, len(got))
	}
}

func TestCmdGenerateWithInvalidLength(t *testing.T) {
	lens := []int{-1, 7, 129}
	defer func() {
		genOpt.length = defaultPasswordLength
	}()

	for _, l := range lens {
		t.Run(fmt.Sprintf("length: %d", l), func(t *testing.T) {
			genOpt.length = l
			_, err := execCmdGenerate()
			if err == nil {
				t.Error("error should be raised")
			} else if err.Error() != "length must be 8 to 128" {
				t.Errorf("invalid error: %s", err.Error())
			}
		})
	}
}

func TestCmdGenerateWithNoUpperCase(t *testing.T) {
	genOpt.noUpperCase = true
	defer func() {
		genOpt.noUpperCase = false
	}()

	for i := 0; i < N; i++ {
		got, err := execCmdGenerate()
		if err != nil {
			t.Error(err)
		}
		pwd := []rune(got)
		if containsAnyRune(pwd, upperRunes) {
			t.Errorf("%s contains uppercase letter", got)
		}
		if !containsAnyRune(pwd, lowerRunes) {
			t.Errorf("%s must contain lowercase letter", got)
		}
		if !containsAnyRune(pwd, numberRunes) {
			t.Errorf("%s must contain number letter", got)
		}
		if !containsAnyRune(pwd, symbolRunes) {
			t.Errorf("%s must contain symbol letter", got)
		}
	}
}

func TestCmdGenerateWithNoLowerCase(t *testing.T) {
	genOpt.noLowerCase = true
	defer func() {
		genOpt.noLowerCase = false
	}()

	for i := 0; i < N; i++ {
		got, err := execCmdGenerate()
		if err != nil {
			t.Error(err)
		}
		pwd := []rune(got)
		if !containsAnyRune(pwd, upperRunes) {
			t.Errorf("%s must contain uppercase letter", got)
		}
		if containsAnyRune(pwd, lowerRunes) {
			t.Errorf("%s contains lowercase letter", got)
		}
		if !containsAnyRune(pwd, numberRunes) {
			t.Errorf("%s must contain number letter", got)
		}
		if !containsAnyRune(pwd, symbolRunes) {
			t.Errorf("%s must contain symbol letter", got)
		}
	}
}

func TestCmdGenerateWithNoNumber(t *testing.T) {
	genOpt.noNumber = true
	defer func() {
		genOpt.noNumber = false
	}()

	for i := 0; i < N; i++ {
		got, err := execCmdGenerate()
		if err != nil {
			t.Error(err)
		}
		pwd := []rune(got)
		if !containsAnyRune(pwd, upperRunes) {
			t.Errorf("%s must contain uppercase letter", got)
		}
		if !containsAnyRune(pwd, lowerRunes) {
			t.Errorf("%s must contain lowercase letter", got)
		}
		if containsAnyRune(pwd, numberRunes) {
			t.Errorf("%s contains number letter", got)
		}
		if !containsAnyRune(pwd, symbolRunes) {
			t.Errorf("%s must contain symbol letter", got)
		}
	}
}

func TestCmdGenerateWithNoSymbol(t *testing.T) {
	genOpt.noSymbol = true
	defer func() {
		genOpt.noSymbol = false
	}()

	for i := 0; i < N; i++ {
		got, err := execCmdGenerate()
		if err != nil {
			t.Error(err)
		}
		pwd := []rune(got)
		if !containsAnyRune(pwd, upperRunes) {
			t.Errorf("%s must contain uppercase letter", got)
		}
		if !containsAnyRune(pwd, lowerRunes) {
			t.Errorf("%s must contain lowercase letter", got)
		}
		if !containsAnyRune(pwd, numberRunes) {
			t.Errorf("%s must contain number letter", got)
		}
		if containsAnyRune(pwd, symbolRunes) {
			t.Errorf("%s contains symbol letter", got)
		}
	}
}

func TestCmdGenerateWithNoAllCharTypes(t *testing.T) {
	genOpt.noUpperCase = true
	genOpt.noLowerCase = true
	genOpt.noNumber = true
	genOpt.noSymbol = true
	defer func() {
		genOpt.noUpperCase = false
		genOpt.noLowerCase = false
		genOpt.noNumber = false
		genOpt.noSymbol = false
	}()

	_, err := execCmdGenerate()
	if err == nil {
		t.Error("error should be raised")
	} else if err.Error() != "one character type is required" {
		t.Errorf("invalid error: %s", err.Error())
	}
}

func execCmdGenerate() (string, error) {
	out := &bytes.Buffer{}
	ctx := newContext(out, "generate")
	err := cmdGenerate.Run(ctx, []string{})
	if err != nil {
		return "", err
	}
	got := strings.TrimRight(out.String(), "\n")
	return got, nil
}
