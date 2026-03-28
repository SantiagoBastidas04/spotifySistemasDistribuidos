package capaControladores

import (
	"fmt"

	fachada "servidor.local/servidorStreaming/capaFachada"
	pb "servidor.local/servidorStreaming/serviciosAudio"
)

// ControladorStreaming implementa la interfaz AudioServiceServer generada por protoc.
type ControladorStreaming struct {
	pb.UnimplementedAudioServiceServer
}

// AudioStream es el procedimiento remoto invocado por el cliente para reproducir un audio.
// Recibe el título del audio y envía los fragmentos MP3 al cliente mediante streaming.
func (s *ControladorStreaming) AudioStream(req *pb.AudioRequest, stream pb.AudioService_AudioStreamServer) error {
	titulo := req.GetFilename()
	fmt.Printf("[RPC] AudioStream llamado con filename=%s\n", titulo)

	if err := fachada.EnviarFragmentosAudio(titulo, stream); err != nil {
		fmt.Printf("[RPC] Error en AudioStream para '%s': %v\n", titulo, err)
		return err
	}

	fmt.Printf("[RPC] AudioStream finalizado para: %s\n", titulo)
	return nil
}