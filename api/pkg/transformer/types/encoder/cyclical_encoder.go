package encoder

import (
	"fmt"
	"github.com/gojek/merlin/pkg/transformer/spec"
	"github.com/gojek/merlin/pkg/transformer/types/converter"
	"math"
	"time"
)

const (
	FloatZero   = 0.0000000001
	Q1LastMonth = 3
	Q2LastMonth = 6
	Q3LastMonth = 9
	H1LastMonth = 6
	February    = 2

	MinInSec  = 60
	HourInSec = 3600
	DayInSec  = 86400
	WeekInSec = 604800

	DaysInSec31 = 2678400
	DaysInSec30 = 2592000
	DaysInSec29 = 2505600
	DaysInSec28 = 2419200

	Q1InSec     = 7776000
	Q1LeapInSec = 7862400
	Q2InSec     = 7862400
	Q3InSec     = 7948800
	Q4InSec     = 7948800

	H1InSec     = 15638400
	H1LeapInSec = 15724800
	H2InSec     = 15897600

	YearInSec     = 31536000
	LeapYearInSec = 31622400

	completeAngle = 2 * math.Pi

	// Unit angles for variable periods
	UnitDaysInSec31 = completeAngle / DaysInSec31
	UnitDaysInSec30 = completeAngle / DaysInSec30
	UnitDaysInSec29 = completeAngle / DaysInSec29
	UnitDaysInSec28 = completeAngle / DaysInSec28

	UnitQ1InSec     = completeAngle / Q1InSec
	UnitQ1LeapInSec = completeAngle / Q1LeapInSec
	UnitQ2InSec     = completeAngle / Q2InSec
	UnitQ3InSec     = completeAngle / Q3InSec
	UnitQ4InSec     = completeAngle / Q4InSec

	UnitH1InSec     = completeAngle / H1InSec
	UnitH1LeapInSec = completeAngle / H1LeapInSec
	UnitH2InSec     = completeAngle / H2InSec

	UnitYearInSec     = completeAngle / YearInSec
	UnitLeapYearInSec = completeAngle / LeapYearInSec
)

// Unit angles for variable periods for each month
var MonthInSec = map[int]float64{
	1:  UnitDaysInSec31,
	2:  UnitDaysInSec28,
	3:  UnitDaysInSec31,
	4:  UnitDaysInSec30,
	5:  UnitDaysInSec31,
	6:  UnitDaysInSec30,
	7:  UnitDaysInSec31,
	8:  UnitDaysInSec31,
	9:  UnitDaysInSec30,
	10: UnitDaysInSec31,
	11: UnitDaysInSec30,
	12: UnitDaysInSec31,
	13: UnitDaysInSec29, // Leap year Feb
}

type CyclicalEncoder struct {
	Period spec.PeriodType
	Min    float64
	Max    float64
}

func NewCyclicalEncoder(config *spec.CyclicalEncoderConfig) (*CyclicalEncoder, error) {
	// by range
	byRange := config.GetByRange()
	if byRange != nil {
		if (byRange.Max - byRange.Min) < FloatZero {
			return nil, fmt.Errorf("max of cyclical range must be larger than min")
		}

		return &CyclicalEncoder{
			Period: spec.PeriodType_UNDEFINED,
			Min:    byRange.Min,
			Max:    byRange.Max,
		}, nil
	}

	// by epoch time
	byEpochTime := config.GetByEpochTime()
	var min, max float64 = 0, 0
	var period spec.PeriodType

	if byEpochTime != nil {
		switch byEpochTime.Period {
		case spec.PeriodType_HOUR:
			period = spec.PeriodType_UNDEFINED
			max = HourInSec
		case spec.PeriodType_DAY:
			period = spec.PeriodType_UNDEFINED
			max = DayInSec
		case spec.PeriodType_WEEK:
			period = spec.PeriodType_UNDEFINED
			max = WeekInSec
		case spec.PeriodType_MONTH, spec.PeriodType_QUARTER, spec.PeriodType_HALF, spec.PeriodType_YEAR:
			period = byEpochTime.Period
			max = 0
		default:
			return nil, fmt.Errorf("invalid or unspported cycle period")
		}

		return &CyclicalEncoder{
			Period: period,
			Min:    min,
			Max:    max,
		}, nil
	}

	return nil, fmt.Errorf("cyclical encoding config invalid or undefined")
}

func (oe *CyclicalEncoder) Encode(values []interface{}, column string) (map[string]interface{}, error) {
	encodedCos := make([]interface{}, 0, len(values))
	encodedSin := make([]interface{}, 0, len(values))

	// config with fixed range
	if oe.Period == spec.PeriodType_UNDEFINED {
		period := oe.Max - oe.Min
		unitAngle := completeAngle / period

		for _, val := range values {
			// Check if value is missing
			if val == nil {
				return nil, fmt.Errorf("missing value")
			}

			// Check if value is valid
			valFloat, err := converter.ToFloat64(val)
			if err != nil {
				return nil, err
			}

			// Encode to sin and cos
			phase := (valFloat - oe.Min) * unitAngle
			encodedCos = append(encodedCos, math.Cos(phase))
			encodedSin = append(encodedSin, math.Sin(phase))
		}
	} else {
		// config with variable range, by epoch time (e.g. different days in each month, leap year etc.)
		for _, val := range values {
			// Check if value is missing
			if val == nil {
				return nil, fmt.Errorf("missing value")
			}

			// Check if value is valid
			valInt, err := converter.ToInt64(val)
			if err != nil {
				return nil, err
			}

			// convert epoch time to golang datetime
			t := time.Unix(valInt, 0).In(time.UTC)
			shareOfPeriod, err := getCycleTime(oe.Period, t)
			if err != nil {
				return nil, err
			}
			unitAngle, err := getUnitAngle(oe.Period, t)
			if err != nil {
				return nil, err
			}

			// Encode to sin and cos
			phase := float64(shareOfPeriod) * unitAngle
			encodedCos = append(encodedCos, math.Cos(phase))
			encodedSin = append(encodedSin, math.Sin(phase))
		}
	}

	return map[string]interface{}{
		column + "_x": encodedCos,
		column + "_y": encodedSin,
	}, nil
}

// Computes the number of seconds past the beginning of a pre-defined cycle
// Only works with PeriodType with variable cycle time such as Month, Year etc
// For period type with fixed cycle time, it is handled differently by encoder and
// does not need the cycle time to be computed
func getCycleTime(periodType spec.PeriodType, t time.Time) (int, error) {
	switch periodType {
	case spec.PeriodType_MONTH:
		dayElapsed := t.Day() - 1
		hr, min, sec := t.Clock()
		elapsed := getElapsedSec(dayElapsed, hr, min, sec)

		return elapsed, nil

	case spec.PeriodType_QUARTER:
		dayElapsed := t.YearDay() - 1
		hr, min, sec := t.Clock()
		elapsed := getElapsedSec(dayElapsed, hr, min, sec)
		var cycleTime int

		if t.Month() <= Q1LastMonth {
			return elapsed, nil
		} else if t.Month() <= Q2LastMonth {
			cycleTime = elapsed - Q1InSec
		} else if t.Month() <= Q3LastMonth {
			cycleTime = elapsed - H1InSec
		} else {
			cycleTime = elapsed - H1InSec - Q3InSec
		}

		if isLeapYear(t.Year()) {
			cycleTime -= DayInSec //minus extra day from leap year
		}

		return cycleTime, nil

	case spec.PeriodType_HALF:
		dayElapsed := t.YearDay() - 1
		hr, min, sec := t.Clock()
		elapsed := getElapsedSec(dayElapsed, hr, min, sec)

		if t.Month() <= 6 {
			return elapsed, nil
		}

		if isLeapYear(t.Year()) {
			return elapsed - H1LeapInSec, nil
		}
		return elapsed - H1InSec, nil

	case spec.PeriodType_YEAR:
		dayElapsed := t.YearDay() - 1
		hr, min, sec := t.Clock()
		elapsed := getElapsedSec(dayElapsed, hr, min, sec)
		return elapsed, nil
	}

	return 0, fmt.Errorf("period type is undefined for this use case")
}

// Convert time duration in days, hour, min, sec to number of seconds
func getElapsedSec(dayElapsed int, hr int, min int, sec int) int {
	elapsed := dayElapsed*DayInSec + hr*HourInSec + min*MinInSec + sec

	return elapsed
}

// Computes the angle in radians represented by per unit second of a pre-defined period
// This is derived from the formula for calculating phase:
// phase = time passed / period * 2pi
// By rearranging the formula (for optimizing computation) into this:
// phase = time pass * 2pi / period
// we define unit angle as (2pi / period)
// The motivation is that we can pre-compute this value once and use it repeatedly.
func getUnitAngle(periodType spec.PeriodType, t time.Time) (float64, error) {
	switch periodType {
	case spec.PeriodType_MONTH:
		if t.Month() == February && isLeapYear(t.Year()) {
			return MonthInSec[13], nil
		}
		return MonthInSec[int(t.Month())], nil
	case spec.PeriodType_QUARTER:
		if t.Month() <= Q1LastMonth {
			if isLeapYear(t.Year()) {
				return UnitQ1LeapInSec, nil
			}
			return UnitQ1InSec, nil
		} else if t.Month() <= Q2LastMonth {
			return UnitQ2InSec, nil
		} else if t.Month() <= Q3LastMonth {
			return UnitQ3InSec, nil
		}
		return UnitQ4InSec, nil
	case spec.PeriodType_HALF:
		if t.Month() <= H1LastMonth {
			if isLeapYear(t.Year()) {
				return UnitH1LeapInSec, nil
			}
			return UnitH1InSec, nil
		}
		return UnitH2InSec, nil
	case spec.PeriodType_YEAR:
		if isLeapYear(t.Year()) {
			return UnitLeapYearInSec, nil
		}
		return UnitYearInSec, nil
	}

	return 0, fmt.Errorf("period type is undefined for this use case")
}

// test if a given year is leap year
// leap year is a year divisible by (4, but not 100) or (4, 100 and 400)
func isLeapYear(year int) bool {
	if (year%4 == 0 && year%100 != 0) || (year%4 == 0 && year%400 == 0) {
		return true
	}
	return false
}
