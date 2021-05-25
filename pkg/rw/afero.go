package rw

import (
	"io"
	"net/url"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	aferoS3 "github.com/fclairamb/afero-s3"
	"github.com/pkg/errors"
	"github.com/spf13/afero"
)

func NewAferoReadCloser(fs afero.Fs, path string) (io.ReadCloser, error) {
	f, err := fs.OpenFile(path, os.O_RDONLY, 0755)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func NewAferoWriteCloser(fs afero.Fs, path string) (io.WriteCloser, error) {
	f, err := fs.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func GetBaseFs(path string) (afero.Fs, error) {
	loc, err := url.Parse(path)
	if err != nil {
		return nil, err
	}
	switch loc.Scheme {
	case "file", "", "local":
		return afero.NewBasePathFs(afero.NewOsFs(), loc.Host+loc.Path), nil
	case "s3", "s3n", "s3a":
		sess := session.Must(session.NewSession())
		return afero.NewBasePathFs(aferoS3.NewFs(loc.Host, sess), loc.Path), nil
	default:
		return nil, errors.WithStack(errors.Errorf("unsupported data location scheme: %s",
			loc.Scheme))
	}
}
