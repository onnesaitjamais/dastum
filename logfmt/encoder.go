/*
#######
##           __         __
##       ___/ /__ ____ / /___ ____ _
##      / _  / _ `(_-</ __/ // /  ' \
##      \_,_/\_,_/___/\__/\_,_/_/_/_/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package logfmt

import (
	"bytes"
	"fmt"
	"strings"
	"unicode/utf8"
)

func cleanKey(r rune) rune {
	if r <= 32 || r == 61 || r == 34 || r == utf8.RuneError {
		return -1
	}

	return r
}

// Encode AFAIRE
func Encode(buf *bytes.Buffer, kv ...interface{}) {
	if len(kv)%2 == 1 {
		buf.WriteString("[>>odd-numbered list<<]")
		return
	}

	for i := 0; i < len(kv); i += 2 {
		if i != 0 {
			buf.WriteRune(32)
		}

		// key
		s, ok := kv[i].(string)
		if ok {
			buf.WriteString(strings.Map(cleanKey, s))
		} else {
			buf.WriteString(strings.Map(cleanKey, fmt.Sprintf("%#v", s)))
		}

		// '=""
		buf.WriteRune(61)

		// value
		buf.WriteString(fmt.Sprintf("%#v", kv[i+1]))
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
