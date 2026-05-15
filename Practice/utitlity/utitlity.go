package utitlity

func CheckType(value any) string {
	switch value.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case float64:
		return "float64"
	default:
		return "error"
	}
}
