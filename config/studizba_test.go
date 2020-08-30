package config

import (
	"fmt"
	"github.com/n0tm/studizba-searcher/tests"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"os"
	"testing"
)

var config studIzba

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}

func setUp() {
	config = studIzba{}
}

func TestGetUniversitiesUrl(t *testing.T) {
	assert.Equal(t, createUrl(urlPathUniversities), config.GetUniversitiesUrl())
}

func TestGetUniversityFilesUrlByUniversityName(t *testing.T) {
	universityName := tests.RandomString(10)
	expected := createUrl(fmt.Sprintf("%s/%s/%s", urlPathUniversities, universityName, urlPathFiles))
	assert.Equal(t, expected, config.GetUniversityFilesUrlByUniversityName(universityName))
}

func TestGetFileDownloadUrlById(t *testing.T) {
	id := rand.Int()
	expected := createUrl(fmt.Sprintf("filearray/file-download-%d.html", id))
	assert.Equal(t, expected, config.GetFileDownloadUrlById(id))
}

func createUrl(path string) string {
	return fmt.Sprintf("%s/%s", baseUrl, path)
}
