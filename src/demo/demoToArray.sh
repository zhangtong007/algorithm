#!/bin/bash

arr=$1

echo $arr |sed -e 's/\]/\}/g' |sed -e 's/\[/\{/g'