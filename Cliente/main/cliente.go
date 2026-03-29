package main

import (
	"fmt"
	"log"

	"google.golang.org/grpc"
	vistas      "cliente.local/grpc-cliente/vistas"
	pbAudios    "servidor.local/servidorDeAudios/serviciosAudio"
	pbStreaming  "servidor.local/servidorStreaming/serviciosAudio"
)

func main() {
	// Conexión al ServidorDeAudios
	connAudios, err := grpc.Dial("localhost:50053", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error conectando con ServidorDeAudios: %v", err)
	}
	defer connAudios.Close()

	// Conexión al ServidorDeStreaming
	connStreaming, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error conectando con ServidorDeStreaming: %v", err)
	}
	defer connStreaming.Close()

	clienteAudios   := pbAudios.NewServiciosAudioClient(connAudios)
	clienteStreaming := pbStreaming.NewAudioServiceClient(connStreaming)

	fmt.Println("Conectado a ServidorDeAudios   (:50053)")
	fmt.Println("Conectado a ServidorDeStreaming (:50051)")

	// Cada operación crea su propio contexto con timeout adecuado
	vistas.MostrarMenuPrincipal(clienteAudios, clienteStreaming)
}