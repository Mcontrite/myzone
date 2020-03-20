package regex

import "regexp"

type Regex struct {
	Str     string
	Pattern string
}

func (this *Regex) VerifyString() bool {
	reg := regexp.MustCompile(this.Pattern)
	return reg.MatchString(this.Str)
}
