package modelos

import "encoding/json"

type MetadataCanciones struct {
	titulo     string
	duracion   int
	tipo       string
	disponible bool
}

//setters and getters

func (m *MetadataCanciones) GetTitulo() string {
	return m.titulo
}

func (m *MetadataCanciones) SetTitulo(titulo string) {
	m.titulo = titulo
}

func (m *MetadataCanciones) GetDuracion() int {
	return m.duracion
}

func (m *MetadataCanciones) SetDuracion(duracion int) {
	m.duracion = duracion
}

func (m *MetadataCanciones) GetTipo() string {
	return m.tipo
}

func (m *MetadataCanciones) SetTipo(tipo string) {
	m.tipo = tipo
}

func (m *MetadataCanciones) GetDisponible() bool {
	return m.disponible
}

func (m *MetadataCanciones) SetDisponible(disponible bool) {
	m.disponible = disponible
}

// codificar a json
func (m MetadataCanciones) MarshalJSON() ([]byte, error) {
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
func (m *MetadataCanciones) UnmarshalJSON(data []byte) error {
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
