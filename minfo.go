package minfo

import (
	"context"
	"fmt"
	"net/http"
	"os/exec"
	"strings"
)

func check(c context.Context, path string) error {
	if !strings.HasPrefix(path, "http") {
		return nil
	}
	r, err := http.NewRequest("GET", path, nil)
	if err != nil {
		return err
	}
	r.Header.Set("Range", "bytes=0-100")
	resp, err := http.DefaultClient.Do(r.WithContext(c))
	if err != nil {
		return err
	}
	stat := resp.StatusCode
	if stat >= 400 {
		return fmt.Errorf("check: %q: bad http status %d", path, stat)
	}
	return nil
}

func ReadURL(c context.Context, path string) (media File, err error) {
	if err = check(c, path); err != nil {
		return
	}
	out, err := exec.CommandContext(c, "mediainfo", "--Output=JSON", path).Output()
	if err != nil {
		return
	}
	return media, media.Decode(out)
}

/*
// Cant seem to get this working
func Read(r io.Reader) (media File, err error) {
	cmd := exec.Command("mediainfo", "--Output=JSON", os.Stdin.Name())
	cmd.Stdin = bufio.NewReader(r)
	out, err := cmd.Output()
	if err != nil {
		return
	}
	return media, media.Decode(out)
}
*/
