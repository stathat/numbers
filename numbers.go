package numbers

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Scale(num float64, digits int) float64 {
	if num == 0.0 {
		return 0.0
	}
	if math.IsNaN(num) {
		return 0.0
	}
	if math.IsInf(num, 0) {
		return num
	}

	original := num
	num = math.Abs(num)

	d := math.Ceil(math.Log10(num))
	power := digits - int(d)
	magnitude := math.Pow10(power)
	shifted := math.Ceil(num * magnitude)
	result := shifted / magnitude

	result = math.Copysign(result, original)
	return result
}

func SlideScale(num float64) float64 {
	if math.Abs(num) < 100.0 {
		return Scale(num, 1)
	}
	return Scale(num, 2)
}

func ScaleDown(num float64, digits int) float64 {
	if num == 0.0 {
		return 0.0
	}
	if math.IsNaN(num) {
		return 0.0
	}
	if math.IsInf(num, 0) {
		return num
	}

	original := num
	num = math.Abs(num)

	result := 0.0

	if num < 1.0 {
		result = math.Floor(num*10.0) / 10.0
	} else {
		d := math.Ceil(math.Log10(num))
		power := digits - int(d)
		magnitude := math.Pow10(power)
		shifted := math.Floor(num * magnitude)
		result = shifted / magnitude
	}

	result = math.Copysign(result, original)
	return result
}

func SlideScaleDown(num float64) float64 {
	if math.Abs(num) < 100.0 {
		return ScaleDown(num, 1)
	}
	return ScaleDown(num, 2)
}

func CentsToDollars(cents float64) string {
	dollars := cents / 100.0
	return fmt.Sprintf("$%.2f", dollars)
}

func AddDelimiters(num float64) string {
	s := strconv.FormatFloat(num, 'f', 3, 64)
	if num < 1000.0 {
		return s
	}
	pieces := strings.Split(s, ".")
	digits := []string(nil)

	inum := pieces[0]
	x := len(inum)
	for x > 2 {
		digits = append(digits, inum[x-3:x])
		x -= 3
	}
	if x > 0 {
		digits = append(digits, inum[0:x])
	}

	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}

	// drop the fraction
	// result := strings.Join(digits, ",") + "." + pieces[1]
	result := strings.Join(digits, ",")

	return result
}

func AddDelimitersInt(num int) string {
	inum := strconv.Itoa(num)
	digits := []string(nil)

	x := len(inum)
	for x > 2 {
		digits = append(digits, inum[x-3:x])
		x -= 3
	}
	if x > 0 {
		digits = append(digits, inum[0:x])
	}

	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}

	// drop the fraction
	// result := strings.Join(digits, ",") + "." + pieces[1]
	result := strings.Join(digits, ",")

	return result
}

func Display(num float64) string {
	if num >= 10000.0 {
		return Humanize(num)
	}
	if num >= 1000.0 {
		return AddDelimiters(num)
	}

	if num > 1.0 && math.Floor(num) == num {
		return fmt.Sprintf("%.0f", num)
	}

	if num == 0.0 {
		return "0"
	}

	return fmt.Sprintf("%.3f", num)
}

func Humanize(num float64) string {
	var exponent = 0.0
	if num != 0 {
		exponent = math.Floor(math.Log10(math.Abs(num)))
	}

	if exponent >= 3 {
		unit := "K"
		dispExponent := 3
		if exponent >= 15 {
			unit = "Q"
			dispExponent = 15
		} else if exponent >= 12 {
			unit = "T"
			dispExponent = 12
		} else if exponent >= 9 {
			unit = "G"
			dispExponent = 9
		} else if exponent >= 6 {
			unit = "M"
			dispExponent = 6
		}
		num /= math.Pow10(dispExponent)
		return fmt.Sprintf("%.2f%s", num, unit)
	}

	return AddDelimiters(num)
}

func Words(num float64) string {
	var exponent = 0.0
	if num != 0 {
		exponent = math.Floor(math.Log10(math.Abs(num)))
	}

	if exponent >= 3 {
		unit := "thousand"
		dispExponent := 3
		if exponent >= 15 {
			unit = "quadrillion"
			dispExponent = 15
		} else if exponent >= 12 {
			unit = "trillion"
			dispExponent = 12
		} else if exponent >= 9 {
			unit = "billion"
			dispExponent = 9
		} else if exponent >= 6 {
			unit = "million"
			dispExponent = 6
		}
		num /= math.Pow10(dispExponent)
		return fmt.Sprintf("%.2f %s", num, unit)
	}

	return AddDelimiters(num)
}

func Percentage(current, old float64) float64 {
	if old == 0.0 {
		return 0.0
	}
	return 100.0 * (current - old) / old
}

func DisplayPercentage(percentage float64) string {
	return Display(percentage)
}

// using midpoint method
func PercentageMid(current, old float64) float64 {
	mid := 0.5 * (current + old)
	if mid == 0.0 {
		return 0.0
	}
	return 100.0 * (current - old) / mid
}

func Megabytes(bytes uint64) float64 {
	return float64(bytes) / (1024.0 * 1024.0)
}
