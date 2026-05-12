package co.edu.unicauca.configuracion.servicios;

import java.io.BufferedReader;
import java.io.InputStreamReader;
import java.io.OutputStream;
import java.net.ServerSocket;
import java.net.Socket;
import java.nio.charset.StandardCharsets;
import co.edu.unicauca.capaDeControladores.RegistroAdministradores;

public class ServidorHTTPCallback {

    // Puerto que quedó asignado a esta instancia
    public static int puertoAsignado = -1;

    public static void iniciar() {
        // Intentar puertos desde 8080 hasta 8089
        for (int puerto = 8080; puerto <= 8089; puerto++) {
            try {
                final int puertoFinal = puerto;
                ServerSocket serverSocket = new ServerSocket(puertoFinal);
                puertoAsignado = puertoFinal;

                Thread hilo = new Thread(() -> {
                    System.out.println("[HTTP] Servidor de callbacks escuchando en :"
                            + puertoFinal);
                    while (true) {
                        try {
                            Socket cliente = serverSocket.accept();
                            new Thread(() -> manejarConexion(cliente)).start();
                        } catch (Exception e) {
                            System.out.println("[HTTP] Error aceptando conexión: "
                                    + e.getMessage());
                        }
                    }
                });
                hilo.setDaemon(true);
                hilo.start();
                return; // salir en cuanto encuentre un puerto libre

            } catch (Exception e) {
                // Puerto ocupado, intentar el siguiente
            }
        }
        System.out.println("[HTTP] No se encontró puerto disponible entre 8080 y 8089");
    }

    private static void manejarConexion(Socket cliente) {
        try (
            BufferedReader entrada = new BufferedReader(
                new InputStreamReader(cliente.getInputStream(), StandardCharsets.UTF_8));
            OutputStream salida = cliente.getOutputStream()
        ) {
            String linea;
            int contentLength = 0;
            while (!(linea = entrada.readLine()).isEmpty()) {
                if (linea.toLowerCase().startsWith("content-length:")) {
                    contentLength = Integer.parseInt(linea.split(":")[1].trim());
                }
            }

            char[] cuerpoChars = new char[contentLength];
            entrada.read(cuerpoChars, 0, contentLength);
            String cuerpo = new String(cuerpoChars);

            System.out.println("\n[HTTP] Notificación recibida: " + cuerpo);

            String idAudio   = extraerCampoJson(cuerpo, "idAudio");
            String fechaHora = extraerCampoJson(cuerpo, "fechaHoraReproduccion");

            RegistroAdministradores.notificarTodos(idAudio, fechaHora);

            String respuesta = "HTTP/1.1 200 OK\r\nContent-Length: 0\r\n\r\n";
            salida.write(respuesta.getBytes(StandardCharsets.UTF_8));
            salida.flush();

        } catch (Exception e) {
            System.out.println("[HTTP] Error procesando conexión: " + e.getMessage());
        } finally {
            try { cliente.close(); } catch (Exception ignored) {}
        }
    }

    private static String extraerCampoJson(String json, String campo) {
        String clave = "\"" + campo + "\":\"";
        int inicio = json.indexOf(clave);
        if (inicio == -1) return "";
        inicio += clave.length();
        int fin = json.indexOf("\"", inicio);
        return fin == -1 ? "" : json.substring(inicio, fin);
    }
}