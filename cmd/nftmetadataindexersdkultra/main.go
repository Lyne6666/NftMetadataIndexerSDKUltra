// cmd/nftmetadataindexersdkultra/main.go
package main

import (
	"flag"
	"log"
	"os"
	
	"nftmetadataindexersdkultra/internal/nftmetadataindexersdkultra"
)

func main() {
	verbose := flag.Bool("verbose", false, "Enable verbose logging")
	flag.Parse()
	
	app := nftmetadataindexersdkultra.NewApp(*verbose)
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
