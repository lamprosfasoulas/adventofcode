#!/bin/bash

pattern="$1"
file="tinyinput.txt"
file2="transposed.txt"
matches=$(grep -oE "XMAS" "$file")
matches2=$(grep -oE "SAMX" "$file")
matches3=$(grep -oE "XMAS" "$file2")
matches4=$(grep -oE "SAMX" "$file2")
total=0
for match in $matches; do
    total=$((total + 1))
done
for match in $matches2; do
    total=$((total + 1))
done
for match in $matches3; do
    total=$((total + 1))
done
for match in $matches4; do
    total=$((total + 1))
done
echo "Total is $total"
