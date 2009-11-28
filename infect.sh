#!/bin/bash
# A script to covert a linux profile to have all of my prefered settings.

dotfiles=`ls`;
excludes=( "README.md" "infect.sh" );

# Checks to see if a file should be excluded
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

# actually move files 
function doitnow() {
	if [ -f $2 ]; then
		mkdir -p ~/tmp/oldDotFiles/;
		mv $2 ~/tmp/oldDotFiles/;
	fi;

	ln -s $1 $2
	echo "$1 => $2";
}

for file in $dotfiles; do
	exclude $file;
	ex="$?";
	if [ $ex == 0 ]; then
		destfile=`echo ~/.$file`;
		file="`pwd`/$file";

		if [ -d $file ]; then
			if [ -d $destfile ]; then
				echo "$destfile exists"; # Push subfiles onto array?
			else
				doitnow $file $destfile;
			fi;
		else
			doitnow $file $destfile;
		fi;
	fi;
done;

#echo $dotfiles;
#echo ${excludes[*]};
