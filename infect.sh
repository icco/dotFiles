#!/bin/bash

dotfiles=`ls`;
excludes=( "README.md" "infect.sh" );

function exclude() {
	if [ -z "$1" ]; then
		return 1;
	fi;

	for i in ${excludes[*]}; do
		if [ $i == $1 ]; then
			return 1;
		fi;
	done;

	return 0
}

for file in $dotfiles; do
	exclude $file;
	ex="$?";
	if [ $ex == 0 ]; then
		destfile=`echo ~/.$file`;
		if [ -d $file ]; then
			echo "$file/ : `ls $file`"; # needs to be recursive...?
		else
			echo "$file => $destfile"; #ln -s
		fi;
	fi;
done;

#echo $dotfiles;
#echo ${excludes[*]};
