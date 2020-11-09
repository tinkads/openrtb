package vast

type Mode string

const (
	MODE_ALL  Mode = "all"  // All companion ads must be displayed.
	MODE_ANY  Mode = "any"  // At least one companion ad must be displayed.
	MODE_NONE Mode = "none" // Companion ad display is optional.
)
