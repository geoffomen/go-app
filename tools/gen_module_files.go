package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

var (
	groupId    string = "github.com/geoffomen"
	artifactId string = "go-app"
	appName    string = "myapp"
)

type moduleValues struct {
	GroupId    string
	ArtifactId string
	AppName    string
	ModuleName string
}

// at project root dir execute: go run tools/gen_module_files.go
func main() {
	var (
		fp        *os.File
		templates *template.Template
		subdirs   []string
	)

	values := moduleValues{
		GroupId:    groupId,
		ArtifactId: artifactId,
		AppName:    appName,
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter this module's name (no spaces): ")
	mn, _ := reader.ReadString('\n')
	fmt.Printf("module's name will be: %s", strings.Trim(mn, " \n"))
	values.ModuleName = strings.Trim(mn, " \n")

	// values.ModuleName = "test"

	wd, _ := os.Getwd()
	fmt.Printf("current working directory\n: %s", wd)

	/*
	 * Create directories
	 */
	moduleDir := fmt.Sprintf("internal/app/%s/%s", values.AppName, values.ModuleName)
	if err := os.MkdirAll(moduleDir, 0755); err != nil {
		log.Panicf("error attempting to create application directory '%s', err: %s", values.AppName, err)
	}

	subdirs = []string{
		fmt.Sprintf("%s/%sctl", moduleDir, values.ModuleName),
		fmt.Sprintf("%s/%simp", moduleDir, values.ModuleName),
	}

	for _, dirname := range subdirs {
		if err := os.MkdirAll(dirname, 0755); err != nil {
			log.Panicf("unable to create subdirectory %s, err: %s", dirname, err)
		}
	}

	/*
	 * Process templates
	 */
	var err error
	if templates, err = template.ParseGlob("tools/moduletemplates/*.tmpl"); err != nil {
		// if templates, err = template.ParseGlob("moduletemplates/*.tmpl"); err != nil {
		log.Panicf("error parsing root templates files: %s", err)
	}

	rootFsMapping := map[string]string{
		"controller.tmpl": fmt.Sprintf("%s/%sctl/controller.go", moduleDir, values.ModuleName),
		"service.tmpl":    fmt.Sprintf("%s/%simp/service.go", moduleDir, values.ModuleName),
		"iface.tmpl":      fmt.Sprintf("%s/iface.go", moduleDir),
		"req_dto.tmpl":    fmt.Sprintf("%s/req_dto.go", moduleDir),
		"rsp_dto.tmpl":    fmt.Sprintf("%s/rsp_dto.go", moduleDir),
	}

	for templateName, outputPath := range rootFsMapping {
		if fp, err = os.Create(outputPath); err != nil {
			log.Panicf("unable to create file %s for writing, err: %s", outputPath, err)
		}

		defer fp.Close()

		if err = templates.ExecuteTemplate(fp, templateName, values); err != nil {
			log.Panicf("unable to exeucte template: %s, err: %s", templateName, err)
		}
	}

	fmt.Printf("\n🎉Congratulations! Your new module is ready. locate at: %s\n", moduleDir)
}
