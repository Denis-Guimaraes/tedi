package output

type Output interface {
	Write(extractedTexts map[string][]string)
}

type stdout struct{}

type csv struct {
	outputPath string
}

func New(outputType string) Output {
	switch outputType {
	case "stdout":
		return &stdout{}
	case "csv":
		return &csv{}
	default:
		return &stdout{}
	}
}
