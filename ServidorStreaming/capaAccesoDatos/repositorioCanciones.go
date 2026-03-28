package capaAccesoDatos
import (
	"fmt"
	"os"
	"path/filepath"
)

const directorioAudios = "audios"

// AbrirArchivo busca el archivo .mp3 cuyo nombre base coincide con el título recibido.
// Retorna el *os.File abierto o un error si no existe.
func AbrirArchivo(titulo string) (*os.File, error) {
	// Construye la ruta: audios/<titulo>.mp3
	ruta := filepath.Join(directorioAudios, titulo+".mp3")
	archivo, err := os.Open(ruta)
	if err != nil {
		return nil, fmt.Errorf("archivo no encontrado para '%s': %w", titulo, err)
	}
	return archivo, nil
}