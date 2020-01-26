package util

import (
	"bufio"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

func GetFileName(ext string) string {
	filePath := getFileStrl(ext)
	return filePath
}

func getFileStrl(ext string) string {
	filename := getFilel(ext)
	_, dirname, _, _ := runtime.Caller(0)
	filePath := getDirectoryName(dirname, filename)
	return filePath
}

func getFilel(nameFile string) []string {
	filename := []string{nameFile}
	return filename
}

func GetFileNameWithExtension(folder string, prefix string, ext string) string {
	env := os.Getenv("ENV")
	if len(env) == 0 {
		env = "development"
	}
	filePath := getFileStr(folder, prefix, env, ext)
	return filePath
}

func getFileStr(folder string, prefix string, env string, ext string) string {
	filename := getFile(folder, prefix, env, ext)
	_, dirname, _, _ := runtime.Caller(0)
	filePath := getDirectoryName(dirname, filename)
	return filePath
}

func getFile(folder string, prefix string, env string, ext string) []string {
	filename := []string{folder, prefix, env, ext}
	return filename
}

func getDirectoryName(dirname string, filename []string) string {
	filePath := path.Join(filepath.Dir(dirname), strings.Join(filename, ""))
	return filePath
}

func GetLocationMap(fromFilename string) map[string]string {
	locations := scanLocationLines(GetFileName(fromFilename))
	return locations
}

func scanLocationLines(path string) map[string]string {
	file, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var locationMap = make(map[string]string)
	//var lines []string
	for scanner.Scan() {
		locationSplit := strings.Split(scanner.Text(), ",")
		//log.Println(locationSplit[0], locationSplit[1])
		locationMap[locationSplit[0]] = locationSplit[1]
		//lines = append(lines, scanner.Text())
	}
	return locationMap
}
