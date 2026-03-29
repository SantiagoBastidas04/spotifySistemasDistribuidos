module cliente.local/grpc-cliente

go 1.24.5

// Apunta a los directorios de cada servidor
replace servidor.local/servidorDeAudios => ../ServidorDeAudios

replace servidor.local/servidorStreaming => ../ServidorStreaming

require (
	github.com/faiface/beep v1.1.0
	google.golang.org/grpc v1.79.3
	servidor.local/servidorDeAudios v0.0.0
	servidor.local/servidorStreaming v0.0.0
)

require (
	github.com/hajimehoshi/go-mp3 v0.3.0 // indirect
	github.com/hajimehoshi/oto v0.7.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/exp v0.0.0-20190306152737-a1d7652674e8 // indirect
	golang.org/x/image v0.0.0-20190227222117-0694c2d4d067 // indirect
	golang.org/x/mobile v0.0.0-20190415191353-3e0bab5405d6 // indirect
	golang.org/x/net v0.48.0 // indirect
	golang.org/x/sys v0.39.0 // indirect
	golang.org/x/text v0.32.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20251202230838-ff82c1b0f217 // indirect
	google.golang.org/protobuf v1.36.11 // indirect
)
