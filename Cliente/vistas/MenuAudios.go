package vistas

import (
	"bufio"
	"context"
	"fmt"
	"strings"

	ctrl "cliente.local/grpc-cliente/capaControladores"
	pbAudios   "servidor.local/servidorDeAudios/serviciosAudio"
	pbStreaming "servidor.local/servidorStreaming/serviciosAudio"
)

// MostrarListaAudios obtiene y despliega los audios de un tipo y permite seleccionar uno.
func MostrarListaAudios(
	clienteAudios   pbAudios.ServiciosAudioClient,
	clienteStreaming pbStreaming.AudioServiceClient,
	ctx context.Context,
	idTipo int32,
	reader *bufio.Reader,
) {
	resp, err := ctrl.ObtenerAudiosPorTipo(clienteAudios, ctx, idTipo)
	if err != nil {
		fmt.Printf("Error al obtener audios: %v\n", err)
		return
	}
	if resp.Estado.Codigo != 200 {
		fmt.Printf("Error del servidor: %s\n", resp.Estado.Mensaje)
		return
	}

	fmt.Println("\n--- Audios Disponibles ---")
	for i, a := range resp.Audios {
		fmt.Printf("  %d. %s\n", i+1, a.Titulo)
	}
	fmt.Println("  0. Volver")
	fmt.Print("Seleccione un audio: ")

	entrada, _ := reader.ReadString('\n')
	entrada = strings.TrimSpace(entrada)
	if entrada == "0" {
		return
	}

	var idx int
	if _, err := fmt.Sscanf(entrada, "%d", &idx); err != nil || idx < 1 || idx > len(resp.Audios) {
		fmt.Println("Opción no válida.")
		return
	}

	audioSeleccionado := resp.Audios[idx-1]
	MostrarDetalleAudio(clienteAudios, clienteStreaming, ctx, audioSeleccionado.IdAudio, audioSeleccionado.Titulo, reader)
}