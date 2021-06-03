package main

import (
	"fmt"
	"math"
	"strconv"
)

const (
	RoundHalfUp = iota + 1
	RoundHalfDown
	RoundHalfEven
	RoundHalfOdd
)

func round(value float64, Precision int, mode int) float64 {
	var f1, f2, tmpValue float64
	var err error

	maxInt := int(^uint(0) >> 1)
	minInt := -maxInt - 1

	if value == 0 {
		return value
	}

	if Precision < minInt+1 {
		Precision = minInt + 1
	}

	precisionPlaces := 14 - log10Abc(value)

	f1 = pow10(int(math.Abs(float64(Precision))))

	/* If the decimal precision guaranteed by FP arithmetic is higher than
	   the requested Precision BUT is small enough to make sure a non-zero value
	   is returned, pre-round the result to the precision */
	var usePrecision int64
	if precisionPlaces > Precision && precisionPlaces-15 < Precision {
		if precisionPlaces < minInt+1 {
			usePrecision = int64(minInt + 1)
		} else {
			usePrecision = int64(precisionPlaces)
		}

		f2 = pow10(int(math.Abs(float64(usePrecision))))
		if usePrecision >= 0 {
			tmpValue = value * f2
		} else {
			tmpValue = value / f2
		}

		/* preround the result (tmpValue will always be something * 1e14,
		   thus never larger than 1e15 here) */
		tmpValue = roundHelper(tmpValue, mode)

		usePrecision = int64(Precision - precisionPlaces)

		if usePrecision < int64(minInt+1) {
			usePrecision = int64(minInt + 1)
		}

		/* now correctly move the decimal point */
		f2 = pow10(int(math.Abs(float64(usePrecision))))
		/* because Precision < precision_places */
		tmpValue = tmpValue / f2
	} else {
		/* adjust the value */
		if Precision >= 0 {
			tmpValue = value * f1
		} else {
			tmpValue = value / f1
		}

		/* This value is beyond our precision, so rounding it is pointless */
		if math.Abs(tmpValue) >= 1e15 {
			return value
		}
	}

	/* round the temp value */
	tmpValue = roundHelper(tmpValue, mode)

	if math.Abs(float64(Precision)) < 23 {
		if Precision > 0.0 {
			tmpValue = tmpValue / f1
		} else {
			tmpValue = tmpValue * f1
		}
	} else {

		sf := fmt.Sprintf("%15fe%d", tmpValue, -Precision)

		tmpValue, err = strconv.ParseFloat(sf, 64)
		if err != nil {
			tmpValue = value
		}
	}

	return tmpValue
}

func log10Abc(value float64) int {
	var result int
	value = math.Abs(value)

	if value < 1e-8 || value > 1e22 {
		result = int(math.Floor(math.Log(value)))
	} else {
		values := []float64{1e-8, 1e-7, 1e-6, 1e-5, 1e-4, 1e-3, 1e-2, 1e-1,
			1e0, 1e1, 1e2, 1e3, 1e4, 1e5, 1e6, 1e7,
			1e8, 1e9, 1e10, 1e11, 1e12, 1e13, 1e14, 1e15,
			1e16, 1e17, 1e18, 1e19, 1e20, 1e21, 1e22}

		result = 15

		if value < values[result] {
			result -= 8
		} else {
			result += 8
		}

		if value < values[result] {
			result -= 4
		} else {
			result += 4
		}

		if value < values[result] {
			result -= 2
		} else {
			result += 2
		}

		if value < values[result] {
			result -= 1
		} else {
			result += 1
		}

		if value < values[result] {
			result -= 1
		}

		result -= 8
	}

	return result
}

func pow10(power int) float64 {
	powers := []float64{
		1e0, 1e1, 1e2, 1e3, 1e4, 1e5, 1e6, 1e7,
		1e8, 1e9, 1e10, 1e11, 1e12, 1e13, 1e14, 1e15,
		1e16, 1e17, 1e18, 1e19, 1e20, 1e21, 1e22,
	}

	if power < 0 || power > 22 {
		return math.Pow10(power)
	}

	return powers[power]
}

func roundHelper(value float64, mode int) float64 {
	var tmpValue float64

	if value >= 0.0 {
		tmpValue = math.Floor(value + 0.5)

		if mode == RoundHalfDown && value == (-0.5+tmpValue) ||
			mode == RoundHalfEven && value == (0.5+2*math.Floor(tmpValue/2)) ||
			mode == RoundHalfOdd && value == (0.5+2*math.Floor(tmpValue/2)-1) {

			tmpValue = tmpValue - 1.0
		}
	} else {
		tmpValue = math.Ceil(value - 0.5)
		if mode == RoundHalfDown && value == (0.5+tmpValue) ||
			mode == RoundHalfEven && value == (-0.5+2*math.Ceil(tmpValue/2.0)) ||
			mode == RoundHalfOdd && value == (-0.5+2*math.Ceil(tmpValue/2.0)+1.0) {

			tmpValue = tmpValue + 1.0
		}

	}
	return tmpValue
}
