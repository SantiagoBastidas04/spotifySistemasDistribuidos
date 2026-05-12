package capaFachada

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	repo "servidor.local/servidorStreaming/capaAccesoDatos"
	pb "servidor.local/servidorStreaming/serviciosAudio"
)

// puertosAdministradores lista todos los puertos donde pueden haber admins escuchando
var puertosAdministradores = []string{
	"http://localhost:8080/callback/reproduccion",
	"http://localhost:8081/callback/reproduccion",
	"http://localhost:8082/callback/reproduccion",
}

type notificacionReproduccion struct {
	IdAudio         string `json:"idAudio"`
	FechaHoraReprod string `json:"fechaHoraReproduccion"`
}

func EnviarFragmentosAudio(titulo string, stream pb.AudioService_AudioStreamServer) error {
	log.Printf("[Fachada] Preparando envío de fragmentos para: %s", titulo)

	archivo, err := repo.AbrirArchivo(titulo)
	if err != nil {
		return fmt.Errorf("no se pudo abrir el archivo: %w", err)
	}
	defer archivo.Close()

	// Notificar al inicio de la reproducción
	go notificarATodosLosAdmins(titulo)

	buf := make([]byte, 32*1024)
	noFragmento := 0

	for {
		n, err := archivo.Read(buf)
		if err == io.EOF {
			log.Printf("[Fachada] Envío completo: %s (%d fragmentos)", titulo, noFragmento)
			break
		}
		if err != nil {
			return fmt.Errorf("error leyendo archivo: %w", err)
		}
		if n > 0 {
			noFragmento++
			if err := stream.Send(&pb.AudioChunk{Data: buf[:n]}); err != nil {
				return fmt.Errorf("error enviando fragmento #%d: %w", noFragmento, err)
			}
			log.Printf("[Fachada] Fragmento #%d enviado (%d bytes)", noFragmento, n)
		}
	}

	return nil
}

// notificarATodosLosAdmins intenta notificar a cada puerto de la lista.
// Si un admin no está corriendo en ese puerto, simplemente lo ignora.
func notificarATodosLosAdmins(tituloAudio string) {
	notificacion := notificacionReproduccion{
		IdAudio:         tituloAudio,
		FechaHoraReprod: time.Now().Format("2006-01-02 15:04:05"),
	}
	body, err := json.Marshal(notificacion)
	if err != nil {
		log.Printf("[Callback] Error serializando notificación: %v", err)
		return
	}

	for _, url := range puertosAdministradores {
		resp, err := http.Post(url, "application/json", bytes.NewBuffer(body))
		if err != nil {
			// Admin no está corriendo en ese puerto, se ignora silenciosamente
			log.Printf("[Callback] Admin no disponible en %s", url)
			continue
		}
		resp.Body.Close()
		log.Printf("[Callback] Admin notificado en %s: %s", url, resp.Status)
	}
}
