alias ls="ls --color=auto"

tmx() {
    tmux attach -t $1 || tmux new -s $1
}

function mvim() {
    vim -p `git status --porcelain | sed -ne 's/^ M //p'`
}

function grim() {
    vim +/$1 -p $(git grep -l $1)
}

# my own binaries
export PATH=$PATH:$HOME/bin

export GOPATH=$HOME/go
