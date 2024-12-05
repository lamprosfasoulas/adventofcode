#!/bin/bash

file="$1"

if [ -z "$file" ]; then
    echo "Choose a file"
    exit 1
fi

total=0
matches=$(grep -oE "do\(\)|don't\(\)|mul\([0-9]{1,3},[0-9]{1,3}\)" "$file")
run="do"
for match in $matches; do
    echo "$match"
    if [[ "$match" == "do()" ]]; then
        run="do"
        continue
    elif [[ "$match" == "don't()" ]]; then
        run="dont"
        continue
    fi
    if [[ "$run" == "do" ]]; then
        num1=$(echo "$match" | cut -d',' -f1 | cut -d'(' -f2)
        num2=$(echo "$match" | cut -d',' -f2 | cut -d')' -f1)
        result=$((num1 * num2))
        total=$((total + result))
    fi
done

echo "Total is $total"
