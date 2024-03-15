package sfx

import (
	"fmt"
	"log"

	"github.com/veandco/go-sdl2/mix"
)

type Audio struct {
	Tracklist   map[string]string
	SfxList     map[string]string
	TrackChunks map[string]*mix.Chunk
}

func CreateAudio() *Audio {
	if err := mix.OpenAudio(22050, mix.DEFAULT_FORMAT, 2, 4096); err != nil {
		log.Println(err)
	}
	return &Audio{
		Tracklist: map[string]string{
			"test1": "assets/sfx/test1.mp3",
		},
		SfxList: map[string]string{
			"testsfx1": "assets/sfx/test1.mp3",
		},
	}
}

func (audio *Audio) GenerateChunks() {
	chunks := make(map[string]*mix.Chunk)

	for key, val := range audio.Tracklist {
		chunk, err := mix.LoadWAV(val)
		if err != nil {
			fmt.Println(err)
			continue
		}
		chunks[key] = chunk
	}

	audio.TrackChunks = chunks
}

func (audio *Audio) PlayTrack(id string) {

	chunk := audio.TrackChunks[id]

	if chunk == nil {
		return
	}
	chunk.Play(-1, 0)
}

func (audio *Audio) PlayTrackLoop(id string) {

	chunk := audio.TrackChunks[id]

	if chunk == nil {
		return
	}
	chunk.Play(-1, -1)
}
