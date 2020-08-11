#!/usr/bin/env bash

# docs: https://dds.cr.usgs.gov/srtm/version2_1/Documentation/

lat=$1
lon=$2

# lat prefix 00
printf -v latP1 "%02g" $lat
# lat prefix N/S
if [[ $latP1 == -* ]]; then
  lat=S${latP1#*-}
else
  lat=N${latP1}
fi

# lon prefix 000
printf -v lonP1 "%03g" $lon
# lon prefix E/W
if [[ $lonP1 == -* ]]; then
  lon=W${lonP1#*-}
else
  lon=E${lonP1}
fi

file=${lat}${lon}.hgt.zip
curl -O https://dds.cr.usgs.gov/srtm/version2_1/SRTM3/Eurasia/${file}
unzip ${file}
rm -f ${file}
