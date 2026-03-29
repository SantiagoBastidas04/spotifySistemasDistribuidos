package capaControladores

import (
	"context"
	"fmt"
	"io"

	util "cliente.local/grpc-cliente/utilidades"
	pb "servidor.local/servidorStreaming/serviciosAudio"
)

// ReproducirAudio llama al RPC AudioStream del ServidorDeStreaming y reproduce el audio.
// Usa context.Background() sin timeout para no cortar audios largos.
func ReproducirAudio(cliente pb.AudioServiceClient, filename string) {
	fmt.Printf("[Cliente → RPC] Llamando AudioStream con filename=%s...\n", filename)

	// Sin timeout el streaming debe durar lo que dure el audio completo
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stream, err := cliente.AudioStream(ctx, &pb.AudioRequest{Filename: filename})
	if err != nil {
		fmt.Printf("Error al iniciar streaming: %v\n", err)
		return
	}

	fmt.Println("Recibiendo y reproduciendo audio en vivo...")
	reader, writer := io.Pipe()
	canalSincronizacion := make(chan struct{})

	go util.DecodificarReproducir(reader, canalSincronizacion)
	util.RecibirAudio(stream, writer, canalSincronizacion)
}