package main

import (
	"log"
	"os"

	"github.com/halfenif/torrent-client/torrentfile"
)

func main() {

	if len(os.Args) != 3 {
		println("Usage: go run main.go torrent_file.torrent output_file")
		panic("실행 아규먼트 갯수를 확인하시기 바랍니다.")
	}

	inPath := os.Args[1]
	outPath := os.Args[2]

	tf, err := torrentfile.Open(inPath)
	if err != nil {
		log.Fatal(err)
	}

	err = tf.DownloadToFile(outPath)
	if err != nil {
		log.Fatal(err)
	}
}
