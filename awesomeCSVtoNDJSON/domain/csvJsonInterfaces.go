package domain

type Input interface {
	InputFile2Json() []InputFile
}

type Output interface {
	CreateOutput() ([]OutputFile, error)
}

type Numbers interface {
	Request() (string, error)
}
