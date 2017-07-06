#!/bin/sh
PKGS="xmonad stow rofi feh compton rxvt-unicode-256color i3lock "
DIRS="bin xmonad compton fonts x"
sudo apt-get install $PKGS
cd ~/xmonad && stow -S -n $DIRS
cd ~/bin/src && make
echo "Sign up for an unsplash api key to set new backgrounds on login"
echo "https://unsplash.com/join"
echo "https://unsplash.com/oauth/applications/new"
echo -n "unsplash key (Application ID): "
read UKEY

echo 'alias lockscr="golock -s 23 && i3lock -u -i ~/.lock.png"' >> ~/.bashrc
echo 'export PATH="$PATH:$HOME/bin"' >> ~/.bashrc
echo "unsplash_key : $UKEY" >> ~/.env.yml
echo 'bg_location: /usr/local/google/home/sworne/.bg' >> ~/.env.yml
echo 'bg_arg: --bg-fill' >> ~/.env.yml
echo 'bg_arg1: --no-xinerama' >> ~/.env.yml
echo 'bg_cmd: feh' ~/.env.yml
