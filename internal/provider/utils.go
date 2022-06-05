package provider

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/mattermost/mattermost-server/v6/model"
)

var random *rand.Rand

func fmtErr(resp *model.Response, err error) error {
	if resp == nil {
		return err
	}

	return fmt.Errorf("request %s failed with status %d: %v", resp.RequestId, resp.StatusCode, err)
}

func expandStringMap(m map[string]interface{}) map[string]string {
	r := make(map[string]string, len(m))
	for k, v := range m {
		r[k] = fmt.Sprintf("%v", v) // works for most types, unlike v.(string)
	}

	return r
}

func init() {
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
}

var passwordChars = struct {
	Lower    string
	Upper    string
	Digits   string
	Specials string
}{
	"abcdefghijklmnopqrstuvwxyz",
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
	"0123456789",
	"!@#$%^&*()-+",
}

func genPassword(size, numSpecials, numDigits, numUpper int) string {
	var password strings.Builder

	for i := 0; i < numSpecials; i++ {
		password.WriteString(
			string(passwordChars.Specials[random.Intn(len(passwordChars.Specials))]),
		)
	}
	for i := 0; i < numDigits; i++ {
		password.WriteString(
			string(passwordChars.Digits[random.Intn(len(passwordChars.Digits))]),
		)
	}
	for i := 0; i < numUpper; i++ {
		password.WriteString(
			string(passwordChars.Upper[random.Intn(len(passwordChars.Upper))]),
		)
	}
	for i := 0; i < size-numSpecials-numDigits-numUpper; i++ {
		password.WriteString(
			string(passwordChars.Lower[random.Intn(len(passwordChars.Lower))]),
		)
	}
	runes := []rune(password.String())
	random.Shuffle(len(runes), func(i, j int) {
		runes[i], runes[j] = runes[j], runes[i]
	})
	return string(runes)
}
