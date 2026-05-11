package co.edu.unicauca.capaDeControladores;

import java.rmi.RemoteException;
import java.rmi.server.UnicastRemoteObject;
import co.edu.unicauca.fachadaServices.DTO.NotificacionReproduccionDTO;

public class CallBackImp extends UnicastRemoteObject implements CallBackInt {

    public CallBackImp() throws RemoteException {
        super();
    }

    @Override
    public void notificar(NotificacionReproduccionDTO notificacion) throws RemoteException {
        mostrarNotificacion(notificacion);
    }

    private void mostrarNotificacion(NotificacionReproduccionDTO notificacion) {
        System.out.println("\n===== NOTIFICACIÓN DE REPRODUCCIÓN =====");
        System.out.println("ID del audio  : " + notificacion.getIdAudio());
        System.out.println("Fecha y hora  : " + notificacion.getFechaHoraReproduccion());
        System.out.println("=========================================\n");
    }
}