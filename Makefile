NULL=
gopath=$(GOPATH)

#VERSION=`git describe --tags`
#BUILD=`date +%FT%T%z`

PACKAGES=\
	"github.com/google/gousb"                \
	"github.com/sferris/howler-controller"   \
	"gopkg.in/yaml.v2"                       \
	"gopkg.in/urfave/cli.v2"                 \
	"golang.org/x/sys/unix"                  \
	$(NULL)

all: $(PACKAGES)
	GOPATH=$(gopath) go build 

$(PACKAGES): 
	GOPATH=$(gopath) go get $@
clean:
