# Specify minimum Vagrant version and Vagrant API version
Vagrant.require_version ">= 1.6.0"
VAGRANTFILE_API_VERSION = "2"

# Create box
Vagrant.configure("2") do |config|
  config.vm.define "docker-rmi"
  config.vm.box = "hashicorp/bionic64"
  config.vm.synced_folder ".", "/home/vagrant/go/src/github.com/docker-rmi"
  config.ssh.extra_args = ["-t", "cd /home/vagrant/go/src/github.com/docker-rmi; bash --login"]
  config.vm.provider "virtualbox" do |vb|
      vb.name = "docker-rmi"
      vb.cpus = 2
      vb.memory = 2048
  end
  config.vm.provision "shell", inline: <<-SHELL
    echo "export GOPATH=/home/vagrant/go" >> /home/vagrant/.bashrc
    echo "export PATH=$PATH:/usr/local/go/bin" >> /home/vagrant/.bashrc
    source /home/vagrant/.bashrc

    # Install golang-1.14.3
    if [ ! -f "/usr/local/go/bin/go" ]; then
      curl -s -L -o go1.14.3.linux-amd64.tar.gz https://dl.google.com/go/go1.14.3.linux-amd64.tar.gz
      sudo tar -C /usr/local -xzf go1.14.3.linux-amd64.tar.gz
      sudo chmod +x /usr/local/go
      rm -f go1.14.3.linux-amd64.tar.gz
    fi

    # Install docker-rmi
    cd /home/vagrant/go/src/github.com/docker-rmi
    sudo make install
  SHELL
  config.vm.provision "docker" do |d|
    d.pull_images "traefik"
    d.pull_images "traefik:v2.3.2"
    d.pull_images "traefik:v2.3"
    d.pull_images "postgres"
    d.pull_images "postgres:alpine"
    d.pull_images "postgres:9.6.19-alpine"
    d.pull_images "ubuntu"
    d.pull_images "ubuntu:16.04"
    d.pull_images "ubuntu:18.04"
    d.pull_images "redis"
    d.pull_images "redis:buster"
    d.pull_images "redis:alpine3.12"
    d.pull_images "node"
    d.pull_images "node:stretch-slim"
    d.pull_images "node:slim"
    d.pull_images "mysql"
    d.pull_images "mysql:8.0.22"
    d.pull_images "mysql:8.0"
    d.pull_images "mariadb"
    d.pull_images "mariadb:focal"
    d.pull_images "mariadb:10.5.7"
  end
end
