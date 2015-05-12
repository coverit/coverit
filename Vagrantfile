# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure(2) do |config|
  config.vm.host_name = "coverit.dev"

  config.vm.box = "coverit/golang-dev"
  config.vm.box_version = "20150512.0.0"

  config.vm.network :forwarded_port, host: 5801, guest: 3000 # api
  config.vm.network :forwarded_port, host: 5802, guest: 27017 # mongo
  config.vm.network :forwarded_port, host: 5803, guest: 5803 # reserved

  config.vm.provider "virtualbox" do |vb|
    vb.memory = "4096"
  end

  config.vm.provision "shell", privileged: false, inline: <<-SHELL
    # Setup go workspace, see https://golang.org/doc/code.html
    mkdir -p ~/go/src/github.com/coverit
    ln -s /vagrant ~/go/src/github.com/coverit/coverit
  SHELL

  config.vm.provision "shell", inline: <<-SHELL
    # Accelerate `docker pull` with daocloud.io
    echo "DOCKER_OPTS=\\\"--registry-mirror=http://c3de84d2.m.daocloud.io \\\$DOCKER_OPTS\\\"" >>/etc/default/docker

    # Install mongodb
    apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv 7F0CEB10
    echo "deb http://repo.mongodb.org/apt/ubuntu "$(lsb_release -sc)"/mongodb-org/3.0 multiverse" | tee /etc/apt/sources.list.d/mongodb-org-3.0.list
    apt-get update
    apt-get install -y mongodb-org
  SHELL

end
