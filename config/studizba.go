package config

import "fmt"

const baseUrl = "https://studizba.com"

const urlPathUniversities = "hs"
const urlPathFiles = "files"

type studIzba struct{}

func (si studIzba) GetUniversitiesUrl() string {
	return makeUrl(urlPathUniversities)
}

func (si studIzba) GetUniversityFilesUrlByUniversityName(name string) string {
	return makeUrl(fmt.Sprintf("%s/%s/%s", urlPathUniversities, name, urlPathFiles))
}

func (si studIzba) GetFileDownloadUrlById(id int) string {
	return makeUrl(fmt.Sprintf("filearray/file-download-%d.html", id))
}

func makeUrl(urlPath string) string {
	return fmt.Sprintf("%s/%s", baseUrl, urlPath)
}
