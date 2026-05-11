package capaControladores

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"

    fachada "servidor.local/servidorDeAudios/capaFachada"
)

// RespuestaAlmacenamiento es el JSON que se devuelve al Administrador
// tras almacenar un audio exitosamente.
type RespuestaAlmacenamiento struct {
    IdAudio int    `json:"idAudio"`
    Mensaje string `json:"mensaje"`
}

// AlmacenarAudio maneja POST /canciones/almacenamiento
// Espera multipart/form-data con campos: titulo, artista, genero y archivo (MP3).
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

    titulo  := r.FormValue("titulo")
    artista := r.FormValue("artista")
    genero  := r.FormValue("genero")

    idAsignado, err := fachada.AlmacenarCancion(titulo, artista, genero, data)
    if err != nil {
        http.Error(w, "Error almacenando canción: "+err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(RespuestaAlmacenamiento{
        IdAudio: idAsignado,
        Mensaje: fmt.Sprintf("Audio '%s' almacenado correctamente con id %d", titulo, idAsignado),
    })
}

// ListarAudios maneja GET /canciones
// Devuelve los audios de tipo Música (tipo 1) en formato JSON.
func ListarAudios(w http.ResponseWriter, r *http.Request) {
    fmt.Println("[REST] GET /canciones recibido")

    audios, err := fachada.ObtenerAudiosPorTipo(1)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(audios)
}