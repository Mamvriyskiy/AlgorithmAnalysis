#!/bin/bash

k=0
for ((i = 1; i < 2; i++)) do
    if [[ $i -lt 10 ]]; then
        a="0""$i"""
    else 
        a="$i"
    fi

    # mv "tests/$a.a" "tests/$a.txt"

    START=$(date +%s)
    # python3 main.py < tests/"$a" > test.txt
    go run *.go < tests/"$a" > test.txt
    END=$(date +%s)
    DIFF=$(($END - $START))

    echo "Time: $DIFF seconds" 

    if [ -z "$(diff -w "test.txt" "tests/$a.a")" ]; then
        echo "$a TEST PASSED"
        k=$(($k+1))
    else
        # diff "test.txt" "tests/$a.txt"
        echo "$a TEST FAILED"
    fi

done

echo "$k"