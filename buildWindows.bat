go build -v  ./cmd/stats-codegen && stats-codegen.exe
go build -v -o fesl.exe ./cmd/goheroes
go build -v ./cmd/tpl-bindata && tpl-bindata.exe
go build -v -o magma.exe ./cmd/heroes-api
start ./fesl.exe &&  start ./magma.exe