/*
#######
##           __         __
##       ___/ /__ ____ / /___ ____ _
##      / _  / _ `(_-</ __/ // /  ' \
##      \_,_/\_,_/___/\__/\_,_/_/_/_/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package failure

import (
	"bytes"
	"fmt"

	"github.com/arnumina/dastum/logfmt"
)

// Failure AFAIRE
type Failure struct {
	err error
	ctx []interface{}
	msg string
}

// New AFAIRE
func New(err error) *Failure {
	return &Failure{
		err: err,
	}
}

// Set AFAIRE
func (f *Failure) Set(key string, value interface{}) *Failure {
	f.ctx = append(f.ctx, key, value)
	return f
}

// Setf AFAIRE
func (f *Failure) Setf(key, format string, args ...interface{}) *Failure {
	return f.Set(key, fmt.Sprintf(format, args...))
}

// Msg AFAIRE
func (f *Failure) Msg(msg string) *Failure {
	f.msg = msg
	return f
}

// Msgf AFAIRE
func (f *Failure) Msgf(format string, args ...interface{}) *Failure {
	return f.Msg(fmt.Sprintf(format, args...))
}

// Error AFAIRE
func (f *Failure) Error() string {
	buf := bytes.Buffer{}
	buf.WriteString(f.msg)

	if len(f.ctx) > 0 {
		buf.WriteString(": ")
		logfmt.Encode(&buf, f.ctx...)
	}

	if f.err != nil {
		buf.WriteString(" >>> ")
		buf.WriteString(f.err.Error())
	}

	return buf.String()
}

// Unwrap AFAIRE
func (f *Failure) Unwrap() error {
	return f.err
}

// TODO AFAIRE
func TODO() *Failure {
	return New(nil).Msg("TODO")
}

/*
######################################################################################################## @(°_°)@ #######
*/
