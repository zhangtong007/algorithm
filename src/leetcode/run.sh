#!/bin/bash

file=$1

newfile=`echo $file | cut -c 9-`

mv $file $newfile
