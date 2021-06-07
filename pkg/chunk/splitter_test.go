package chunk

import (
	"context"
	"fmt"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSplitter_Split(t *testing.T) {
	l, err := url.Parse("")
	fmt.Print(l.Scheme, l.Path, err)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	reader := strings.NewReader("123456789012345678901234567890")
	splitter := NewSplitter(8)
	datas, _ := splitter.Split(ctx, reader)
	data := <-datas
	assert.Equal(t, "12345678", string(data))
	data = <-datas
	assert.Equal(t, "90123456", string(data))
	data = <-datas
	assert.Equal(t, "78901234", string(data))
	data = <-datas
	assert.Equal(t, "567890", string(data))
}

func TestSplitter_Split_withcancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	reader := strings.NewReader("123456789012345678901234567890")
	splitter := NewSplitter(8)
	datas, errs := splitter.Split(ctx, reader)
	data := <-datas
	assert.Equal(t, "12345678", string(data))
	data = <-datas
	assert.Equal(t, "90123456", string(data))
	cancel()
	select {
	case err := <-errs:
		assert.Error(t, err)
	case <-time.After(time.Second):
		assert.Fail(t, "no error")
	}
}
