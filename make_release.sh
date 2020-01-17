PROJECT_NAME=trainStatus
RELEASE_VERSION="$1"

if [ -z "$RELEASE_VERSION" ]
then
      RELEASE_VERSION="latest"
fi

CDIR=`pwd`
RDIR="${CDIR}/release/${RELEASE_VERSION}"

export GOPATH="${CDIR}"
export CGO_ENABLED=0

mkdir -p "$RDIR"

for GOARCH in ""amd64 386""; do
  for GOOS in ""linux windows darwin""; do
    NAME="${PROJECT_NAME}_${RELEASE_VERSION}_${GOOS}_${GOARCH}"
    go build -o ${RDIR}/${NAME}
    echo ${RDIR}/${NAME}
  done
done
