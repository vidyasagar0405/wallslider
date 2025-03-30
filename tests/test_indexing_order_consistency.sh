#!/usr/bin/env bash

# Does the indexing order change ??
# The answer is No
# why is the script saved ??
# FOR YOUR REFERENCE!

WALLPAPER_DIR="$HOME/Pictures/wallpapers/"

for i in {1..3}; do
    ./wallslider test "$WALLPAPER_DIR" > /dev/null
    mv ~/.config/wallslider/index.json "index${i}.json"
    # Extract the 'index_arr' field using jq and write it to a new file
    jq '.index_arr' "index${i}.json" > "index${i}_c.json"
    rm "index${i}.json"
done

for i in {1..2}; do
    if diff "index${i}_c.json" "index$((i+1))_c.json" > /dev/null; then
        echo "OK"
    else
        echo "Failed"
    fi
done

rm index*
