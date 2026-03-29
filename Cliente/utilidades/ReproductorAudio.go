package utilidades

import (
	"fmt"
	"io"
	"log"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"

	pb "servidor.local/servidorStreaming/serviciosAudio"
)

// DecodificarReproducir decodifica el stream MP3 desde el pipe y lo reproduce.
// Cierra canalSincronizacion cuando la reproducción finaliza.
func DecodificarReproducir(reader io.Reader, canalSincronizacion chan struct{}) {
	streamer, format, err := mp3.Decode(io.NopCloser(reader))
	if err != nil {
		log.Fatalf("Error decodificando MP3: %v", err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/2))

	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		close(canalSincronizacion)
	})))
}

// RecibirAudio recibe los fragmentos (chunks) del servidor vía streaming gRPC
// y los escribe en el pipe para que DecodificarReproducir los consuma.
func RecibirAudio(
	stream pb.AudioService_AudioStreamClient,
	writer *io.PipeWriter,
	canalSincronizacion chan struct{},
) {
	noFragmento := 0
	for {
		fragmento, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("\nAudio recibido completo.")
			writer.Close()
			break
		}
		if err != nil {
			log.Printf("Error recibiendo fragmento: %v", err)
			writer.CloseWithError(err)
			break
		}
		noFragmento++
		fmt.Printf("\rFragmento #%d recibido (%d bytes) reproduciendo...", noFragmento, len(fragmento.Data))

		if _, err := writer.Write(fragmento.Data); err != nil {
			log.Printf("Error escribiendo en pipe: %v", err)
			break
		}
	}

	// Esperar a que el speaker termine de reproducir
	<-canalSincronizacion
	fmt.Println("\nReproducción finalizada.")
}