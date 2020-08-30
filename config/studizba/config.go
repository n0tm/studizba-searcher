package studizba

import "fmt"

const baseUrl = "https://studizba.com"

type StudIzba struct{}

func (si StudIzba) GetUniversitiesUrl() string {
	return makeUrl("hs")
}

func (si StudIzba) GetFileDownloadUrlById(id int) string {
	return makeUrl(fmt.Sprintf("filearray/file-download-%d.html", id))
}

func makeUrl(urlPath string) string {
	return fmt.Sprintf("%s/%s", baseUrl, urlPath)
}
