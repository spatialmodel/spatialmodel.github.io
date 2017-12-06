package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/ctessum/geom"
	"github.com/ctessum/geom/encoding/osm"
	"github.com/ctessum/geom/encoding/shp"
)

var osmFile, tag, value, out string
var noxrate, pm25rate, nh3rate, soxrate, vocrate float64

func init() {
	flag.StringVar(&osmFile, "osmfile", "beijing_china.osm.pbf", "path to the OpenStreetMap .pbf data")
	flag.StringVar(&tag, "tag", "", "the tag for which data should be extracted")
	flag.StringVar(&value, "value", "", "the value of the given tag for which data should be extracted")
	flag.StringVar(&out, "out", "emissions.shp", "path to the output shapefile")
	flag.Float64Var(&noxrate, "noxrate", 0, "rate of NOx emissions per unit area or length for polygon and line features,"+
		" respectively, and per point for point features")
	flag.Float64Var(&pm25rate, "pm25rate", 0, "rate of primary PM2.5 emissions per unit area or length for polygon and line features,"+
		" respectively, and per point for point features")
	flag.Float64Var(&nh3rate, "nh3rate", 0, "rate of NH3 emissions per unit area or length for polygon and line features,"+
		" respectively, and per point for point features")
	flag.Float64Var(&soxrate, "soxrate", 0, "rate of SOx emissions per unit area or length for polygon and line features,"+
		" respectively, and per point for point features")
	flag.Float64Var(&vocrate, "vocrate", 0, "rate of VOC emissions per unit area or length for polygon and line features,"+
		" respectively, and per point for point features")
	flag.Parse()
}

func main() {
	f, err := os.Open(osmFile)
	if err != nil {
		log.Fatal(err)
	}
	data, err := osm.ExtractTag(f, tag, value)
	if err != nil {
		log.Fatal(err)
	}
	features, err := data.Geom()
	if err != nil {
		log.Fatal(err)
	}
	typ, err := osm.DominantType(features)
	if err != nil {
		panic(err)
		log.Fatal(err)
	}
	switch typ {
	case osm.Point:
		writePoint(features)
	case osm.Line:
		writeLine(features)
	case osm.Poly:
		writePolygon(features)
	case osm.Collection:
		log.Fatal("geometry collections not currently supported")
	default:
		panic("invalid type")
	}

}

func writePolygon(features []*osm.GeomTags) {
	type dtype struct {
		geom.Polygon
		PM2_5, NOx, SOx, NH3, VOC float64
	}
	o, err := shp.NewEncoder(out, dtype{})
	if err != nil {
		log.Fatal(err)
	}
	for _, gt := range features {
		switch gt.Geom.(type) {
		case geom.Polygon:
			p := gt.Geom.(geom.Polygon)
			d := dtype{
				Polygon: p,
				PM2_5:   pm25rate * p.Area(),
				NOx:     noxrate * p.Area(),
				SOx:     soxrate * p.Area(),
				NH3:     nh3rate * p.Area(),
				VOC:     vocrate * p.Area(),
			}
			if err := o.Encode(d); err != nil {
				log.Fatal(err)
			}
		}
	}
	o.Close()
	writePrj()
}

func writeLine(features []*osm.GeomTags) {
	type dtype struct {
		geom.MultiLineString
		PM2_5, NOx, SOx, NH3, VOC float64
	}
	o, err := shp.NewEncoder(out, dtype{})
	if err != nil {
		log.Fatal(err)
	}
	for _, gt := range features {
		switch gt.Geom.(type) {
		case geom.LineString:
			l := gt.Geom.(geom.LineString)
			d := dtype{
				MultiLineString: geom.MultiLineString{l},
				PM2_5:           pm25rate * l.Length(),
				NOx:             noxrate * l.Length(),
				SOx:             soxrate * l.Length(),
				NH3:             nh3rate * l.Length(),
				VOC:             vocrate * l.Length(),
			}
			if err := o.Encode(d); err != nil {
				log.Fatal(err)
			}
		case geom.MultiLineString:
			l := gt.Geom.(geom.MultiLineString)
			d := dtype{
				MultiLineString: l,
				PM2_5:           pm25rate * l.Length(),
				NOx:             noxrate * l.Length(),
				SOx:             soxrate * l.Length(),
				NH3:             nh3rate * l.Length(),
				VOC:             vocrate * l.Length(),
			}
			if err := o.Encode(d); err != nil {
				log.Fatal(err)
			}
		}
	}
	o.Close()
	writePrj()
}

func writePoint(features []*osm.GeomTags) {
	type dtype struct {
		geom.Point
		PM2_5, NOx, SOx, NH3, VOC float64
	}
	o, err := shp.NewEncoder(out, dtype{})
	if err != nil {
		log.Fatal(err)
	}
	for _, gt := range features {
		switch gt.Geom.(type) {
		case geom.Point:
			p := gt.Geom.(geom.Point)
			d := dtype{
				Point: p,
				PM2_5: pm25rate,
				NOx:   noxrate,
				SOx:   soxrate,
				NH3:   nh3rate,
				VOC:   vocrate,
			}
			if err := o.Encode(d); err != nil {
				log.Fatal(err)
			}
		case geom.MultiPoint:
			mp := gt.Geom.(geom.MultiPoint)
			for _, p := range mp {
				d := dtype{
					Point: p,
					PM2_5: pm25rate,
					NOx:   noxrate,
					SOx:   soxrate,
					NH3:   nh3rate,
					VOC:   vocrate,
				}
				if err := o.Encode(d); err != nil {
					log.Fatal(err)
				}
			}
		}
	}
	o.Close()
	writePrj()
}

func writePrj() {
	fname := strings.TrimSuffix(out, filepath.Ext(out)) + ".prj"
	f, err := os.Create(fname)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(f, `GEOGCS["GCS_WGS_1984",DATUM["D_WGS_1984",SPHEROID["WGS_1984",6378137,298.257223563]],PRIMEM["Greenwich",0],UNIT["Degree",0.017453292519943295]]`)
	f.Close()
}
