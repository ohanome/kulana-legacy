#!/bin/bash
set -e
go build
mkdir -p usr/local/bin
cp ./kulana usr/local/bin/.
cd ..
dpkg-deb --build --root-owner-group kulana
mv kulana.deb kulana/.
cd kulana
rm -rf usr