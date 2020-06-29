#!/usr/bin/env bash

if [[ -z "$1" ]]; then
  echo "Usage: $0 tested_executable"
  exit 1
fi
tested_exec="$(pwd)/$1"

here="$(pwd)/$(dirname "$0")"
test_cases_file="testcases.txt"
test_cases_filepath="$here/$test_cases_file"
nbr_failures=0

n_tests="$(wc -l "$test_cases_filepath" | cut -d ' ' -f7)"
i=1
echo "Running $n_tests tests..."
echo

while read -r line; do
  regex="$(echo "$line" | cut -d " " -f1)"
  str="$(echo "$line" | cut -d " " -f2)"
  expected_outcome="$(echo "$line" | cut -d " " -f3)"
  echo "[$i/$n_tests] Matching '$str' with regex '$regex'"

  outcome="$($tested_exec "$regex" "$str")"
  if [[ "$outcome" != "$expected_outcome" ]]; then
    nbr_failures=$((nbr_failures+1))
    echo ">>> ERROR: expected $expected_outcome, returned $outcome"
  fi
  i=$((i+1))
done < "$test_cases_filepath"

echo

if [[ $nbr_failures == 0 ]]; then
  echo "All tests passed successfully"
  exit 0
else
  echo "Oh no!... $nbr_failures tests failed!!!"
  exit 1
fi
