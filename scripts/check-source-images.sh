#!/usr/bin/env bash
# check-source-images.sh returns a non-zero status if
# we don't have any source images.

image_dir="assets/images/"
num="$(find "${image_dir}" -type f -regextype posix-extended -regex '^.*\.[jpegpn]{3,}$'|wc -l)"
if [ "$num" -lt 1 ]; then
  echo >&2 "${image_dir} directory has zero valid images"
  echo >&2 "add images before build"
  exit 1
fi
