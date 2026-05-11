package co.edu.unicauca.configuracion.servicios;

import com.sun.net.httpserver.HttpExchange;
import com.sun.net.httpserver.HttpServer;
import java.io.InputStream;
import java.net.InetSocketAddress;
import java.nio.charset.StandardCharsets;
import co.edu.unicauca.capaDeControladores.RegistroAdministradores;


public class ServidorHTTPCallback {

    private static final int PUERTO_HTTP = 8080;

    public static void iniciar() throws Exception {
        HttpServer servidor = HttpServer.create(
                new InetSocketAddress(PUERTO_HTTP), 0);
        servidor.createContext("/callback/reproduccion",
                ServidorHTTPCallback::manejarCallback);
        servidor.start();
        System.out.println("[HTTP] Servidor de callbacks escuchando en :"
                + PUERTO_HTTP);
    }

    private static void manejarCallback(HttpExchange exchange) {
        try {
            if (!"POST".equals(exchange.getRequestMethod())) {
                exchange.sendResponseHeaders(405, -1);
                return;
            }

            InputStream is = exchange.getRequestBody();
            String cuerpo = new String(is.readAllBytes(), StandardCharsets.UTF_8);
            System.out.println("[HTTP] Notificación recibida: " + cuerpo);

            String idAudio   = extraerCampoJson(cuerpo, "idAudio");
            String fechaHora = extraerCampoJson(cuerpo, "fechaHoraReproduccion");

            RegistroAdministradores.notificarTodos(idAudio, fechaHora);

            exchange.sendResponseHeaders(200, -1);
        } catch (Exception e) {
            System.out.println("[HTTP] Error procesando callback: " + e.getMessage());
        } finally {
            exchange.close();
        }
    }

    // Extrae el valor de un campo string de un JSON simple sin dependencias externas
    private static String extraerCampoJson(String json, String campo) {
        String clave = "\"" + campo + "\":\"";
        int inicio = json.indexOf(clave);
        if (inicio == -1) return "";
        inicio += clave.length();
        int fin = json.indexOf("\"", inicio);
        return fin == -1 ? "" : json.substring(inicio, fin);
    }
}