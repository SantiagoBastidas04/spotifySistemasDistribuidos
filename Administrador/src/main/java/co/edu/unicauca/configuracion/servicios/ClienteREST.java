package co.edu.unicauca.configuracion.servicios;

import java.io.*;
import java.net.HttpURLConnection;
import java.net.URL;
import java.nio.file.Files;

/**
 * Cliente REST que se comunica con el ServidorDeAudios (Go).
 * Usa multipart/form-data para enviar el archivo MP3 con sus metadatos.
 */
public class ClienteREST {

    private static final String URL_ALMACENAR =
            "http://localhost:5000/canciones/almacenamiento";
    private static final String URL_LISTAR =
            "http://localhost:5000/canciones";
    private static final String BOUNDARY =
            "----FormBoundary7MA4YWxkTrZu0gW";

    public static boolean almacenarAudio(
            String titulo, String artista, String genero, String rutaArchivo) {
        try {
            File archivo = new File(rutaArchivo);
            if (!archivo.exists()) {
                System.out.println("Archivo no encontrado: " + rutaArchivo);
                return false;
            }

            HttpURLConnection conn = abrirConexionMultipart(URL_ALMACENAR);

            try (OutputStream os = conn.getOutputStream();
                 PrintWriter writer = new PrintWriter(
                         new OutputStreamWriter(os, "UTF-8"), true)) {

                agregarCampoTexto(writer, "titulo",  titulo);
                agregarCampoTexto(writer, "artista", artista);
                agregarCampoTexto(writer, "genero",  genero);
                agregarArchivo(writer, os, archivo);
                writer.append("--" + BOUNDARY + "--").append("\r\n");
            }

            int codigo = conn.getResponseCode();
            System.out.println("[REST] Respuesta almacenar: " + codigo);
            return codigo == 201;

        } catch (Exception e) {
            System.out.println("[REST] Error al almacenar: " + e.getMessage());
            return false;
        }
    }

    public static String listarAudios() {
        try {
            HttpURLConnection conn =
                    (HttpURLConnection) new URL(URL_LISTAR).openConnection();
            conn.setRequestMethod("GET");

            try (BufferedReader br = new BufferedReader(
                    new InputStreamReader(conn.getInputStream()))) {
                StringBuilder sb = new StringBuilder();
                String linea;
                while ((linea = br.readLine()) != null) {
                    sb.append(linea).append("\n");
                }
                return sb.toString();
            }
        } catch (Exception e) {
            return "Error al listar: " + e.getMessage();
        }
    }

    private static HttpURLConnection abrirConexionMultipart(String url)
            throws Exception {
        HttpURLConnection conn =
                (HttpURLConnection) new URL(url).openConnection();
        conn.setDoOutput(true);
        conn.setRequestMethod("POST");
        conn.setRequestProperty("Content-Type",
                "multipart/form-data; boundary=" + BOUNDARY);
        return conn;
    }

    private static void agregarCampoTexto(
            PrintWriter writer, String nombre, String valor) {
        writer.append("--" + BOUNDARY).append("\r\n");
        writer.append("Content-Disposition: form-data; name=\""
                + nombre + "\"").append("\r\n");
        writer.append("\r\n");
        writer.append(valor).append("\r\n");
    }

    private static void agregarArchivo(
            PrintWriter writer, OutputStream os, File archivo)
            throws IOException {
        writer.append("--" + BOUNDARY).append("\r\n");
        writer.append("Content-Disposition: form-data; name=\"archivo\"; filename=\""
                + archivo.getName() + "\"").append("\r\n");
        writer.append("Content-Type: audio/mpeg").append("\r\n");
        writer.append("\r\n");
        writer.flush();
        Files.copy(archivo.toPath(), os);
        os.flush();
        writer.append("\r\n");
    }
}