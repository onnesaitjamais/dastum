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

// Level AFAIRE
type Level int

// AFAIRE
const (
	TraceLevel Level = iota
	DebugLevel
	InfoLevel
	NoticeLevel
	WarningLevel
	ErrorLevel
	CriticalLevel
)

// GetLevelFromString AFAIRE
func GetLevelFromString(level string) Level {
	switch level {
	case "debug":
		return DebugLevel
	case "info":
		return InfoLevel
	case "notice":
		return NoticeLevel
	case "warning":
		return WarningLevel
	case "error":
		return ErrorLevel
	case "critical":
		return CriticalLevel
	default:
		return TraceLevel
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
