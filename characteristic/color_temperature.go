// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeColorTemperature = "CE"

type ColorTemperature struct {
	*Int
}

func NewColorTemperature() *ColorTemperature {
	char := NewInt(TypeColorTemperature)
	char.Format = FormatUInt32
	char.Perms = []string{PermRead, PermWrite, PermEvents}

	char.SetValue(0)

	return &ColorTemperature{char}
}
