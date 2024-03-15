#!/bin/sh -l

# supplying the args from the actions.yml 
program_input=$(app $1)

# Set the output variable
echo "version=$program_input" >>"$GITHUB_OUTPUT"