#!/bin/bash

# Check if arguments were provided
if [ $# -eq 0 ]; then
    echo "Usage: $0 <command>"
    echo "Example: $0 'ls -la'"
    exit 1
fi

# Configuration
RUNS=100
TOTAL=0
COMMAND="$*"

echo "Running '$COMMAND' $RUNS times..."
echo "-------------------"

# Run the command multiple times
for i in $(seq 1 $RUNS); do
    START=$(date +%s.%N)
    eval $COMMAND > /dev/null 2>&1
    END=$(date +%s.%N)
    
    # Calculate time taken
    TIME=$(echo "$END - $START" | bc)
    TOTAL=$(echo "$TOTAL + $TIME" | bc)
    
    printf "Run %2d: %.4f seconds\n" $i $TIME
done

# Calculate and display average
AVERAGE=$(echo "scale=4; $TOTAL / $RUNS" | bc)
echo "-------------------"
echo "Average time: $AVERAGE seconds"