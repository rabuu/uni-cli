# disable file completions
complete -c unicd -f

# course completion
complete -c unicd -a '(uni course list --fish)'

# flag completion
complete -c unicd -l export -s x -d 'Go to the export directory'
complete -c unicd -l material -s m -d 'Go to the material directory'
