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

// MostrarDetalleAudio obtiene y muestra los metadatos completos del audio seleccionado.
func MostrarDetalleAudio(
	clienteAudios   pbAudios.ServiciosAudioClient,
	clienteStreaming pbStreaming.AudioServiceClient,
	ctx context.Context,
	idAudio int32,
	tituloAudio string,
	reader *bufio.Reader,
) {
	resp, err := ctrl.ObtenerMetadata(clienteAudios, ctx, idAudio)
	if err != nil {
		fmt.Printf("Error al obtener metadatos: %v\n", err)
		return
	}
	if resp.Estado.Codigo != 200 {
		fmt.Printf("Error del servidor: %s\n", resp.Estado.Mensaje)
		return
	}

	fmt.Println("\n--- Detalle del Audio ---")
	imprimirMetadato(resp)

	for {
		fmt.Println("\n  1. Reproducir audio")
		fmt.Println("  2. Volver al menú principal")
		fmt.Print("Seleccione: ")

		opcion, _ := reader.ReadString('\n')
		opcion = strings.TrimSpace(opcion)

		switch opcion {
		case "1":
			ctrl.ReproducirAudio(clienteStreaming, ctx, tituloAudio)
		case "2":
			return
		default:
			fmt.Println("Opción no válida.")
		}
	}
}

// imprimirMetadato muestra los campos según el tipo concreto del oneof.
func imprimirMetadato(resp *pbAudios.RespuestaMetadata) {
	switch m := resp.Metadato.(type) {

	case *pbAudios.RespuestaMetadata_Musica:
		mu := m.Musica
		fmt.Println("  Tipo:             Música")
		fmt.Printf("  Título:           %s\n", mu.Titulo)
		fmt.Printf("  Artista:          %s\n", mu.ArtistaPrincipal)
		fmt.Printf("  Álbum:            %s\n", mu.Album)
		fmt.Printf("  Género:           %s\n", mu.GeneroMusical)
		fmt.Printf("  Sello:            %s\n", mu.SelloDiscografico)
		fmt.Printf("  Año lanzamiento:  %d\n", mu.AnioLanzamiento)

	case *pbAudios.RespuestaMetadata_Podcast:
		po := m.Podcast
		fmt.Println("  Tipo:             Podcast")
		fmt.Printf("  Podcast:          %s\n", po.TituloPodcast)
		fmt.Printf("  Episodio:         %s\n", po.TituloEpisodio)
		fmt.Printf("  Anfitrión:        %s\n", po.Anfitrion)
		fmt.Printf("  Temporada/EP:     %s\n", po.TemporadaEpisodio)
		fmt.Printf("  Clasificación:    %s\n", po.ClasificacionContenido)
		fmt.Printf("  Notas:            %s\n", po.NotasShow)

	case *pbAudios.RespuestaMetadata_Audiolibro:
		al := m.Audiolibro
		fmt.Println("  Tipo:             Audiolibro")
		fmt.Printf("  Título:           %s\n", al.Titulo)
		fmt.Printf("  Autor:            %s\n", al.Autor)
		fmt.Printf("  Narrador:         %s\n", al.Narrador)
		fmt.Printf("  Editorial:        %s\n", al.Editorial)
		fmt.Printf("  ISBN:             %s\n", al.Isbn)
		fmt.Printf("  Capítulo:         %s\n", al.Capitulo)

	case *pbAudios.RespuestaMetadata_RuidoBlanco:
		rb := m.RuidoBlanco
		fmt.Println("  Tipo:             Ruido Blanco")
		fmt.Printf("  Tipo de sonido:   %s\n", rb.TipoSonido)
		fmt.Printf("  Fuente:           %s\n", rb.FuenteAudio)
		fmt.Printf("  Uso sugerido:     %s\n", rb.UsoSugerido)
		fmt.Printf("  Proveedor:        %s\n", rb.ProveedorContenido)
		fmt.Printf("  Duración bucle:   %d min\n", rb.DuracionBucle)
		fmt.Printf("  Frecuencia:       %s\n", rb.FrecuenciaDominante)

	default:
		fmt.Println("  Tipo de metadato desconocido.")
	}
}