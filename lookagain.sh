# reverse the output, cut the first 3 and show the rest, reverse, cut the first 2

find -name '*.sh' -printf '%f\n'| rev | cut -c 4- | rev

