package create

import (
	"fmt"
	"ms-cli/pkg/shared"
	"os"
	"strings"

	cp "github.com/otiai10/copy"
	"github.com/sirupsen/logrus"
)

func CreateMicroservice() {
	msNamePromptContent := shared.PromptContent{
		ErrorMsg: "Please, provide the microservice name",
		Label:    "What's the microservice name?",
	}
	msName := shared.PromptGetInput(msNamePromptContent)

	msTypePromptContent := shared.PromptContent{
		ErrorMsg: "Please provide a microservice type.",
		Label:    "Select the microservice type to build",
	}
	msType := shared.PromptGetSelect(msTypePromptContent, []string{SQSListener, Ambassador, Bulkhead, API})

	languagePromptContent := shared.PromptContent{
		ErrorMsg: "Please provide the language type to use",
		Label:    "Select the language type to build",
	}
	var language string
	switch {
	case msType == SQSListener:
		language = shared.PromptGetSelect(languagePromptContent, []string{Typescript})
	case msType == Ambassador || msType == Bulkhead:
		language = shared.PromptGetSelect(languagePromptContent, []string{Typescript, Go})
	default:
		language = shared.PromptGetSelect(languagePromptContent, []string{Typescript, Go, Python})
	}
	fmt.Printf("Name: %s, type: %s, language: %s", msName, msType, language)
	sourceDir := fmt.Sprintf("pkg/templates/%s/%s", strings.ToLower(language), strings.ToLower(msType))
	target, err := os.Getwd()
	if err != nil {
		logrus.Fatal(err)
	}
	targetDir := fmt.Sprintf("%s/%s", target, msName)
	err = cp.Copy(sourceDir, targetDir)
	if err != nil {
		logrus.Fatal(err)
	}
}
