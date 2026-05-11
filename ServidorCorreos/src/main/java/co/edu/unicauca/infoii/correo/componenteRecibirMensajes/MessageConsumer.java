package co.edu.unicauca.infoii.correo.componenteRecibirMensajes;

import org.springframework.amqp.rabbit.annotation.RabbitListener;
import org.springframework.stereotype.Service;
import co.edu.unicauca.infoii.correo.DTOs.CancionAlmacenarDTOInput;
import co.edu.unicauca.infoii.correo.commons.Simulacion;

@Service
public class MessageConsumer {

    private static final String FRASE_MOTIVADORA =
            "¡La música es el idioma universal de la humanidad! ";

    @RabbitListener(queues = "notificaciones_canciones")
    public void recibirMensaje(CancionAlmacenarDTOInput cancion) {
        System.out.println("\nNuevo audio registrado. Simulando envío de correo...");
        Simulacion.simular(1000, "Enviando correo");

        System.out.println("\n========== CORREO SIMULADO ==========");
        System.out.println("-------------------------------------");
        System.out.println("ID Audio          : " + cancion.getIdAudio());
        System.out.println("Título            : " + cancion.getTitulo());
        System.out.println("Artista           : " + cancion.getArtista());
        System.out.println("Género            : " + cancion.getGenero());
        System.out.println("Fecha de registro : " + cancion.getFechaHoraRegistro());
        System.out.println("-------------------------------------");
        System.out.println(FRASE_MOTIVADORA);
        System.out.println("=====================================\n");
    }
}