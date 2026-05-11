package co.edu.unicauca.infoii.correo.DTOs;

import lombok.AllArgsConstructor;
import lombok.Getter;
import lombok.NoArgsConstructor;
import lombok.Setter;

@Getter
@Setter
@AllArgsConstructor
@NoArgsConstructor
public class CancionAlmacenarDTOInput {
    private int    idAudio;
    private String titulo;
    private String artista;
    private String genero;
    private String fechaHoraRegistro;
}