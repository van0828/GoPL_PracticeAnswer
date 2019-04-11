package ch2_1

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvein float64

const (
	AbsoluteZeroC Celsius    = -273.15
	FreezingC     Celsius    = 0
	BoilingC      Celsius    = 100
	AbsoluteZeroK Kelvein    = 0
	FreezingK     Kelvein    = 273.15
	BoilingK      Kelvein    = 373.15
	AbsoluteZeroF Fahrenheit = -459.66999999999996
	FreezingF     Fahrenheit = 32
	BoilingF      Fahrenheit = 212
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvein) String() string    { return fmt.Sprintf("%g°K", k) }

