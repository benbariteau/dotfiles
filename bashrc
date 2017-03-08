if [ -f ~/.bashrc.local ]; then
    source ~/.bashrc.local
fi

alias ls="ls --color=auto"

alias tmx="tmux attach -t"

function mvim() {
    vim -p `git status --porcelain | sed -ne 's/^ M //p'`
}

function grim() {
    vim +/$1 -p $(git grep -l $1)
}

# my own binaries
export PATH=$PATH:$HOME/bin

# add go binaries to path
export PATH=$PATH:$GOPATH/bin
