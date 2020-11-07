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
      vb.name = "containerd-linux"
      vb.cpus = 2
      vb.memory = 2048
  end
  config.vm.provision "shell", inline: <<-SHELL
    apt-get update
    apt-get install -y unzip gcc runc
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
  SHELL
end