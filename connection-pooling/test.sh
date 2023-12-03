#!/bin/bash

# URL of your Go server
URL="http://localhost:9082/query"

# Maximum number of connections to test
MAX_CONNECTIONS=20

# Temporary files for success and failure counts
SUCCESS_FILE=$(mktemp)
FAIL_FILE=$(mktemp)

# Initialize counters to 0
echo 0 > "$SUCCESS_FILE"
echo 0 > "$FAIL_FILE"

# Function to make a request and count successes and failures
make_request() {
    if curl -s "$URL" > /dev/null; then
        echo $(($(cat "$SUCCESS_FILE") + 1)) > "$SUCCESS_FILE"
    else
        echo $(($(cat "$FAIL_FILE") + 1)) > "$FAIL_FILE"
    fi
}

# Make requests in parallel
for i in $(seq 1 $MAX_CONNECTIONS); do
    make_request &
done

# Wait for all background processes to finish
wait

# Read final counts
SUCCESS=$(cat "$SUCCESS_FILE")
FAIL=$(cat "$FAIL_FILE")

# Print results
echo "Successful connections: $SUCCESS"
echo "Failed connections: $FAIL"

# Clean up
rm "$SUCCESS_FILE" "$FAIL_FILE"
