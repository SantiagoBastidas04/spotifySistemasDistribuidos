package capaControladores

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	fachada "servidor.local/servidorDeAudios/capaFachada"
)

// ResumenAudioDTO es la representación serializable de un audio para la API REST.
// Usa campos exportados (mayúscula) para que encoding/json los incluya.
type ResumenAudioDTO struct {
	IdAudio int    `json:"idAudio"`
	Titulo  string `json:"titulo"`
	IdTipo  int    `json:"idTipo"`
}

type RespuestaAlmacenamiento struct {
	IdAudio int    `json:"idAudio"`
	Mensaje string `json:"mensaje"`
}

func AlmacenarAudio(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[REST] POST /canciones/almacenamiento recibido")

	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseMultipartForm(50 << 20); err != nil {
		http.Error(w, "Error parseando formulario multipart", http.StatusBadRequest)
		return
	}

	archivo, _, err := r.FormFile("archivo")
	if err != nil {
		http.Error(w, "Error leyendo campo 'archivo'", http.StatusBadRequest)
		return
	}
	defer archivo.Close()

	data, err := io.ReadAll(archivo)
	if err != nil {
		http.Error(w, "Error leyendo bytes del archivo", http.StatusInternalServerError)
		return
	}

	titulo := r.FormValue("titulo")
	artista := r.FormValue("artista")
	genero := r.FormValue("genero")

	fmt.Printf("[REST] Datos recibidos -> titulo='%s' artista='%s' genero='%s'\n",
		titulo, artista, genero)

	idAsignado, err := fachada.AlmacenarCancion(titulo, artista, genero, data)
	if err != nil {
		http.Error(w, "Error almacenando canción: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(RespuestaAlmacenamiento{
		IdAudio: idAsignado,
		Mensaje: fmt.Sprintf("Audio '%s' almacenado con id %d", titulo, idAsignado),
	})
}

// ListarAudios maneja GET /canciones
// Convierte los modelos internos a DTOs exportables antes de serializar.
func ListarAudios(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[REST] GET /canciones recibido")

	audios, err := fachada.ObtenerTodosLosAudios()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Convertir a DTOs exportables
	var dtos []ResumenAudioDTO
	for _, a := range audios {
		dtos = append(dtos, ResumenAudioDTO{
			IdAudio: a.GetIdAudio(),
			Titulo:  a.GetTitulo(),
			IdTipo:  a.GetIdTipo(),
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dtos)
}
