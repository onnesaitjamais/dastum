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
	"bytes"
	"sync"
	"time"

	"github.com/arnumina/dastum/logfmt"
)

// DefaultFormatter AFAIRE
type DefaultFormatter struct {
	buf   bytes.Buffer
	mutex sync.Mutex
}

// NewDefaultFormatter AFAIRE
func NewDefaultFormatter() *DefaultFormatter {
	return &DefaultFormatter{}
}

// Format AFAIRE
func (df *DefaultFormatter) Format(level Level, runner, msg string, ctx []interface{}, output Output) []byte {
	df.mutex.Lock()
	defer df.mutex.Unlock()

	df.buf.Reset()

	if output.LogDateTime() {
		df.buf.WriteString(time.Now().Format("2006-01-02T15:04:05.000"))
	}

	if output.LogLevel() {
		switch level {
		case TraceLevel:
			df.buf.WriteString("{TRA} ")
		case DebugLevel:
			df.buf.WriteString("{DEB} ")
		case InfoLevel:
			df.buf.WriteString("{INF} ")
		case NoticeLevel:
			df.buf.WriteString("{NOT} ")
		case WarningLevel:
			df.buf.WriteString("{WAR} ")
		case ErrorLevel:
			df.buf.WriteString("{ERR} ")
		case CriticalLevel:
			df.buf.WriteString("{CRI} ")
		default:
			df.buf.WriteString("{???} ")
		}
	}

	df.buf.WriteString(runner)
	df.buf.WriteRune(32)
	df.buf.WriteString(msg)

	if len(ctx) != 0 {
		df.buf.WriteString("> ")
		logfmt.Encode(&df.buf, ctx...)
	}

	if output.AddNewLine() {
		df.buf.WriteString("\n")
	}

	return df.buf.Bytes()
}

/*
######################################################################################################## @(°_°)@ #######
*/
