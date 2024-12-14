#!/bin/bash

DIR=${1:-./...}

# Set the color variable
green='\033[0;32m'
red='\033[0;31m'

# Clear the color after that
clear='\033[0m'

docker exec -i golang_socialmedia go test -v $DIR -cover | tee test_output.log

FAILED_TESTS=$(grep -E "^--- FAIL:" test_output.log)

if [ -n "$FAILED_TESTS" ]; then
    echo ""
    echo "${red}==== FAILED TESTS ====${clear}"
    echo "$FAILED_TESTS"
else
    echo ""
    printf "${green}All tests passed!${clear}"
    echo ""
fi

rm test_output.log