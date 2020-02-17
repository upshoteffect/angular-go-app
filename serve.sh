#!/bin/sh

ng serve &

gin --port 4201 --path . --build ./src/web-api/ --i --all &

wait