package notification

// Category correlates to a MAAS notification category.
type Category string

const (
	ERROR   Category = "error"
	WARNING Category = "warning"
	SUCCESS Category = "success"
	INFO    Category = "info"
)
