package domain

type ServiceInput struct {
	fileData [][]string
}

func NewServiceInput(file [][]string) *ServiceInput {
	return &ServiceInput{
		fileData: file,
	}
}

func (s *ServiceInput) InputFile2Json() []InputFile {
	var file []InputFile
	var eachLine InputFile

	for _, line := range s.fileData {
		eachLine = *EnvVarsInput(line[0], line[1], line[2], line[3])
		file = append(file, eachLine)
	}

	return file
}
