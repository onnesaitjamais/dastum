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

// Formatter AFAIRE
type Formatter interface {
	Format(level Level, runner, msg string, ctx []interface{}, out Output) []byte
}

/*
######################################################################################################## @(°_°)@ #######
*/
