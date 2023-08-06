#!/bin/bash

echo "Work with default parameters? (Y/n)"
read useDefaults
useDefaults=${useDefaults:-y}

if [[ $useDefaults == "n" || $useDefaults == "no" ]]; then
  echo "Write the inputter to use: [csv]"
  read inputter

  echo "Write the outputter to use: [json]"
  read outputter

fi

inputter=${inputter:-csv}
outputter=${outputter:-json}

echo "Write the client to request: "
read client
go run cmd/cli/main.go -inputter $inputter -outputter $outputter $client
