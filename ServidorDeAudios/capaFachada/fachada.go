package capaFachada

import (
	"fmt"

	repo "servidor.local/servidorDeAudios/capaAccesoDatos"
	"servidor.local/servidorDeAudios/modelos"
	cola "servidor.local/servidorDeAudios/componenteConexionCola"
)

func ObtenerTipos() []modelos.TipoAudio {
	fmt.Println("[RPC] ObtenerTipos llamado")
	return repo.ObtenerTipos()
}

func ObtenerAudiosPorTipo(idTipo int) ([]modelos.ResumenAudio, error) {
	fmt.Printf("[RPC] ObtenerAudiosPorTipo llamado con idTipo=%d\n", idTipo)

	tipos := repo.ObtenerTipos()
	tipoValido := false
	for _, t := range tipos {
		if t.GetId() == idTipo {
			tipoValido = true
			break
		}
	}

	if !tipoValido {
		return nil, fmt.Errorf("tipo de audio con id %d no existe", idTipo)
	}

	return repo.ObtenerAudiosPorTipo(idTipo), nil
}

func ObtenerMetadata(idAudio int) (interface{}, int, error) {
	fmt.Printf("[RPC] ObtenerMetadata llamado con idAudio=%d\n", idAudio)

	idTipo := repo.ObtenerTipoPorAudio(idAudio)

	switch idTipo {
	case 1:
		musica, encontrado := repo.ObtenerMusica(idAudio)
		if !encontrado {
			return nil, -1, fmt.Errorf("música con id %d no encontrada", idAudio)
		}
		return musica, 1, nil
	case 2:
		podcast, encontrado := repo.ObtenerPodcast(idAudio)
		if !encontrado {
			return nil, -1, fmt.Errorf("podcast con id %d no encontrado", idAudio)
		}
		return podcast, 2, nil
	case 3:
		audiolibro, encontrado := repo.ObtenerAudiolibro(idAudio)
		if !encontrado {
			return nil, -1, fmt.Errorf("audiolibro con id %d no encontrado", idAudio)
		}
		return audiolibro, 3, nil
	case 4:
		ruidoBlanco, encontrado := repo.ObtenerRuidoBlanco(idAudio)
		if !encontrado {
			return nil, -1, fmt.Errorf("ruido blanco con id %d no encontrado", idAudio)
		}
		return ruidoBlanco, 4, nil
	default:
		return nil, -1, fmt.Errorf("audio con id %d no encontrado", idAudio)
	}
}
func AlmacenarCancion(titulo, artista, genero string, data []byte) (int, error) {
    fmt.Printf("[Fachada] AlmacenarCancion titulo=%s artista=%s genero=%s\n", titulo, artista, genero)

    // Paso 1: guardar el archivo MP3 en disco
    if err := repo.GuardarArchivoFisico(titulo, artista, genero, data); err != nil {
        return 0, fmt.Errorf("error guardando archivo: %w", err)
    }

    // Paso 2: registrar metadatos en memoria para que el gRPC los sirva
    nuevaMusica := modelos.Musica{}
    nuevaMusica.SetTitulo(titulo)
    nuevaMusica.SetArtistaPrincipal(artista)
    nuevaMusica.SetGeneroMusical(genero)
    nuevaMusica.SetAlbum("Sin álbum")
    nuevaMusica.SetSelloDiscografico("Independiente")
    nuevaMusica.SetAnioLanzamiento(0)
    idAsignado := repo.AgregarMusica(nuevaMusica)

    fmt.Printf("[Fachada] Audio registrado en memoria con id=%d\n", idAsignado)

    // Paso 3: notificar asíncronamente al servidor de correos vía RabbitMQ
    go publicarEnCola(idAsignado, titulo, artista, genero)

    return idAsignado, nil
}

// publicarEnCola crea una conexión RabbitMQ, publica la notificación y la cierra.
// Se ejecuta en una goroutine para no bloquear la respuesta al Administrador.
func publicarEnCola(idAudio int, titulo, artista, genero string) {
    publisher, err := cola.NewRabbitPublisher()
    if err != nil {
        fmt.Printf("[Fachada] Error conectando RabbitMQ: %v\n", err)
        return
    }
    defer publisher.Cerrar()

    err = publisher.PublicarNotificacion(cola.NotificacionCancion{
        IdAudio: idAudio,
        Titulo:  titulo,
        Artista: artista,
        Genero:  genero,
    })
    if err != nil {
        fmt.Printf("[Fachada] Error publicando en cola: %v\n", err)
    }
}