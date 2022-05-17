#!/bin/bash
echo "making $1 folder and sample data"
mkdir -p $1/1
mkdir -p $1/1/1.1
mkdir -p $1/1/1.1/1.1.1
touch $1/1/1.txt
touch $1/1/1.1/1.1txt
touch $1/1/1.1/1.1.1/1.1.1txt
touch $1/1/1.1/1.1.1/1.1.2txt

mkdir -p $1/2
touch $1/2/2.txt

touch $1/$1.txt