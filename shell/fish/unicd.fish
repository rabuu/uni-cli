function unicd
	set dir "$(uni path $argv)"
	if test -n "$dir"
		cd $dir
	end
end
