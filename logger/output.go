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

import "io"

// Output AFAIRE
type Output interface {
	LogDateTime() bool
	LogLevel() bool
	AddNewLine() bool
	Log(level Level, buf []byte) error
	Close() error
}

// OutputProperties AFAIRE
type OutputProperties struct {
	logDateTime bool
	logLevel    bool
	addNewLine  bool
}

// NewOutputProperties AFAIRE
func NewOutputProperties(dateTime, level, newLine bool) *OutputProperties {
	return &OutputProperties{
		logDateTime: dateTime,
		logLevel:    level,
		addNewLine:  newLine,
	}
}

// LogDateTime AFAIRE
func (op *OutputProperties) LogDateTime() bool {
	return op.logDateTime
}

// LogLevel AFAIRE
func (op *OutputProperties) LogLevel() bool {
	return op.logLevel
}

// AddNewLine AFAIRE
func (op *OutputProperties) AddNewLine() bool {
	return op.addNewLine
}

// WriterOutput AFAIRE
type WriterOutput struct {
	io.Writer
	*OutputProperties
}

// NewWriterOutput AFAIRE
func NewWriterOutput(iow io.Writer, op *OutputProperties) *WriterOutput {
	return &WriterOutput{
		Writer:           iow,
		OutputProperties: op,
	}
}

// Log AFAIRE
func (o *WriterOutput) Log(_ Level, buf []byte) error {
	_, err := o.Write(buf)
	return err
}

/*
######################################################################################################## @(°_°)@ #######
*/
