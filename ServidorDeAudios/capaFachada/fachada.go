package capaFachada

import (
	"fmt"

	repo "servidor.local/grpc-servidor/capaAccesoDatos"
	"servidor.local/grpc-servidor/modelos"
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
