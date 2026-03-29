package capaControladores

import (
	"context"
	"fmt"
	"io"

	util "cliente.local/grpc-cliente/utilidades"
	pb "servidor.local/servidorStreaming/serviciosAudio"
)

// ReproducirAudio llama al RPC AudioStream del ServidorDeStreaming y reproduce el audio recibido.
func ReproducirAudio(cliente pb.AudioServiceClient, ctx context.Context, filename string) {
	fmt.Printf("[Cliente -> RPC] Llamando AudioStream con filename=%s...\n", filename)

	stream, err := cliente.AudioStream(ctx, &pb.AudioRequest{Filename: filename})
	if err != nil {
		fmt.Printf("Error al iniciar streaming: %v\n", err)
		return
	}

	fmt.Println("Recibiendo y reproduciendo audio en vivo...")
	reader, writer := io.Pipe()
	canalSincronizacion := make(chan struct{})

	// Goroutine: decodifica y reproduce los fragmentos recibidos
	go util.DecodificarReproducir(reader, canalSincronizacion)

	// Recibe los fragmentos del servidor y los escribe en el pipe
	util.RecibirAudio(stream, writer, canalSincronizacion)
}