unicd ()
{
	dir="$(uni path $@)"
	if [ -n "$dir" ]; then
		cd $dir
	fi
}
