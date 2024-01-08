package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

var (
	groupId string = "ibingli.com"
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
		GroupId: groupId,
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter this application name (no spaces): ")
	an, _ := reader.ReadString('\n')
	fmt.Printf("application's name will be: %s\n", strings.Trim(an, " \n"))
	values.AppName = strings.Trim(an, " \n")

	fmt.Print("Enter this module's name (no spaces): ")
	mn, _ := reader.ReadString('\n')
	fmt.Printf("module's name will be: %s\n", strings.Trim(mn, " \n"))
	values.ModuleName = strings.Trim(mn, " \n")

	wd, _ := os.Getwd()
	fmt.Printf("current working directory: %s\n", wd)

	/*
	 * Create directories
	 */
	moduleDir := fmt.Sprintf("internal/app/%s/%s", values.AppName, values.ModuleName)
	if err := os.MkdirAll(moduleDir, 0755); err != nil {
		log.Panicf("error attempting to create directory '%s', err: %s", moduleDir, err)
	}

	subdirs = []string{
		fmt.Sprintf("%s/%sCtl", moduleDir, values.ModuleName),
		fmt.Sprintf("%s/%sSrv", moduleDir, values.ModuleName),
		fmt.Sprintf("%s/%sDm", moduleDir, values.ModuleName),
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
	if templates, err = template.ParseGlob("tools/moduleTemplates/*.tmpl"); err != nil {
		log.Panicf("error parsing root templates files: %s", err)
	}

	rootFsMapping := map[string]string{
		"README.tmpl":        fmt.Sprintf("%s/README.md", moduleDir),
		"http_ctl_base.tmpl": fmt.Sprintf("%s/%sCtl/http_ctl_base.go", moduleDir, values.ModuleName),
		"http_ctl.tmpl":      fmt.Sprintf("%s/%sCtl/http_ctl.go", moduleDir, values.ModuleName),
		"service_base.tmpl":  fmt.Sprintf("%s/%sSrv/service_base.go", moduleDir, values.ModuleName),
		"service.tmpl":       fmt.Sprintf("%s/%sSrv/service.go", moduleDir, values.ModuleName),
		"repo.tmpl":          fmt.Sprintf("%s/%sSrv/repo.go", moduleDir, values.ModuleName),
		"srv_iface.tmpl":     fmt.Sprintf("%s/%sDm/srv_iface.go", moduleDir, values.ModuleName),
		"req_dto.tmpl":       fmt.Sprintf("%s/%sDm/req_dto.go", moduleDir, values.ModuleName),
		"rsp_dto.tmpl":       fmt.Sprintf("%s/%sDm/rsp_dto.go", moduleDir, values.ModuleName),
		"vo.tmpl":            fmt.Sprintf("%s/%sDm/vo.go", moduleDir, values.ModuleName),
		"entity.tmpl":        fmt.Sprintf("%s/%sDm/entity.go", moduleDir, values.ModuleName),
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

	rp := fmt.Sprintf("s/%sEntity/\\u%sEntity/g", values.ModuleName, values.ModuleName)
	cmd := exec.Command("sed", rp, rootFsMapping["entity.tmpl"])
	out, err := cmd.Output()
	if err != nil {
		log.Panicf("unable to exeucte sed command, err: %s", err)
	}
	f, err := os.Create(rootFsMapping["entity.tmpl"])
	if err != nil {
		log.Panicf("unable to open file: %s, err: %s", rootFsMapping["entity.tmpl"], err)
	}
	f.Write(out)
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	cmd = exec.Command("sed", rp, rootFsMapping["repo.tmpl"])
	out, err = cmd.Output()
	if err != nil {
		log.Panicf("unable to exeucte sed command, err: %s", err)
	}
	f, err = os.Create(rootFsMapping["repo.tmpl"])
	if err != nil {
		log.Panicf("unable to open file: %s, err: %s", rootFsMapping["repo.tmpl"], err)
	}
	f.Write(out)
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n🎉Congratulations! Your new module is ready. locate at: %s\n", moduleDir)
}
