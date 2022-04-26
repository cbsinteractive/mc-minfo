package minfo

import (
	"context"
	"os/exec"
)

func ReadURL(c context.Context, path string) (media File, err error) {
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
