#!/bin/bash
SERVER=10.72.46.240
PORT=1323
case $1 in

    list)
    curl $SERVER:$PORT/list
    ;;

    install)
    curl $SERVER:$PORT/install/$2
    ;;

    delete)
    curl $SERVER:$PORT/delete/$2
    ;;

esac