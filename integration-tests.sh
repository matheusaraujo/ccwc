#!/usr/bin/env bats

setup() {
    go build
}

run_wc_test() {
    local options="$1"
    local test_file="$2"

    local ccwc_output
    local ccwc_time
    if [ -z "$options" ]; then
        ccwc_output=$( { time ./ccwc "$test_file"; } 2>&1 )
    else
        ccwc_output=$( { time ./ccwc "$options" "$test_file"; } 2>&1 )
    fi
    ccwc_time=$(echo "$ccwc_output" | grep real | awk '{print $2}')
    ccwc_output="$(echo "$ccwc_output" | sed 1q | tr -s '[:space:]')"

    local wc_output
    local wc_time
    if [ -z "$options" ]; then
        wc_output=$( { time wc "$test_file"; } 2>&1 )
    else
        wc_output=$( { time wc "$options" "$test_file"; } 2>&1 )
    fi
    wc_time=$(echo "$wc_output" | grep real | awk '{print $2}')
    wc_output="$(echo "$wc_output" | sed 1q | tr -s '[:space:]')"

    echo "ccwc execution time: $ccwc_time"
    echo "wc execution time: $wc_time"

    wc_time=$(echo "$wc_time" | sed -r "s/0m//g" | sed -r "s/s//g")
    ccwc_time=$(echo "$ccwc_time" | sed -r "s/0m//g" | sed -r "s/s//g")
    percentage=$(echo "scale=2; ($ccwc_time / $wc_time)" | bc)
    echo "performance: $percentage"

    echo "$ccwc_output"
    echo "$wc_output"

    [ "$ccwc_output" == "$wc_output" ]
}


@test "wc with no options" {
    run_wc_test "" "test.txt"
}

@test "wc with -c option (byte count)" {
    run_wc_test "-c" "test.txt"
}

@test "wc with -w option (word count)" {
    run_wc_test "-w" "test.txt"
}

@test "wc with -l option (line count)" {
    run_wc_test "-l" "test.txt"
}

@test "wc with -m option (character count)" {
    run_wc_test "-m" "test.txt"
}
