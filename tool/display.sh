NAME="Go-Learning"
PROJECT=("™ 乂❤‿❤乂 Project ${NAME} 乂❤‿❤乂 ™" "✦ Project ${NAME} ╭∩╮（︶︿︶）╭∩╮ ✦" "♥ (╥﹏╥) Project ${NAME} (╥﹏╥) ♥" "☠ Project ${NAME} ☠" "ツ Project ${NAME} ツ")

rand=$[ $RANDOM % ${#PROJECT[@]} ]
DATA=${PROJECT[$rand]}
clear

if which tput >/dev/null 2>&1; then
    ncolors=$(tput colors)
fi
if [ -t 1 ] && [ -n "$ncolors" ] && [ "$ncolors" -ge 8 ]; then
  RED="$(tput setaf 1)"
  GREEN="$(tput setaf 2)"
  YELLOW="$(tput setaf 3)"
  BLUE="$(tput setaf 4)"
  BOLD="$(tput bold)"
  NORMAL="$(tput sgr0)"
else
  RED=""
  GREEN=""
  YELLOW=""
  BLUE=""
  BOLD=""
  NORMAL=""
fi

echo	$GREEN
echo	${PROJECT[$rand]}
echo	$BLUE
printf '\n%s\n' '        :::      ::::::::   '
printf '%s\n' '      :+:      :+:    :+:   '
printf '%s\n' '    +:+ +:+         +:+     '
printf '%s\n' '  +#+  +:+       +#+        '
printf '%s\n' '+#+#+#+#+#+   +#+           '
printf '%s\n' '     #+#    #+#             '
printf '%s\n\n' '    ###   ########.fr       '
echo	$GREEN
echo	${PROJECT[$rand]}
echo	$NORMAL
