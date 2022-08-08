package minfo

import (
	"context"
	"fmt"
	"net/http"
	"os/exec"
)

func check(path string) error {
	r, err := http.Head(path)
	if err != nil {
		return err
	}
	stat := r.StatusCode
	if stat >= 400 {
		return fmt.Errorf("check: %q: bad http status %d", path, stat)
	}
	return nil
}

func ReadURL(c context.Context, path string) (media File, err error) {
	if err = check(path); err != nil {
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
