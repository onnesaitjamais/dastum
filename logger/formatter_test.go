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

type nopFormatter struct{}

func (e *nopFormatter) Format(_ Level, _, _ string, _ []interface{}, _ Output) []byte {
	return []byte{}
}

/*
######################################################################################################## @(°_°)@ #######
*/
