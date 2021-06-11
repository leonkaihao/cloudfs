package utils

import (
	"io"
	"net/url"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	aferoS3 "github.com/fclairamb/afero-s3"
	"github.com/pkg/errors"
	"github.com/spf13/afero"
)

type FS = afero.Fs

func GetBaseFs(baseurl string) (FS, error) {
	base, err := url.Parse(baseurl)
	if err != nil {
		return nil, err
	}
	switch base.Scheme {
	case "file", "", "local":
		return afero.NewBasePathFs(afero.NewOsFs(), base.Host+base.Path), nil
	case "s3", "s3n", "s3a":
		sess := session.Must(session.NewSession())
		return afero.NewBasePathFs(aferoS3.NewFs(base.Host, sess), base.Path), nil
	default:
		return nil, errors.WithStack(errors.Errorf("unsupported data location scheme: %s",
			base.Scheme))
	}
}

func NewAferoReadCloser(fs FS, path string) (io.ReadCloser, error) {
	f, err := fs.OpenFile(path, os.O_RDONLY, 0755)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func NewAferoWriteCloser(fs FS, path string) (io.WriteCloser, error) {
	f, err := fs.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		return nil, err
	}
	return f, nil
}
