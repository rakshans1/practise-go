#!/bin/sh

set -e

if tmux has-session -t practise-go 2> /dev/null; then
  tmux attach -t practise-go
  exit
fi

tmux new-session -d -s practise-go -n editor
tmux send-keys -t practise-go:editor "v " Enter
tmux split-window -t practise-go:editor -h
tmux resize-pane -t practise-go:1.2 -R 50
tmux attach -t practise-go:editor.top
