package track

import (
	"io"
	"os"

	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
)

type Track struct {
	FileName string
	stopChan chan bool
}

func MakeTrack(fileName string) *Track {
	return &Track{
		FileName: fileName,
		stopChan: make(chan bool, 1),
	}
}

func (t *Track) Stop() {
	t.stopChan <- true
}

func (t *Track) Play() error {
	file, err := os.Open(t.FileName)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder, err := mp3.NewDecoder(file)
	if err != nil {
		return err
	}
	defer decoder.Close()

	p, err := oto.NewPlayer(decoder.SampleRate(), 2, 2, 8192)
	if err != nil {
		return err
	}
	defer p.Close()

	buf := make([]byte, 8)

	for {
		select {
		case <-t.stopChan:
			return nil
		default:
			r, err := decoder.Read(buf)

			if r > 0 {
				if _, err := p.Write(buf); err != nil {
					return err
				}
			}

			if err == io.EOF {
				decoder.Seek(0, 0)
			} else if err != nil {
				return err
			}
		}
	}
}
