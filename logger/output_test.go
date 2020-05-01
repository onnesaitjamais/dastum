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

type nopOutput struct {
	*OutputProperties
}

func newNopOutput() *nopOutput {
	return &nopOutput{
		OutputProperties: NewOutputProperties(false, false, false),
	}
}

func (o *nopOutput) Log(_ Level, _ []byte) error {
	return nil
}

func (o *nopOutput) Close() error {
	return nil
}

/*
type bufferOutput struct {
	*OutputProperties
	bytes.Buffer
}

func newBufferOutput() *bufferOutput {
	return &bufferOutput{
		OutputProperties: NewOutputProperties(false, true, false),
	}
}

func (o *bufferOutput) Log(_ Level, buf []byte) error {
	_, err := o.Write(buf)
	return err
}

func (o *bufferOutput) Close() error {
	return nil
}
*/

/*
######################################################################################################## @(°_°)@ #######
*/
