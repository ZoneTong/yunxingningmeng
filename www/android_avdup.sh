#!/bin/bash
source ~/.bashrc
name=`emulator -list-avds|head -1`
emulator -avd $name &