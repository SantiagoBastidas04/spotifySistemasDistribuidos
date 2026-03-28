package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	controladores "servidor.local/servidorStreaming/capaControladores"
	pb "servidor.local/servidorStreaming/serviciosAudio"
)

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Error al abrir el puerto: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAudioServiceServer(grpcServer, &controladores.ControladorStreaming{})

	fmt.Println("ServidorDeStreaming escuchando en puerto 50051...")
	fmt.Println("Directorio de audios: ./audios/")

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}