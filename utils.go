package fias

import (
	"time"
)

const (
	FORMAT_DATE = "060102" // YYMMDD
	FORMAT_TIME = "150405" // HHMMSS

	STX byte = 0x02 // start of text
	ETX byte = 0x03 // end of text
)

func ParseDA(daString string) (time.Time, error) {
	return time.Parse(FORMAT_DATE, daString)
}

func GetDA() string {
	df := DateField{
		Value: time.Now(),
	}
	// log.Print(df.Format())
	return string(FID_Date) + df.Format()
}

func GetTI() string {
	tf := TimeField{
		Value: time.Now(),
	}
	return string(FID_Time) + tf.Format()
}

func FormatMessage(msg string) []byte {
	var result []byte
	result = append(result, STX)
	result = append(result, []byte(msg)...)
	result = append(result, ETX)
	return result
}

func FieldsetToString(fields []FieldID) string {
	var fieldset = ""
	for _, field := range fields {
		fieldset += string(field)
	}
	return fieldset
}
