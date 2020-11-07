# docker-rmi

Have too many docker images on your local docker daemon, and not sure which ones to delete?

Remove docker images with a prompt: Press **y** to delete ✅, and **n** to skip deletion ❌.

## Key Features

- Delete docker images with confirmation. Press **y** to delete ✅, and **n** to skip ❌.
- Images (especially the big ones) gets deleted in the background, so you don't wait, you just move onto the next selection.

## Screencast
[![asciicast](https://asciinema.org/a/371122.svg)](https://asciinema.org/a/371122)

## Wanna try it out!?

```
User: Hey! I don't feel comfortable trying it out on my laptop.
User: I fear docker-rmi will delete some images accidentally.
Me: No worries! just run "vagrant up" and it will spin up an ubuntu 18.04 VM
Me: ..and setup ~20 docker images for you to delete and try it out!
Me: Once the VM is up, just do "vagrant ssh docker-rmi" and "docker-rmi"
Me: Vagrant setup takes ~10 mins, go get a coffee :) while we set things up for you!
User: This is awesome!
```

## Installation

```
$ git clone git@github.com:shishir-a412ed/docker-rmi.git
$ cd docker-rmi
$ sudo make install (You don't need sudo on macOS)
```

## Verify

```
$ which docker-rmi
$ docker-rmi (This should start the application)
```
**NOTE**: `docker-rmi` switches your terminal to `raw` mode. This is needed to display ✅ and ❌.<br/>
**Ctrl+C** won't work if you would like to exit in between.<br/>
If you want to exit before `docker-rmi` completes, just close the tab or terminal.

## Cleanup
```
make clean
```
This will delete your binary: `docker-rmi`

```
vagrant destroy
```
This will destroy your vagrant VM.

## Currently supported environments

- macOS Catalina (Version 10.15.5)
- Ubuntu (>=16.04)
- Centos
- Fedora
