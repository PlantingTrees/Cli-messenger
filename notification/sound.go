package notification

import (
	"embed"
	"log"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
)

//go:embed assets/pop-up-audio.wav
var soundFS embed.FS

var isSpeakerInit = false

// PlaySound plays the embedded sound file
func PlaySound() {
	f, err := soundFS.Open("assets/pop-up-audio.wav")
	if err != nil {
		log.Println("Embedded sound not found:", err)
		return
	}

	// Decode the .wav data
	streamer, format, err := wav.Decode(f)
	if err != nil {
		log.Println("Could not decode WAV:", err)
		return
	}
	defer streamer.Close()

	// Initialize Speaker (Once)
	if !isSpeakerInit {
		err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
		if err != nil {
			log.Println("Speaker init failed:", err)
			return
		}
		isSpeakerInit = true
	}

	// Play
	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
}
