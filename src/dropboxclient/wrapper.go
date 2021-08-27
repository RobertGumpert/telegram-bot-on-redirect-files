package dropboxclient

import (
	"io"
	"strings"

	"github.com/tj/go-dropbox"
)

type DropboxClientWrapper struct {
	token  string
	client *dropbox.Client
}

func NewDropboxClient(token string) *DropboxClientWrapper {
	this := &DropboxClientWrapper{
		token: token,
	}
	client := dropbox.New(dropbox.NewConfig(token))
	this.client = client
	return this
}

func (this *DropboxClientWrapper) Upload(filename string, reader io.Reader) (dropbox.Metadata, error) {
	if filename[0] != '/' {
		filename = strings.Join(
			[]string{
				"/",
				filename,
			},
			"",
		)
	}
	out, err := this.client.Files.Upload(
		&dropbox.UploadInput{
			Path:   filename,
			Reader: reader,
			Mute:   true,
			Mode:   dropbox.WriteModeAdd,
		},
	)
	if err != nil {
		return dropbox.Metadata{}, err
	}
	return out.Metadata, nil
}
