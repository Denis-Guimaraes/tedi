package output

type Output interface {
	Write(extractedTexts map[string][]string)
}

type stdout struct{}

func New(outputType string) Output {
	switch outputType {
	case "stdout":
		return stdout{}
	default:
		return stdout{}
	}
}
