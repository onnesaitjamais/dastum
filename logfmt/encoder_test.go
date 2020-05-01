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
	"testing"
)

func TestEncode(t *testing.T) {
	buf := bytes.Buffer{}

	data := []struct {
		kv   []interface{}
		want string
	}{
		// 00
		{
			[]interface{}{},
			"",
		},
		// 01
		{
			[]interface{}{nil, nil},
			"=<nil>",
		},
		// 02
		{
			[]interface{}{"age", 53},
			"age=53",
		},
		// 03
		{
			[]interface{}{"a\tb\nc", "def"},
			`abc="def"`,
		},
		// 04
		{
			[]interface{}{[]byte("lsm"), "this is a message"},
			`="this is a message"`,
		},
		// 05
		{
			[]interface{}{"", 789.456},
			"=789.456",
		},
		// 06
		{
			[]interface{}{"day", 24, "month", "december", "year", 2019},
			`day=24 month="december" year=2019`,
		},
		// 07
		{
			[]interface{}{"message", "Merry\tChristmas\n"},
			`message="Merry\tChristmas\n"`,
		},
		// 08
		{
			[]interface{}{"odd-numbered list"},
			"[>>odd-numbered list<<]",
		},
	}

	for i, d := range data {
		Encode(&buf, d.kv...)

		if have := buf.String(); have != d.want {
			t.Errorf("test[%02d] => want %v have %v", i, d.want, have)
		}

		buf.Reset()
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
