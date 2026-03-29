package capaControladores

import (
	"context"
	"fmt"
	"time"

	pb "servidor.local/servidorDeAudios/serviciosAudio"
)

const timeoutMetadata = 10 * time.Second

// ObtenerTipos llama al RPC del ServidorDeAudios para obtener los tipos de audio.
func ObtenerTipos(cliente pb.ServiciosAudioClient) (*pb.RespuestaTipos, error) {
	fmt.Println("[Cliente -> RPC] Llamando ObtenerTipos...")
	ctx, cancel := context.WithTimeout(context.Background(), timeoutMetadata)
	defer cancel()
	return cliente.ObtenerTipos(ctx, &pb.SolicitudVacia{})
}

// ObtenerAudiosPorTipo llama al RPC del ServidorDeAudios filtrando por tipo.
func ObtenerAudiosPorTipo(cliente pb.ServiciosAudioClient, idTipo int32) (*pb.RespuestaListaAudios, error) {
	fmt.Printf("[Cliente -> RPC] Llamando ObtenerAudiosPorTipo con idTipo=%d...\n", idTipo)
	ctx, cancel := context.WithTimeout(context.Background(), timeoutMetadata)
	defer cancel()
	return cliente.ObtenerAudiosPorTipo(ctx, &pb.SolicitudPorTipo{IdTipo: idTipo})
}

// ObtenerMetadata llama al RPC del ServidorDeAudios para obtener el detalle de un audio.
func ObtenerMetadata(cliente pb.ServiciosAudioClient, idAudio int32) (*pb.RespuestaMetadata, error) {
	fmt.Printf("[Cliente -> RPC] Llamando ObtenerMetadata con idAudio=%d...\n", idAudio)
	ctx, cancel := context.WithTimeout(context.Background(), timeoutMetadata)
	defer cancel()
	return cliente.ObtenerMetadata(ctx, &pb.SolicitudPorId{IdAudio: idAudio})
}