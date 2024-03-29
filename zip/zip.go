package zip

import (
	"archive/zip"
	"errors"
	"fmt"
	"github.com/Hurobaki/gochunks/utils"
	"io"
	"os"
)

func ZipFiles(zipName string, files []string, path string) error {
	zipFile, err := os.Create(zipName)

	if err != nil {
		return err
	}

	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	for _, file := range files {
		if err := AddFileToZip(zipWriter, utils.FullPath(path, file), file); err != nil {
			return errors.New(fmt.Sprintf("Failed to add file %s to zip: %s", file, err))
		}
	}

	return nil
}

func AddFileToZip(zipWriter *zip.Writer, filePath string, fileName string) error {
	fileToZip, err := os.Open(filePath)

	if err != nil {
		return err
	}
	defer fileToZip.Close()

	info, err := fileToZip.Stat()
	if err != nil {
		return err
	}

	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return err
	}

	header.Name = fileName // To get the fileName instead of its path

	writer, err := zipWriter.CreateHeader(header)

	if err != nil {
		return err
	}
	_, err = io.Copy(writer, fileToZip)

	return err
}
