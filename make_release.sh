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

for goarch in ""amd64 386""; do
  for goos in ""linux windows darwin""; do
    NAME="${PROJECT_NAME}_${RELEASE_VERSION}_${goos}_${goarch}"
    GOOS=${goos} GOARCH=${goarch} go build -o ${RDIR}/${NAME}
    echo ${NAME}
  done
done
