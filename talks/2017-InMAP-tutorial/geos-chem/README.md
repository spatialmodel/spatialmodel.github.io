# GEOS-Chem related files

This directory is dedicated to running the InMAP preprocessor that
converts GEOS-Chem outputs into InMAP inputs.

The InMAP simulations in this tutorial are based on publicly available GEOS-Chem
data for 4x5 degree global domain for year 2013 from ftp.as.harvard.edu. Directions
for downloading the data are [here](http://wiki.seas.harvard.edu/geos-chem/index.php/Downloading_GEOS-Chem_source_code_and_data).

The specific datasets used here are `gcgrid/geos-chem/1yr_benchmarks/` and
`gcgrid/geos-chem/data/ExtData/GEOS_4x5/GEOS_FP/2013/`.

Contents of this directory:

* `inmap_inputs_geoschem4x5.ncf`: InMAP input file created by running the command
`inmap preproc --config=preprocess_config.toml` in the parent directory. This file
can be viewed and inspected with the `ncview` and `ncdump` commands.

* `link_files.go`: A Go-language script for linking GEOS-Chem output files to
files with the name format expected by the InMAP preprocessor. This is necessary
because the benchmark GEOS-Chem files we use are output as monthly averages,
which is a lower temporal resolution than what is useful to InMAP.

* `vegtype.global`: A GEOS-Chem land-use file, which the InMAP preprocessor requires.

__Note__: This dataset is included for tutorial purposed only.
Using this data for serious analysis is not recommended. It has a
number of problems, including:
* The horizontal resolution is very low,
* The data is only output as monthly averages, and
* Not all of the GEOS-Chem chemical species required by InMAP are in the output.
