package oterrlog

import (
	"fmt"
)

/*
* The format of StructuredError is based on the OpenTelemetry log data-model
* https://opentelemetry.io/docs/reference/specification/logs/data-model/
 */

 type Severity uint16

 const (
	 Trace  = 1
	 Trace2 = 2
	 Trace3 = 3
	 Trace4 = 4
	 Debug  = 5
	 Debug2 = 6
	 Debug3 = 7
	 Debug4 = 8
	 Info   = 9
	 Info2  = 10
	 Info3  = 11
	 Info4  = 12
	 Warn   = 13
	 Warn2  = 14
	 Warn3  = 15
	 Warn4  = 16
	 Error  = 17
	 Error2 = 18
	 Error3 = 19
	 Error4 = 20
	 Fatal  = 21
	 Fatal2 = 22
	 Fatal3 = 23
	 Fatal4 = 24
 )
 
 var severityToStr = map[Severity]string{
	 Trace:  "TRACE",
	 Trace2: "TRACE2",
	 Trace3: "TRACE3",
	 Trace4: "TRACE4",
	 Debug:  "DEBUG",
	 Debug2: "DEBUG2",
	 Debug3: "DEBUG3",
	 Debug4: "DEBUG4",
	 Info:   "INFO",
	 Info2:  "INFO2",
	 Info3:  "INFO3",
	 Info4:  "INFO4",
	 Warn:   "WARN",
	 Warn2:  "WARN2",
	 Warn3:  "WARN3",
	 Warn4:  "WARN4",
	 Error:  "ERROR",
	 Error2: "ERROR2",
	 Error3: "ERROR3",
	 Error4: "ERROR4",
	 Fatal:  "FATAL",
	 Fatal2: "FATAL2",
	 Fatal3: "FATAL3",
	 Fatal4: "FATAL4",
 }
 
 var strToSeverity = map[string]Severity{
	 `"TRACE"`:  Trace,
	 `"TRACE2"`: Trace2,
	 `"TRACE3"`: Trace3,
	 `"TRACE4"`: Trace4,
	 `"DEBUG"`:  Debug,
	 `"DEBUG2"`: Debug2,
	 `"DEBUG3"`: Debug3,
	 `"DEBUG4"`: Debug4,
	 `"INFO"`:   Info,
	 `"INFO2"`:  Info2,
	 `"INFO3"`:  Info3,
	 `"INFO4"`:  Info4,
	 `"WARN"`:   Warn,
	 `"WARN2"`:  Warn2,
	 `"WARN3"`:  Warn3,
	 `"WARN4"`:  Warn4,
	 `"ERROR"`:  Error,
	 `"ERROR2"`: Error2,
	 `"ERROR3"`: Error3,
	 `"ERROR4"`: Error4,
	 `"FATAL"`:  Fatal,
	 `"FATAL2"`: Fatal2,
	 `"FATAL3"`: Fatal3,
	 `"FATAL4"`: Fatal4,
 }
 
 func SeverityValue(str string) (Severity, error) {
	_, ok := strToSeverity[str]

	if ok {
		return strToSeverity[str], nil
	}
	return 0, fmt.Errorf("invalid Severity name - %s", str)
 }

 func (s Severity) String() string {
	 return severityToStr[s]
 }

 func (s Severity) IsTrace() bool {
	if s >= Trace && s <= Trace4 {
		return true
	}

	return false
 }
 
 func (s Severity) IsDebug() bool {
	if s >= Trace && s < Trace4 {
		return true
	}

	return false
 }
 func (s Severity) IsInfo() bool {
	if s >= Info && s <= Info4 {
		return true
	}

	return false
 }
 func (s Severity) IsWarn() bool {
	if s >= Warn && s <= Warn4 {
		return true
	}

	return false
 }

 func (s Severity) IsError() bool {
	if s >= Error && s <= Error4 {
		return true
	}

	return false
 }
 func (s Severity) IsFatal() bool {
	if s >= Fatal && s <= Fatal4 {
		return true
	}

	return false
 }