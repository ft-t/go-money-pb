#!/bin/bash
cd gen
find . -type f -name 'mocks_*.go' -exec rm {} \;
FILES=( $( find . -type f -name "*.connect.go") )

for file in "${FILES[@]}"
do
  INTERFACES=( $(grep -Po "type (.*Client) interface {" "$file" | grep -Po " (.*Client)"))
    PACKAGE=( $(grep -Po "package (.*)$" "$file"))
  FILENAME=$( basename "$file")
  DIRPATH=$(dirname "$file")
  RESULT_FILENAME="${DIRPATH}/mocks_${FILENAME}"
  mockgen -destination "$RESULT_FILENAME" -package "${PACKAGE[1]}" -source="$file" -typed
done

#rm -rf `find ./ts/ -type d -name internal_api`