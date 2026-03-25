package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	controladores "servidor.local/grpc-servidor/capaControladores"
	pb "servidor.local/grpc-servidor/serviciosAudio"
)

func main() {
	listener, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("Error al abrir el puerto: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterServiciosAudioServer(grpcServer, &controladores.ControladorAudios{})

	fmt.Println("ServidorDeAudios escuchando en puerto 50053...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}