#!/bin/sh

# Init tmux session

#tmux new -s auth -d 'make auth-service'
#tmux new -s gateway -d 'make gateway-service'
#tmux new -s hotel -d 'make hotel-service'
#tmux new -s payment -d 'make payment-service'
#tmux new -s loyalty -d 'make loyalty-service'
#tmux new -s reservation -d 'make reservation-service'
#
#tmux

make auth-service&
make gateway-service&
make hotel-service&
make payment-service&
make loyalty-service&
make reservation-service&
make stat-service&

wait
