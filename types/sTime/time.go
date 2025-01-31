package sTime

import (
	"fmt"
	"strconv"
	"time"

	"github.com/yasseldg/go-simple/configs/sEnv"
)

var DateSeparator = "."

func init() {
	DateSeparator = sEnv.Get("DateSeparator", DateSeparator)
}

func format(t time.Time, prec int, separator string) (string, int64) {
	if len(separator) == 0 {
		separator = DateSeparator
	}
	df := fmt.Sprintf("2006%s01%s02", separator, separator)
	switch prec {
	case 1:
		return t.Format(df), t.Unix()
	case 2:
		return t.Format(fmt.Sprintf("%s 15:04", df)), t.Unix()
	case 4:
		return t.Format(fmt.Sprintf("%s 15:04:05.000", df)), t.UnixMilli()
	case 5:
		return t.Format(fmt.Sprintf("%s 15:04:05.000000", df)), t.UnixMicro()
	case 6:
		return t.Format(fmt.Sprintf("%s%s15:04:05", df, separator)), t.Unix()
	default:
		return t.Format(fmt.Sprintf("%s 15:04:05", df)), t.Unix()
	}
}

func Format(value any, prec int) string {
	return FormatSep(value, prec, "")
}

func FormatSep(value any, prec int, separator string) string {
	t, err := ToTime(value)
	if err != nil {
		return fmt.Sprintf("%v", value)
	}
	s, _ := format(t, prec, separator)
	return s
}

func ForLog(value any, prec int) string {
	t, err := ToTime(value)
	if err != nil {
		return fmt.Sprintf("%v", value)
	}
	s, d := format(t, prec, "")
	return fmt.Sprintf("%d ( %s )", d, s)
}

func ForWeb(value any, prec int) string {
	t, err := ToTime(value)
	if err != nil {
		return fmt.Sprintf("%v", value)
	}
	s, d := format(t, prec, "")
	return fmt.Sprintf("%d <br> %s", d, s)
}

func ToTime(value any) (time.Time, error) {
	switch v := value.(type) {
	case int:
		// Convertir timestamp de tipo int a time.Time
		return fromInt64(int64(v)), nil
	case int64:
		// Convertir timestamp de tipo int64 a time.Time
		return fromInt64(v), nil
	case time.Time:
		// Devolver valor time.Time directamente
		return v, nil
	case string:
		// Convertir cadena a time.Time
		// Comprobar si la cadena representa un timestamp
		if t, err := strconv.ParseInt(v, 10, 64); err == nil {
			return fromInt64(t), nil
		}
		// Comprobar si la cadena representa una fecha y hora con segundos
		if t, err := time.Parse("2006-01-02 15:04:05", v); err == nil {
			return t, nil
		}
		// Comprobar si la cadena representa una fecha y hora sin segundos
		if t, err := time.Parse("2006-01-02 15:04", v); err == nil {
			return t, nil
		}
		// Comprobar si la cadena representa una fecha
		if t, err := time.Parse("2006-01-02", v); err == nil {
			return t, nil
		}
		// Si no es un timestamp ni una fecha válida, devolver error
		return time.Time{}, fmt.Errorf("cadena %v no es un timestamp ni una fecha válida", v)
	default:
		// Si el valor no es int, int64, time.Time o string, devolver error
		return time.Time{}, fmt.Errorf("valor %v no es de tipo int, int64, time.Time o string", v)
	}
}

func fromInt64(value int64) time.Time {
	if value < 1e12 {
		// Si el valor es menor a 1e12, se trata de un timestamp en segundos
		return time.Unix(value, 0)
	} else if value < 1e15 {
		// Si el valor es menor a 1e15, se trata de un timestamp en milisegundos
		return time.Unix(value/1e3, (value%1e3)*1e6)
	} else {
		// Si el valor es mayor o igual a 1e15, se trata de un timestamp en microsegundos
		return time.Unix(value/1e6, (value%1e6)*1e3)
	}
}
