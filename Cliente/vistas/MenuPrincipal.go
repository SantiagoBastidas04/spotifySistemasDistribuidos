package vistas

import (
    "bufio"
    "fmt"
    "os"
    "strings"

    ctrl        "cliente.local/grpc-cliente/capaControladores"
    pbAudios    "servidor.local/servidorDeAudios/serviciosAudio"
    pbStreaming  "servidor.local/servidorStreaming/serviciosAudio"
)

const archivoUsuarios = "usuarios.txt"

// MostrarMenuPrincipal presenta el menú raíz de la aplicación.
func MostrarMenuPrincipal(
    clienteAudios   pbAudios.ServiciosAudioClient,
    clienteStreaming pbStreaming.AudioServiceClient,
) {
    reader := bufio.NewReader(os.Stdin)

    if !menuAutenticacion(reader) {
        fmt.Println("Saliendo.")
        return
    }

    for {
        fmt.Println("\n|==================================|")
        fmt.Println("|   Sistema de Audio Distribuido   |")
        fmt.Println("|==================================|")
        fmt.Println("|  1. Ver tipos de audio           |")
        fmt.Println("|  2. Salir                        |")
        fmt.Println("|==================================|")
        fmt.Print("Seleccione una opción: ")

        opcion, _ := reader.ReadString('\n')
        opcion = strings.TrimSpace(opcion)

        switch opcion {
        case "1":
            mostrarTipos(clienteAudios, clienteStreaming, reader)
        case "2":
            fmt.Println("Hasta luego.")
            return
        default:
            fmt.Println("Opción no válida. Intente de nuevo.")
        }
    }
}

// menuAutenticacion muestra las opciones de login o registro.
func menuAutenticacion(reader *bufio.Reader) bool {
    for {
        fmt.Println("\n===== BIENVENIDO =====")
        fmt.Println("1. Iniciar sesión")
        fmt.Println("2. Registrarse")
        fmt.Println("0. Salir")
        fmt.Print("Opción: ")

        opcion, _ := reader.ReadString('\n')
        opcion = strings.TrimSpace(opcion)

        switch opcion {
        case "1":
            return iniciarSesion(reader)
        case "2":
            registrarUsuario(reader)
        case "0":
            return false
        default:
            fmt.Println("Opción no válida.")
        }
    }
}

// iniciarSesion verifica nickname y contraseña contra el archivo de usuarios.
func iniciarSesion(reader *bufio.Reader) bool {
    fmt.Println("\n===== INICIO DE SESIÓN =====")
    fmt.Print("Nickname: ")
    nickname, _ := reader.ReadString('\n')
    nickname = strings.TrimSpace(nickname)

    fmt.Print("Contraseña: ")
    password, _ := reader.ReadString('\n')
    password = strings.TrimSpace(password)

    usuarios := cargarUsuarios()
    pass, existe := usuarios[nickname]
    if !existe || pass != password {
        fmt.Println("Credenciales incorrectas.")
        return false
    }

    fmt.Printf("Bienvenido, %s!\n", nickname)
    return true
}

// registrarUsuario solicita nickname y contraseña y los guarda en el archivo.
func registrarUsuario(reader *bufio.Reader) {
    fmt.Println("\n===== REGISTRO =====")
    fmt.Print("Nickname: ")
    nickname, _ := reader.ReadString('\n')
    nickname = strings.TrimSpace(nickname)

    fmt.Print("Contraseña: ")
    password, _ := reader.ReadString('\n')
    password = strings.TrimSpace(password)

    usuarios := cargarUsuarios()
    if _, existe := usuarios[nickname]; existe {
        fmt.Println("Ese nickname ya está registrado.")
        return
    }

    if err := guardarUsuario(nickname, password); err != nil {
        fmt.Printf("Error al registrar: %v\n", err)
        return
    }

    fmt.Printf("Usuario '%s' registrado correctamente.\n", nickname)
}

// cargarUsuarios lee el archivo de usuarios y retorna un mapa nickname -> contraseña.
func cargarUsuarios() map[string]string {
    usuarios := make(map[string]string)

    archivo, err := os.Open(archivoUsuarios)
    if err != nil {
        // Si el archivo no existe aún, retorna mapa vacío
        return usuarios
    }
    defer archivo.Close()

    scanner := bufio.NewScanner(archivo)
    for scanner.Scan() {
        linea := scanner.Text()
        partes := strings.SplitN(linea, ":", 2)
        if len(partes) == 2 {
            usuarios[partes[0]] = partes[1]
        }
    }

    return usuarios
}

// guardarUsuario agrega una línea "nickname:contraseña" al archivo de usuarios.
func guardarUsuario(nickname, password string) error {
    archivo, err := os.OpenFile(archivoUsuarios, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        return fmt.Errorf("error abriendo archivo de usuarios: %w", err)
    }
    defer archivo.Close()

    _, err = fmt.Fprintf(archivo, "%s:%s\n", nickname, password)
    return err
}

// mostrarTipos obtiene los tipos de audio y deja al usuario seleccionar uno.
func mostrarTipos(
    clienteAudios   pbAudios.ServiciosAudioClient,
    clienteStreaming pbStreaming.AudioServiceClient,
    reader *bufio.Reader,
) {
    resp, err := ctrl.ObtenerTipos(clienteAudios)
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

    MostrarListaAudios(clienteAudios, clienteStreaming, idTipo, reader)
}