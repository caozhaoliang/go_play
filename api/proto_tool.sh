#/bin/bash
protoc --go_out=. --proto_path=. -I=. --proto_path=$(dirname $(dirname $(dirname "$PWD"))) *.proto
ls *.pb.go | xargs -n1 -IX bash -c 'sed s/,omitempty// X > X.tmp && mv X{.tmp,}'
