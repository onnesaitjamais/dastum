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

import "os"

type stderrOutput struct {
	*WriterOutput
}

// NewStderrOutput AFAIRE
func NewStderrOutput() Output {
	return &stderrOutput{
		WriterOutput: NewWriterOutput(os.Stderr, NewOutputProperties(true, true, true)),
	}
}

func (o *stderrOutput) Close() error {
	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
