package conv

import "fmt"

type Meters float64
type Miles float64

func (m Meters) String() string { return fmt.Sprintf("%g meters", m) }
func (m Miles) String() string  { return fmt.Sprintf("%g miles", m) }

func MilesToMeters(m Miles) Meters { return Meters(m * 1609.34) }
