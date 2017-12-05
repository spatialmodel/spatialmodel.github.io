package main

import (
	"encoding/csv"
	"log"
	"os"
	"sort"

	"github.com/ctessum/geom/encoding/osm"
)

func main() {
	f, err := os.Open("beijing_china.osm.pbf")
	if err != nil {
		log.Fatal(err)
	}
	tags, err := osm.CountTags(f)
	o, err := os.Create("beijing_tags.csv")
	if err != nil {
		log.Fatal(err)
	}
	sort.Sort(sort.Reverse(&tags))
	w := csv.NewWriter(o)
	if err := w.WriteAll(tags.Table()); err != nil {
		log.Fatal(err)
	}
	w.Flush()
	f.Close()
}
