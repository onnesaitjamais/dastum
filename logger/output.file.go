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

type fileOutput struct {
	*WriterOutput
	file *os.File
}

// NewFileOutput AFAIRE
func NewFileOutput(name string) (Output, error) {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return nil, err
	}

	output := &fileOutput{
		WriterOutput: NewWriterOutput(file, NewOutputProperties(true, true, true)),
		file:         file,
	}

	return output, nil
}

// Close AFAIRE
func (fo *fileOutput) Close() error {
	return fo.file.Close()
}

/*
######################################################################################################## @(°_°)@ #######
*/
