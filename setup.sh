#!/usr/bin/env bash

curl -fL "https://github.com/coursier/launchers/raw/master/cs-x86_64-pc-linux.gz" | gzip -d > cs

chmod u+x ./cs
