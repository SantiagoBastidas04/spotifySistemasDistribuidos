package vistas


import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	ctrl "cliente.local/grpc-cliente/capaControladores"
	pbAudios   "servidor.local/servidorDeAudios/serviciosAudio"
	pbStreaming "servidor.local/servidorStreaming/serviciosAudio"
)

// MostrarMenuPrincipal presenta el menú raíz de la aplicación.
func MostrarMenuPrincipal(
	clienteAudios   pbAudios.ServiciosAudioClient,
	clienteStreaming pbStreaming.AudioServiceClient,
	ctx context.Context,
) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n╔══════════════════════════════════╗")
		fmt.Println("║   Sistema de Audio Distribuido   ║")
		fmt.Println("╠══════════════════════════════════╣")
		fmt.Println("║  1. Ver tipos de audio           ║")
		fmt.Println("║  2. Salir                        ║")
		fmt.Println("╚══════════════════════════════════╝")
		fmt.Print("Seleccione una opción: ")

		opcion, _ := reader.ReadString('\n')
		opcion = strings.TrimSpace(opcion)

		switch opcion {
		case "1":
			mostrarTipos(clienteAudios, clienteStreaming, ctx, reader)
		case "2":
			fmt.Println("Hasta luego.")
			return
		default:
			fmt.Println("Opción no válida. Intente de nuevo.")
		}
	}
}

// mostrarTipos obtiene los tipos de audio y deja al usuario seleccionar uno.
func mostrarTipos(
	clienteAudios   pbAudios.ServiciosAudioClient,
	clienteStreaming pbStreaming.AudioServiceClient,
	ctx context.Context,
	reader *bufio.Reader,
) {
	resp, err := ctrl.ObtenerTipos(clienteAudios, ctx)
	if err != nil {
		fmt.Printf("Error al obtener tipos: %v\n", err)
		return
	}
	if resp.Estado.Codigo != 200 {
		fmt.Printf("Error del servidor: %s\n", resp.Estado.Mensaje)
		return
	}

	fmt.Println("\n--- Tipos de Audio Disponibles ---")
	for _, t := range resp.Tipos {
		fmt.Printf("  %d. %s\n", t.Id, t.Nombre)
	}
	fmt.Println("  0. Volver")
	fmt.Print("Seleccione un tipo: ")

	entrada, _ := reader.ReadString('\n')
	entrada = strings.TrimSpace(entrada)
	if entrada == "0" {
		return
	}

	var idTipo int32
	if _, err := fmt.Sscanf(entrada, "%d", &idTipo); err != nil {
		fmt.Println("Entrada no válida.")
		return
	}

	// Validar que el id existe en la lista retornada
	valido := false
	for _, t := range resp.Tipos {
		if t.Id == idTipo {
			valido = true
			break
		}
	}
	if !valido {
		fmt.Println("Tipo no encontrado.")
		return
	}

	MostrarListaAudios(clienteAudios, clienteStreaming, ctx, idTipo, reader)
}