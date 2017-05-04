package model

// A map of prizes to drop rate for the gold chest
// Drop rates are rounded up.
// Numbers represent Au.
var GoldChest = map[string]int{
	"500":      8,
	"550":      7,
	"600":      7,
	"650":      6,
	"700":      6,
	"750":      5,
	"Aphid":    5,
	"Gekko":    5,
	"800":      5,
	"850":      5,
	"900":      4,
	"1000":     4,
	"Orkan":    3,
	"Gepard":   3,
	"Gareth":   3,
	"1500":     3,
	"Ancile":   3,
	"Zeus":     3,
	"2000":     2,
	"2500":     2,
	"Rogatka":  2,
	"Galahad":  2,
	"3000":     1,
	"Jesse":    1,
	"4000":     1,
	"5000":     1,
	"Fury":     1,
	"Lancelot": 1,
	"Doc":      1,
	"Butch":    1,
}

// A map of prizes to drop rate for the gold chest
// Drop rates are rounded up.
// Numbers represent Au.
var SuperChest = map[string]int{
	"Zeus":     13,
	"2000":     10,
	"2250":     9,
	"2500":     8,
	"Rogatka":  8,
	"Galahad":  8,
	"2750":     7,
	"3000":     7,
	"Jesse":    6,
	"3500":     6,
	"4000":     5,
	"Fury":     4,
	"Lancelot": 4,
	"Doc":      3,
	"Butch":    1,
}
