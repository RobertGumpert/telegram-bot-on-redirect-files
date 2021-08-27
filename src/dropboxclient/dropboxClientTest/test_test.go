package dropboxclienttest

import (
	"application/src/config"
	"application/src/dropboxclient"
	"bufio"
	"log"
	"os"
	"strings"
	"testing"
)

const (
	TestFileName = "test-file.png"
)

func TestUploadFlow(t *testing.T) {
	env, err := config.Reader(config.Test, config.Env)
	if err != nil {
		t.Fatal(err)
	}
	testDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	filePath := strings.Join([]string{testDir, TestFileName}, "/")
	file, err := os.Open(filePath)
	if err != nil {
		t.Fatal(err)
	}
	ioReader := bufio.NewReader(file)
	client := dropboxclient.NewDropboxClient(env.DropboxToken)
	metaData, err := client.Upload(TestFileName, ioReader)
	if err != nil {
		t.Fatal(err)
	}
	log.Println(metaData)
}
