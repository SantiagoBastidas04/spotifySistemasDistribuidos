package capaControladores

import (
	"context"
	"fmt"

	pb "servidor.local/servidorDeAudios/serviciosAudio"
)

// ObtenerTipos llama al RPC del ServidorDeAudios para obtener los tipos de audio registrados.
func ObtenerTipos(cliente pb.ServiciosAudioClient, ctx context.Context) (*pb.RespuestaTipos, error) {
	fmt.Println("[Cliente -> RPC] Llamando ObtenerTipos...")
	return cliente.ObtenerTipos(ctx, &pb.SolicitudVacia{})
}

// ObtenerAudiosPorTipo llama al RPC del ServidorDeAudios filtrando por tipo de audio.
func ObtenerAudiosPorTipo(cliente pb.ServiciosAudioClient, ctx context.Context, idTipo int32) (*pb.RespuestaListaAudios, error) {
	fmt.Printf("[Cliente -> RPC] Llamando ObtenerAudiosPorTipo con idTipo=%d...\n", idTipo)
	return cliente.ObtenerAudiosPorTipo(ctx, &pb.SolicitudPorTipo{IdTipo: idTipo})
}

// ObtenerMetadata llama al RPC del ServidorDeAudios para obtener el detalle de un audio.
func ObtenerMetadata(cliente pb.ServiciosAudioClient, ctx context.Context, idAudio int32) (*pb.RespuestaMetadata, error) {
	fmt.Printf("[Cliente -> RPC] Llamando ObtenerMetadata con idAudio=%d...\n", idAudio)
	return cliente.ObtenerMetadata(ctx, &pb.SolicitudPorId{IdAudio: idAudio})
}