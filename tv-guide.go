package mypkg

import (
	"flag"
	"fmt"
	"os"

	"github.com/xhenner/xmltvparse"
)

var url *string

func main() {
	flag_set := flag.NewFlagSet("tv-guide", flag.ExitOnError)

	url = flag_set.String("url", "", "(required) the url to download from")

	err := flag_set.Parse(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stdout, "error: %s\n\n", err.Error())
		flag_set.PrintDefaults()
		return
	}

	content, err := download(url)
	if err != nil {
		fmt.Fprintf(os.Stdout, "error: %s\n\n", err.Error())
		return
	}

	tvgrid := make(xmltvparse.TvGrid)

	tvgrid.Load([]byte(content))
}
