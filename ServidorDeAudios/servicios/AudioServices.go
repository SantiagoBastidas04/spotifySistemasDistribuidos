package servicios

import (
	. "servidor.local/grpc-servidor/modelos"
)

func CargarMetadataAudios(vec []MetadataAudio) {
	var objAudio1, objAudio2, objAudio3, objAudio4, objAudio5 MetadataAudio
	objAudio1.SetTitulo("Hasta la madre")
	objAudio1.SetTipo("Popular")
	objAudio1.SetDuracion(4)
	objAudio1.SetDisponible(false)

	objAudio2.SetTitulo("No he podido ser feliz")
	objAudio2.SetTipo("Vallenato")
	objAudio2.SetDuracion(4)
	objAudio2.SetDisponible(false)

	objAudio3.SetTitulo("El cartel paranormal de la mega")
	objAudio3.SetTipo("podcast")
	objAudio3.SetDuracion(4)
	objAudio3.SetDisponible(true)

	objAudio4.SetTitulo("AP")
	objAudio4.SetTipo("Reggaeton")
	objAudio4.SetDuracion(4)
	objAudio4.SetDisponible(false)

	objAudio5.SetTitulo("lo que hay x ahi")
	objAudio5.SetTipo("Reggaeton")
	objAudio5.SetDuracion(4)
	objAudio5.SetDisponible(true)

	vec[0] = objAudio1
	vec[1] = objAudio2
	vec[2] = objAudio3
	vec[3] = objAudio4
	vec[4] = objAudio5

}

func BuscarAudio(titulo string, vectorMetadataAudio []MetadataAudio) RespuestaMetadataAudioDTO {
	for i := 0; i < len(vectorMetadataAudio); i++ {
		if vectorMetadataAudio[i].GetTitulo() == titulo {
			var resp RespuestaMetadataAudioDTO
			resp.ObjAudio = vectorMetadataAudio[i]
			resp.Mensaje = "Encontrada"
			resp.Codigo = 200
			return resp

		}
	}
	var resp RespuestaMetadataAudioDTO
	resp.Codigo = 404
	resp.Mensaje = "No encontrada"
	return resp
}
