@echo off
nodemon --signal SIGKILL --ignore ui --ext go --exec "go run . || exit 1"
