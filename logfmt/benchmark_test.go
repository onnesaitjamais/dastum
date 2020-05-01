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

func BenchmarkEncode(b *testing.B) {
	buf := bytes.Buffer{}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Encode(&buf, "day", 24, "month", "april", "year", 2020)
		buf.Reset()
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
