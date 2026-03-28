package capaFachada

import (
	"fmt"
	"io"
	"log"

	repo "servidor.local/servidorStreaming/capaAccesoDatos"
	pb "servidor.local/servidorStreaming/serviciosAudio"
)

// EnviarFragmentosAudio lee el archivo MP3 en fragmentos de 32 KB y los envía
// al cliente a través del stream gRPC.
func EnviarFragmentosAudio(titulo string, stream pb.AudioService_AudioStreamServer) error {
	log.Printf("[Fachada] Preparando envío de fragmentos para: %s", titulo)

	archivo, err := repo.AbrirArchivo(titulo)
	if err != nil {
		return fmt.Errorf("no se pudo abrir el archivo: %w", err)
	}
	defer archivo.Close()

	buf := make([]byte, 32*1024) // 32 KB por fragmento
	noFragmento := 0

	for {
		n, err := archivo.Read(buf)
		if err == io.EOF {
			log.Printf("[Fachada] Envío completo: %s (%d fragmentos)", titulo, noFragmento)
			break
		}
		if err != nil {
			return fmt.Errorf("error leyendo archivo: %w", err)
		}

		if n > 0 {
			noFragmento++
			chunk := &pb.AudioChunk{Data: buf[:n]}
			if err := stream.Send(chunk); err != nil {
				return fmt.Errorf("error enviando fragmento #%d: %w", noFragmento, err)
			}
			log.Printf("[Fachada] Fragmento #%d enviado (%d bytes)", noFragmento, n)
		}
	}

	return nil
}