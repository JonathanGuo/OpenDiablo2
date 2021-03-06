// Code generated by "string2enum -samepkg -linecomment -type CompositeType"; DO NOT EDIT.

package d2enum

import "fmt"

// CompositeTypeFromString returns the CompositeType enum corresponding to s.
func CompositeTypeFromString(s string) CompositeType {
	if len(s) == 0 {
		return 0
	}
	for i := range _CompositeType_index[:len(_CompositeType_index)-1] {
		if s == _CompositeType_name[_CompositeType_index[i]:_CompositeType_index[i+1]] {
			return CompositeType(i)
		}
	}
	panic(fmt.Errorf("unable to locate CompositeType enum corresponding to %q", s))
}

func _(s string) {
	// Check for duplicate string values in type "CompositeType".
	switch s {
	// 0
	case "HD":
	// 1
	case "TR":
	// 2
	case "LG":
	// 3
	case "RA":
	// 4
	case "LA":
	// 5
	case "RH":
	// 6
	case "LH":
	// 7
	case "SH":
	// 8
	case "S1":
	// 9
	case "S2":
	// 10
	case "S3":
	// 11
	case "S4":
	// 12
	case "S5":
	// 13
	case "S6":
	// 14
	case "S7":
	// 15
	case "S8":
	// 16
	case "CompositeTypeMax":
	}
}
