/*
#######
##           __         __
##       ___/ /__ ____ / /___ ____ _
##      / _  / _ `(_-</ __/ // /  ' \
##      \_,_/\_,_/___/\__/\_,_/_/_/_/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package logger

import (
	"testing"
)

var fakeMessage = "Test logging, but use a somewhat realistic message length."

func newLogger(fmt Formatter) *Logger {
	return &Logger{
		level:     InfoLevel,
		runner:    "benchmark",
		formatter: fmt,
		output:    newNopOutput(),
	}
}

func BenchmarkNop(b *testing.B) {
	logger := newLogger(&nopFormatter{})

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		logger.Info(fakeMessage)
	}
}

func BenchmarkDisabled(b *testing.B) {
	logger := newLogger(&nopFormatter{})

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		logger.Debug(fakeMessage) // DebugLevel < InfoLevel
	}
}

func BenchmarkWithoutContext(b *testing.B) {
	logger := newLogger(NewDefaultFormatter())

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		logger.Info(fakeMessage)
	}
}

func BenchmarkWithContext(b *testing.B) {
	logger := newLogger(NewDefaultFormatter())

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		logger.Info(fakeMessage, "day", 24, "month", "april", "year", 2020)
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
