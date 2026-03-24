package modelos

import "encoding/json"

type MetadataAudio struct {
	titulo     string
	duracion   int
	tipo       string
	disponible bool
}

//setters and getters

func (m *MetadataAudio) GetTitulo() string {
	return m.titulo
}

func (m *MetadataAudio) SetTitulo(titulo string) {
	m.titulo = titulo
}

func (m *MetadataAudio) GetDuracion() int {
	return m.duracion
}

func (m *MetadataAudio) SetDuracion(duracion int) {
	m.duracion = duracion
}

func (m *MetadataAudio) GetTipo() string {
	return m.tipo
}

func (m *MetadataAudio) SetTipo(tipo string) {
	m.tipo = tipo
}

func (m *MetadataAudio) GetDisponible() bool {
	return m.disponible
}

func (m *MetadataAudio) SetDisponible(disponible bool) {
	m.disponible = disponible
}

// codificar a json
func (m MetadataAudio) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Titulo     string `json:"titulo"`
		Duracion   int    `json:"duracion"`
		Tipo       string `json:"tipo"`
		Disponible bool   `json:"disponible"`
	}{
		Titulo:     m.titulo,
		Duracion:   m.duracion,
		Tipo:       m.tipo,
		Disponible: m.disponible,
	})
}

// descodificar de json
func (m *MetadataAudio) UnmarshalJSON(data []byte) error {
	aux := struct {
		Titulo     string `json:"titulo"`
		Duracion   int    `json:"duracion"`
		Tipo       string `json:"tipo"`
		Disponible bool   `json:"disponible"`
	}{}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	m.titulo = aux.Titulo
	m.duracion = aux.Duracion
	m.tipo = aux.Tipo
	m.disponible = aux.Disponible
	return nil
}
