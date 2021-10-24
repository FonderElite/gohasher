#!/bin/bash
function check_dependencies(){
        read -p "Would you like to update?:(n=0/y=1): " choice
        if [[ $choice == "0" ]];then
                echo "Skipping Update..."
                sleep 1 
        elif [[ $choice == "1" ]];then
                sudo apt update
        fi
if [ ! go ];
then
        sudo apt install golang
else
        echo "Go lang is present in: $(which go)"
        sleep 1
        echo "Skipping installation..."
        sleep 1
fi
FILE=go.mod 
if test -f "$FILE";
then
        echo "$FILE is present in $(pwd)..."
        sleep 1

else

        read -p "Project name for go mod" projname
        sudo go mod init $projname
fi
}

function main(){
figlet \#Auto Install
check_dependencies
}
main
