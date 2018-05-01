package track

import (
	"os/exec"
)

type MPG123Track struct {
	FileName string
	cmd      *exec.Cmd
}

func (t *MPG123Track) Play() error {
	t.cmd = exec.Command("mpg123", t.FileName)
	if err := t.cmd.Start(); err != nil {
		return err
	}

	return nil
}

func (t *MPG123Track) Stop() error {
	return t.cmd.Process.Kill()
}

func MakeMPG123Track(fileName string) *MPG123Track {
	return &MPG123Track{
		FileName: fileName,
	}
}
