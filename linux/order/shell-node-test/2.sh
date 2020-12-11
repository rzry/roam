#!/bin/zsh
set -e

for k v in $(cat 2.txt); do
	if (( $v >= 0 )); then
		echo ${k/,/}
	fi
done


strings 2.txt | while read a b; do
	if (( $b >= 0 )); then
		echo ${a/,/}
	fi
done 
