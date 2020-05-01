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
	"fmt"
	"os"
)

// Logger AFAIRE
type Logger struct {
	level     Level
	runner    string
	formatter Formatter
	output    Output
	ctx       []interface{}
}

// New AFAIRE
func New(level, runner string, fmt Formatter, out Output) *Logger {
	return &Logger{
		level:     GetLevelFromString(level),
		runner:    runner,
		formatter: fmt,
		output:    out,
	}
}

// Clone AFAIRE
func (l *Logger) Clone(level string, ctx ...interface{}) *Logger {
	logger := &Logger{
		level:     GetLevelFromString(level),
		runner:    l.runner,
		formatter: l.formatter,
		output:    l.output,
	}

	logger.ctx = append(append(logger.ctx, l.ctx...), ctx...)

	return logger
}

func (l *Logger) log(level Level, msg string, ctx ...interface{}) {
	if level < l.level {
		return
	}

	nc := append(append([]interface{}{}, l.ctx...), ctx...)

	if err := l.output.Log(level, l.formatter.Format(level, l.runner, msg, nc, l.output)); err != nil {
		fmt.Fprintf( ///////////////////////////////////////////////////////////////////////////////////////////////////
			os.Stderr,
			"Error when logging to an output >>> %s\n",
			err,
		)
	}
}

// Trace AFAIRE
func (l *Logger) Trace(msg string, ctx ...interface{}) {
	l.log(TraceLevel, msg, ctx...)
}

// Debug AFAIRE
func (l *Logger) Debug(msg string, ctx ...interface{}) {
	l.log(DebugLevel, msg, ctx...)
}

// Info AFAIRE
func (l *Logger) Info(msg string, ctx ...interface{}) {
	l.log(InfoLevel, msg, ctx...)
}

// Notice AFAIRE
func (l *Logger) Notice(msg string, ctx ...interface{}) {
	l.log(NoticeLevel, msg, ctx...)
}

// Warning AFAIRE
func (l *Logger) Warning(msg string, ctx ...interface{}) {
	l.log(WarningLevel, msg, ctx...)
}

// Error AFAIRE
func (l *Logger) Error(msg string, ctx ...interface{}) {
	l.log(ErrorLevel, msg, ctx...)
}

// Critical AFAIRE
func (l *Logger) Critical(msg string, ctx ...interface{}) {
	l.log(CriticalLevel, msg, ctx...)
}

// Close AFAIRE
func (l *Logger) Close() {
	if err := l.output.Close(); err != nil {
		fmt.Fprintf( ///////////////////////////////////////////////////////////////////////////////////////////////////
			os.Stderr,
			"Error when closing the logger >>> %s\n",
			err,
		)
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
