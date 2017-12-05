# Example emissions scenario creation

This directory contains scripts and data for creating and example scenario
of changes in emissions in Beijing, China.

### Contents

* __beijing_china.osm.pbf__: OpenStreetMap data for Beijing, China, downloaded
from https://mapzen.com/data/metro-extracts/metro/beijing_china/. The methods
here would also work for other areas such as all of China, but this example
is limited to Beijing to limit the file size.

* __get_tags.go__: A script written in the [Go](https://golang.org/) language for extracting the
available location tags in the OpenStreetMap data.

* __beijing_tags.csv__: The available tags and the number of occurences of each
one, created by running the above script.

* __process_emissions.go__: A script written in the [Go](https://golang.org/) language for extracting
OpenStreetMap features matching a given tag and value and assigning emissions to them.

* __process_emissions_mac__, __process_emissions_linux__, __process_emissions_windows.exe__:
Executable programs created by compiling the `process_emissions.go` script. They can
be run like
```
process_emissions_xxx -osmfile=beijing_china.osm.pbf -tag=highway -value=traffic_signals -rate=10 -out=../run/emissions.shp
```
where `process_emissions_xxx` is the appropriate executable for your system,
`osmfile` is the path to the OpenStreetMap data, `tag` is the tag key of interest,
`value` is the desired value of the tag of interest (tags are in a `tag=value` format),
rate is the rate of emissions per point or per unit length or area of a feature
(so the units of rate are [mass/(- or length or area)/time], where the mass
and time units will be specified in the InMAP configuration),
and out is the desired path to the output shapefile.


* __build.sh__: script for creating the executables from `process_emissions.go`
