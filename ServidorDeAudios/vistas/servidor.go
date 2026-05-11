package main

import (
    "fmt"
    "log"
    "net"
    "net/http"

    "google.golang.org/grpc"
    controladores "servidor.local/servidorDeAudios/capaControladores"
    pb "servidor.local/servidorDeAudios/serviciosAudio"
)

func main() {
    // El servidor gRPC corre en goroutine para no bloquear el REST
    go iniciarServidorGRPC()

    // El servidor REST bloquea el hilo principal
    iniciarServidorREST()
}

// iniciarServidorGRPC levanta el servidor gRPC en el puerto 50053.
// Es consumido por el Cliente Go para consultas y streaming.
func iniciarServidorGRPC() {
    listener, err := net.Listen("tcp", ":50053")
    if err != nil {
        log.Fatalf("Error abriendo puerto gRPC: %v", err)
    }

    grpcServer := grpc.NewServer()
    pb.RegisterServiciosAudioServer(grpcServer, &controladores.ControladorAudios{})

    fmt.Println("ServidorDeAudios [gRPC] escuchando en :50053")
    if err := grpcServer.Serve(listener); err != nil {
        log.Fatalf("Error en servidor gRPC: %v", err)
    }
}

// iniciarServidorREST levanta el servidor HTTP en el puerto 5000.
// Es consumido por el Administrador Java para almacenar y listar audios.
func iniciarServidorREST() {
    http.HandleFunc("/canciones/almacenamiento", controladores.AlmacenarAudio)
    http.HandleFunc("/canciones", controladores.ListarAudios)

    fmt.Println("ServidorDeAudios [REST] escuchando en :5000")
    if err := http.ListenAndServe(":5000", nil); err != nil {
        log.Fatalf("Error en servidor REST: %v", err)
    }
}