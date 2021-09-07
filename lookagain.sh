# Create a file lookagain.sh, which will look, 
#from the current directory and its sub-folders for:
#all the files ending with .sh.

#That command will only show the name of the files without the .sh. 
#The files will be in descending order (as shown in the below example).

find . -name '*.sh' -prune -o -print