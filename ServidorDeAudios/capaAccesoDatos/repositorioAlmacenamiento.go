package capaAccesoDatos

import (
    "fmt"
    "os"
    "path/filepath"
)

const directorioAudios = "audios"

// GuardarArchivoFisico persiste el binario MP3 en disco.
// El nombre del archivo usa el título tal cual, para que el ServidorStreaming
// pueda encontrarlo por título cuando el cliente pida reproducirlo.
func GuardarArchivoFisico(titulo, artista, genero string, data []byte) error {
    if err := os.MkdirAll(directorioAudios, os.ModePerm); err != nil {
        return fmt.Errorf("error creando directorio de audios: %w", err)
    }

    // El ServidorStreaming busca los archivos como: audios/<titulo>.mp3
    // Por eso el nombre del archivo es solo el título.
    nombreArchivo := titulo + ".mp3"
    ruta := filepath.Join(directorioAudios, nombreArchivo)

    if err := os.WriteFile(ruta, data, 0644); err != nil {
        return fmt.Errorf("error guardando archivo en disco: %w", err)
    }

    fmt.Printf("[Disco] Archivo guardado: %s\n", ruta)
    return nil
}