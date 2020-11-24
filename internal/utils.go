package internal

import (
	"errors"
	"strings"
)

// TrimLeadingWhitespace removes leading whitespace
func TrimLeadingWhitespace(commandStr string) (newCommand string) {
	lines := strings.Split(strings.ReplaceAll(commandStr, "\t", "    "), "\n")
	leadingWhitespace := 1<<31 - 1
	for _, line := range lines {
		for pos, char := range line {
			if char == ' ' || char == '\t' {
				continue
			}
			// first non-whitespace character
			if pos < leadingWhitespace {
				leadingWhitespace = pos
			}
			// is not needed further
			break
		}
	}
	for i := 0; i < len(lines); i++ {
		if lines[i] == "" || strings.Trim(lines[i], " \t") == "" {
			continue
		}
		if len(lines[i]) < leadingWhitespace {
			newCommand += lines[i] + "\n" // or not..
		} else {
			newCommand += lines[i][leadingWhitespace:] + "\n"
		}
	}
	return
}

// ConvertListInterfaceToString ...
func ConvertListInterfaceToString(m []interface{}) []string {
	response := []string{}
	for _, v := range m {
		if v != nil {
			response = append(response, v.(string))
		}
	}
	return response
}

var (
	// ...
	PathEmptyError = errors.New("provided path is empty")
	//...
	DirPathRootDirError = errors.New("dir path is root directory")
)

// GetParentDirPath Os libraries behave bizarely on windows as
// they will replace slashes with other values.
// This causes issues & errors when submitting the request
func GetParentDirPath(filePath string) (string, error) {

	// TODO: this function is not tested well. replace with filepath.Dir
	if filePath == "" {
		return "", PathEmptyError
	}

	pathParts := strings.Split(filePath, "/")

	// if length of pathParts is just two items then the parent should be the root directory
	if len(pathParts) == 2 {
		return "", DirPathRootDirError
	}

	dirPath := strings.Join(pathParts[0:len(pathParts)-1], "/")

	return dirPath, nil
}
