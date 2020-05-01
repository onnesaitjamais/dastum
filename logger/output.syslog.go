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

import "log/syslog"

// GetFacilityFromString AFAIRE
func GetFacilityFromString(facility string) syslog.Priority {
	switch facility {
	case "local0":
		return syslog.LOG_LOCAL0
	case "local1":
		return syslog.LOG_LOCAL1
	case "local2":
		return syslog.LOG_LOCAL2
	case "local3":
		return syslog.LOG_LOCAL3
	case "local4":
		return syslog.LOG_LOCAL4
	case "local5":
		return syslog.LOG_LOCAL5
	case "local6":
		return syslog.LOG_LOCAL6
	default:
		return syslog.LOG_LOCAL7
	}
}

type syslogOutput struct {
	*syslog.Writer
	*OutputProperties
}

// NewSyslogOutput AFAIRE
func NewSyslogOutput(facility string) (Output, error) {
	writer, err := syslog.New(GetFacilityFromString(facility), "dastum")
	if err != nil {
		return nil, err
	}

	output := &syslogOutput{
		Writer:           writer,
		OutputProperties: NewOutputProperties(false, false, false),
	}

	return output, nil
}

// Log AFAIRE
func (so *syslogOutput) Log(level Level, buf []byte) error {
	switch level {
	case InfoLevel:
		return so.Info(string(buf))
	case NoticeLevel:
		return so.Notice(string(buf))
	case WarningLevel:
		return so.Warning(string(buf))
	case ErrorLevel:
		return so.Err(string(buf))
	case CriticalLevel:
		return so.Crit(string(buf))
	default:
		return so.Debug(string(buf)) // TRACE & DEBUG
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
