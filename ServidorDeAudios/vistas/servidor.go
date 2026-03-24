package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"servidor.local/grpc-servidor/modelos"
	"servidor.local/grpc-servidor/servicios"
	alias "servidor.local/grpc-servidor/serviciosCancion"
)

var vectorMetadataAudios = make([]modelos.MetadataAudio, 5)

// Estructura que implementa el servicio
type servidorCanciones struct {
	alias.UnimplementedServiciosCancionesServer
}

// Implementación del procedimiento remoto buscarAudio
func (s *servidorCanciones) BuscarAudio(ctx context.Context, req *alias.PeticionDTO) (*alias.RespuestaMetadataAudioDTO, error) {

	titulo := req.GetTitulo()
	resp := servicios.BuscarAudio(titulo, vectorMetadataAudios)

	var respuesta alias.RespuestaMetadataAudioDTO
	respuesta.Codigo = int32(resp.Codigo)
	respuesta.Mensaje = resp.Mensaje

	if resp.Codigo == 200 {
		respuesta.ObjAudio = new(alias.MetadataAudio)
		respuesta.ObjAudio.Titulo = resp.ObjAudio.GetTitulo()
		respuesta.ObjAudio.Duracion = int32(resp.ObjAudio.GetDuracion())
		respuesta.ObjAudio.Tipo = resp.ObjAudio.GetTipo()
		respuesta.ObjAudio.Disponible = resp.ObjAudio.GetDisponible()

	}

	return &respuesta, nil

}

func main() {

	listener, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("Error al abrir el puerto: %v", err)
	}

	// Crear servidor gRPC
	grpcServer := grpc.NewServer()

	// Registrar el servicio
	alias.RegisterServiciosCancionesServer(grpcServer, &servidorCanciones{})

	// Cargar canciones en el vector
	servicios.CargarMetadataAudios(vectorMetadataAudios)

	// Iniciar el servidor
	log.Println("Servidor gRPC escuchando en puerto 50053...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
