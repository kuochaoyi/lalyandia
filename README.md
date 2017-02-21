# Lalyandia

![](http://fotos.subefotos.com/533ce797e9a7d1d27bf2bd30cc9c959bo.png)

The emulator of Lineage 2. I'm sure everyone know that Java emulators have some problems with performance, bugs and etc...
Basically I planning to use datapack from PTS.

Emulator will have in base the next components:
- Actors (With pool of actors, cluster, state machine), all packages from ProtoActor.
- Postgres as basic database
- Tile38 for geoposition features
- vndr for vendor way

## Install

- Protobuf binary by using link `https://github.com/google/protobuf/releases`
- `go get` on this repository
- and `make` for build binaries
