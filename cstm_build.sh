# usage: sh build.sh $(go_os) $(go_arch) $(version)
# e.g. sh build.sh linux 386 3.1

ENV=${1:-darwin}
ARCH=${2:-amd64}
VERSION=${3:-1.0}

GOENV=$ENV GOARCH=$ARCH go build -o ./dist/swgserv_$ENV-$ARCH.$VERSION