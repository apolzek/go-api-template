#!/bin/bash
tempfile=$(mktemp)
for i in $(seq 1 2000); do
    echo "curl -X GET http://localhost:8000/api/v1/products" >> "$tempfile"
done

cat "$tempfile" | xargs -P 100 -I {} sh -c "{}"

rm "$tempfile"
