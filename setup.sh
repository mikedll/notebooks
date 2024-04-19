#!/usr/bin/env bash

if [ ! -f cs ]; then
	echo "Fetching coursier"
	curl -fL "https://github.com/coursier/launchers/raw/master/cs-x86_64-pc-linux.gz" | gzip -d > cs
	chmod u+x ./cs
fi

if [ ! -d target/deps ]; then
	mkdir target/deps
fi

if [ ! -f target/deps/os-lib_3-0.10.0.jar ]; then
	wget https://repo1.maven.org/maven2/com/lihaoyi/os-lib_3/0.10.0/os-lib_3-0.10.0.jar -O target/deps/os-lib_3-0.10.0.jar
fi

if [ ! -f target/deps/ujson_3-3.3.0.jar ]; then
  wget https://repo1.maven.org/maven2/com/lihaoyi/ujson_3/3.3.0/ujson_3-3.3.0.jar -O target/deps/ujson_3-3.3.0.jar
fi

echo "setup.sh Done"