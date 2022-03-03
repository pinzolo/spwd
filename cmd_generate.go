package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

var cmdGenerate = &Command{
	Run:       runGenerate,
	UsageLine: "generate [-l <length>] [-U] [-L] [-N] [-S]",
	Short:     "Generate new password",
	Long: `Generate new password 
Some rules can be specified with flags.

OPTIONS
        -l, --length <length>
            Generates a password of the length specified in <value>.
            Values from 8 to 128 can be specified. (default: 16)
     
        -U, --no-uppercase
            Passwords to be generated do not contain uppercase letters.
            Default value is false.
     
        -L, --no-lowercase
            Passwords to be generated do not contain lowercase letters.
            Default value is false.
     
        -N, --no-number
            Passwords to be generated do not contain number letters.
            Default value is false.
     
        -S, --no-symbol
            Passwords to be generated do not contain symbol letters.
            Default value is false.
`,
}

const (
	defaultPasswordLength = 16
	minPasswordLength     = 8
	maxPasswordLength     = 128
)

type cmdGenerateOption struct {
	length      int
	noUpperCase bool
	noLowerCase bool
	noNumber    bool
	noSymbol    bool
}

func (o cmdGenerateOption) validate() error {
	if o.length < minPasswordLength || maxPasswordLength < o.length {
		return fmt.Errorf("length must be %d to %d", minPasswordLength, maxPasswordLength)
	}

	if o.noUpperCase && o.noLowerCase && o.noNumber && o.noSymbol {
		return errors.New("one character type is required")
	}

	return nil
}

var (
	genOpt      = cmdGenerateOption{}
	upperRunes  = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	lowerRunes  = []rune("abcdefghijklmnopqrstuvwxyz")
	numberRunes = []rune("0123456789")
	symbolRunes = []rune("!@#$%^&*-=+?")
)

func init() {
	cmdGenerate.Flag.IntVar(&genOpt.length, "l", defaultPasswordLength, "Length of password")
	cmdGenerate.Flag.IntVar(&genOpt.length, "length", defaultPasswordLength, "Length of password")
	cmdGenerate.Flag.BoolVar(&genOpt.noUpperCase, "U", false, "No uppercase letter(s)")
	cmdGenerate.Flag.BoolVar(&genOpt.noUpperCase, "no-uppercase", false, "No uppercase letter(s)")
	cmdGenerate.Flag.BoolVar(&genOpt.noLowerCase, "L", false, "No lowercase letter(s)")
	cmdGenerate.Flag.BoolVar(&genOpt.noLowerCase, "no-lowercase", false, "No lowercase letter(s)")
	cmdGenerate.Flag.BoolVar(&genOpt.noNumber, "N", false, "No number letter(s)")
	cmdGenerate.Flag.BoolVar(&genOpt.noNumber, "no-number", false, "No number letter(s)")
	cmdGenerate.Flag.BoolVar(&genOpt.noSymbol, "S", false, "No symbol letter(s)")
	cmdGenerate.Flag.BoolVar(&genOpt.noSymbol, "no-symbol", false, "No symbol letter(s)")
}

func runGenerate(ctx context, args []string) error {
	if err := genOpt.validate(); err != nil {
		return err
	}

	runes := prepareRunes()
	rand.Seed(time.Now().UnixNano())
	pwd := generatePassword(runes)
	_, err := fmt.Fprintln(ctx.out, pwd)
	return err
}

func prepareRunes() []rune {
	var runes []rune
	if !genOpt.noUpperCase {
		runes = append(runes, upperRunes...)
	}
	if !genOpt.noLowerCase {
		runes = append(runes, lowerRunes...)
	}
	if !genOpt.noNumber {
		runes = append(runes, numberRunes...)
	}
	if !genOpt.noSymbol {
		runes = append(runes, symbolRunes...)
	}
	return runes
}

func generatePassword(runes []rune) string {
	pwd := make([]rune, genOpt.length)
	for i := 0; i < genOpt.length; i++ {
		pwd[i] = runes[rand.Intn(len(runes))]
	}
	if isValidPassword(pwd) {
		return string(pwd)
	}

	return generatePassword(runes)
}

func isValidPassword(pwd []rune) bool {
	valid := true
	if !genOpt.noUpperCase {
		valid = valid && containsAnyRune(pwd, upperRunes)
	}
	if !genOpt.noLowerCase {
		valid = valid && containsAnyRune(pwd, lowerRunes)
	}
	if !genOpt.noNumber {
		valid = valid && containsAnyRune(pwd, numberRunes)
	}
	if !genOpt.noSymbol {
		valid = valid && containsAnyRune(pwd, symbolRunes)
	}
	return valid
}

func containsAnyRune(pwd, runes []rune) bool {
	for _, r1 := range pwd {
		for _, r2 := range runes {
			if r1 == r2 {
				return true
			}
		}
	}
	return false
}
