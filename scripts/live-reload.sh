#!/bin/bash
# Start live reload of our app with the air tool.

project_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")"/.. && pwd)"

# We need go.
if ! type "go" >/dev/null 2>&1; then
  echo >&2 "go tool is not available, cannot continue"
	exit 1
fi

# Install air if it is not available.
if ! type "air" >/dev/null 2>&1; then
  cd
  go get -u github.com/cosmtrek/air
fi

cd "${project_dir}" || exit 1
air
