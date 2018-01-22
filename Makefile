# **************************************************************************** #
#                                                                              #
#                                                         :::      ::::::::    #
#    Makefile                                           :+:      :+:    :+:    #
#                                                     +:+ +:+         +:+      #
#    By: piquerue <marvin@42.fr>                    +#+  +:+       +#+         #
#                                                 +#+#+#+#+#+   +#+            #
#    Created: 2016/11/02 14:20:07 by piquerue          #+#    #+#              #
#    Updated: 2017/10/10 18:18:02 by piquerue         ###   ########.fr        #
#                                                                              #
# **************************************************************************** #

NAME = wolf3d

INCLUDE = $(shell find . -type f | grep "\.go")

INCLUDE_2 = *.go

all: $(NAME)


$(NAME):
	sh tool/display.sh
	@go build -o a.out $(INCLUDE_2)

.PHONY : all
