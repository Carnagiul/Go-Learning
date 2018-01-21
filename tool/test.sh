printf "%s$\n" "Updating Project"
if git pull --rebase --stat origin master
then
  printf "%s\n" "GG"
else
  printf "%s\n" 'There was an error updating. Try again later?'
fi
