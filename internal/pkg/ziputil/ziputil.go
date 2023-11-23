package ziputil

import (
	"archive/zip"
	"bytes"
	"io"
	"os"
	"strings"
)

// CompressToBuffer ..
func CompressToBuffer(files []string) ([]byte, error) {
	targetFiles := make([]*os.File, 0)
	for _, file := range files {
		f1, err := os.Open(file)
		if err != nil {
			//log.Errorf("打开文件出错，%s， %v", file, err)
			continue
		}
		targetFiles = append(targetFiles, f1)
	}
	defer func() {
		for i := range targetFiles {
			targetFiles[i].Close()
		}
	}()
	// Create a buffer to write our archive to.
	buf := new(bytes.Buffer)
	// Create a new zip archive.
	w := zip.NewWriter(buf)
	for _, file := range targetFiles {
		err := compress(file, "", w)
		if err != nil {
			return nil, err
		}
	}
	w.Close()
	return buf.Bytes(), nil
}

// CompressToFile ..
func CompressToFile(files []string, dest string) error {
	targetFiles := make([]*os.File, 0)
	for _, file := range files {
		f1, err := os.Open(file)
		if err != nil {
			//log.Errorf("打开文件出错，%s， %v", file, err)
			continue
		}
		targetFiles = append(targetFiles, f1)
	}
	defer func() {
		for i := range targetFiles {
			targetFiles[i].Close()
		}
	}()
	d, _ := os.Create(dest)
	defer d.Close()
	w := zip.NewWriter(d)
	defer w.Close()
	for _, file := range targetFiles {
		err := compress(file, ".", w)
		if err != nil {
			return err
		}
	}
	return nil
}

func compress(file *os.File, prefix string, zw *zip.Writer) error {
	info, err := file.Stat()
	if err != nil {
		return err
	}
	if info.IsDir() {
		prefix = prefix + "/" + info.Name()
		fileInfos, err := file.Readdir(-1)
		if err != nil {
			return err
		}
		for _, fi := range fileInfos {
			f, err := os.Open(file.Name() + "/" + fi.Name())
			if err != nil {
				return err
			}
			err = compress(f, prefix, zw)
			if err != nil {
				return err
			}
		}
	} else {
		header, err := zip.FileInfoHeader(info)
		header.Name = prefix + "/" + header.Name
		if err != nil {
			return err
		}
		writer, err := zw.CreateHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(writer, file)
		file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

// DeCompress 解压
func DeCompress(zipFile, dest string) error {
	reader, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer reader.Close()
	for _, file := range reader.File {
		rc, err := file.Open()
		if err != nil {
			return err
		}
		defer rc.Close()
		filename := dest + file.Name
		err = os.MkdirAll(getDir(filename), 0755)
		if err != nil {
			return err
		}
		w, err := os.Create(filename)
		if err != nil {
			return err
		}
		defer w.Close()
		_, err = io.Copy(w, rc)
		if err != nil {
			return err
		}
		w.Close()
		rc.Close()
	}
	return nil
}

func getDir(path string) string {
	return subString(path, 0, strings.LastIndex(path, "/"))
}

func subString(str string, start, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		panic("start is wrong")
	}

	if end < start || end > length {
		panic("end is wrong")
	}

	return string(rs[start:end])
}
