package co.edu.unicauca.main;

import co.edu.unicauca.capaDeControladores.CallBackImp;
import co.edu.unicauca.capaDeControladores.RegistroAdministradores;
import co.edu.unicauca.configuracion.servicios.ClienteREST;
import co.edu.unicauca.configuracion.servicios.ServidorHTTPCallback;
import co.edu.unicauca.utilidades.UtilidadesConsola;

public class Main {

    public static void main(String[] args) {
        iniciarServidorCallback();
        registrarEsteAdministrador();
        mostrarMenu();
    }

    private static void iniciarServidorCallback() {
        try {
            ServidorHTTPCallback.iniciar();
        } catch (Exception e) {
            System.out.println("[Main] Error iniciando servidor HTTP: "
                    + e.getMessage());
        }
    }

    private static void registrarEsteAdministrador() {
        try {
            CallBackImp miCallback = new CallBackImp();
            RegistroAdministradores.registrar(miCallback);
            System.out.println("Administrador listo. Esperando notificaciones...");
        } catch (Exception e) {
            System.out.println("[Main] Error registrando callback: "
                    + e.getMessage());
        }
    }

    private static void mostrarMenu() {
        int opcion;
        do {
            System.out.println("\n===== MENÚ ADMINISTRADOR =====");
            System.out.println("1. Almacenar audio y metadatos");
            System.out.println("2. Listar audios");
            System.out.println("0. Salir");
            System.out.print("Opción: ");
            opcion = UtilidadesConsola.leerEntero();

            switch (opcion) {
                case 1 -> almacenarAudio();
                case 2 -> listarAudios();
                case 0 -> System.out.println("Saliendo...");
                default -> System.out.println("Opción inválida.");
            }
        } while (opcion != 0);
    }

    private static void almacenarAudio() {
        System.out.print("Título: ");
        String titulo  = UtilidadesConsola.leerCadena();
        System.out.print("Artista: ");
        String artista = UtilidadesConsola.leerCadena();
        System.out.print("Género: ");
        String genero  = UtilidadesConsola.leerCadena();
        System.out.print("Ruta del archivo MP3: ");
        String ruta    = UtilidadesConsola.leerCadena();

        boolean ok = ClienteREST.almacenarAudio(titulo, artista, genero, ruta);
        System.out.println(ok
                ? "Audio almacenado correctamente."
                : "Error al almacenar el audio.");
    }

    private static void listarAudios() {
        String resultado = ClienteREST.listarAudios();
        System.out.println("\n=== Audios disponibles ===");
        System.out.println(resultado);
    }
}