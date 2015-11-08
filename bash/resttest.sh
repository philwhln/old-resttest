#!/usr/bin/env bash

# store values in temp file
values=`mktemp`

# sequence through pages, cap at 999
for i in `seq 1 999`; do

  # fetch json for page
  json=`curl -sf http://resttest.bench.co/transactions/$i.json`

  # abort loop if 404 returned
  if [ "$?" -ne "0" ]; then
    break
  fi

  # parse values out of json
  echo $json | jq '.transactions[].Amount|tonumber' >> $values

done

# sum values stored in temp file, output value
paste -s'd+' $values | bc

