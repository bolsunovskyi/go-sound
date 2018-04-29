package track

import (
	"errors"
	"os"
)

type Tracks struct {
	items map[string]*Track
}

func MakeTracks() *Tracks {
	return &Tracks{
		items: make(map[string]*Track),
	}
}

func (t *Tracks) AddMultipleTracks(namePath ...string) error {
	if len(namePath)%2 != 0 {
		return errors.New("wrong params number")
	}

	for i := 0; i < len(namePath); i += 2 {
		if err := t.AddTrack(namePath[i], namePath[i+1]); err != nil {
			return err
		}
	}

	return nil
}

func (t *Tracks) AddTrack(name, path string) error {
	if _, err := os.Stat(path); err != nil {
		return err
	}

	t.items[name] = MakeTrack(path)
	return nil
}

func (t *Tracks) Stop(name string) error {
	tr, ok := t.items[name]
	if !ok {
		return errors.New("track does not exists")
	}

	tr.Stop()
	return nil
}

func (t *Tracks) Play(name string) error {
	tr, ok := t.items[name]
	if !ok {
		return errors.New("track does not exists")
	}

	return tr.Play()
}
