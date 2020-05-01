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

type stdoutOutput struct {
	*WriterOutput
}

// NewStdoutOutput AFAIRE
func NewStdoutOutput() Output {
	return &stdoutOutput{
		WriterOutput: NewWriterOutput(os.Stdout, NewOutputProperties(true, true, true)),
	}
}

func (o *stdoutOutput) Close() error {
	return nil
}

/*
######################################################################################################## @(°_°)@ #######
*/
