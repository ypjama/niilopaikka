#!/usr/bin/env bash
# check-source-images.sh returns a non-zero status if
# we don't have any source images.

project_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")"/.. && pwd)"
image_dir="${project_dir}/assets/images/"

# Loop sub directories.
while IFS= read -r -d '' dir
do
  # Check that this sub directory has atleast one valid image file.
  num="$(find "${dir}" -type f -regextype posix-extended -regex '^.+\.(jpg|jpeg|png){1}$' | wc -l)"
  if [ "$num" -lt 1 ]; then
    echo >&2 "${dir} directory has zero valid images"
    echo >&2 "add images before build"
    exit 1
  fi

done < <(find "${image_dir}" -maxdepth 1 -mindepth 1 -type d -print0)
