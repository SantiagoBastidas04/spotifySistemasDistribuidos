package co.edu.unicauca.capaDeControladores;

import java.rmi.RemoteException;
import java.util.ArrayList;
import java.util.List;
import co.edu.unicauca.fachadaServices.DTO.NotificacionReproduccionDTO;


public class RegistroAdministradores {

    private static final List<CallBackInt> administradoresRegistrados = new ArrayList<>();

    public static synchronized void registrar(CallBackInt callback) {
        administradoresRegistrados.add(callback);
        System.out.println("[Registro] Administrador registrado. Total activos: ");
    }

    public static synchronized void notificarTodos(String idAudio, String fechaHoraReproduccion) {
        System.out.println("[Registro] Notificando a administrador(es)...");

        NotificacionReproduccionDTO dto = new NotificacionReproduccionDTO(idAudio, fechaHoraReproduccion);

        for (CallBackInt callback : administradoresRegistrados) {
            try {
                callback.notificar(dto);
            } catch (RemoteException e) {
                System.out.println("[Registro] Error notificando a un administrador: "
                        + e.getMessage());
            }
        }
    }
}