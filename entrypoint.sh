#!/bin/sh -l

# Run your Go program and set GITHUB_OUTPUT
program_input=$(app $ACTIONS_INPUT)

#echo "$program_input"
# Set the output variable
echo "version=$program_input" >>"$GITHUB_OUTPUT"