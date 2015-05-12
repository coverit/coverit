# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure(2) do |config|
  config.vm.host_name = "coverit.dev"

  config.vm.box = "coverit/golang-dev"
  config.vm.box_version = "20150512.0.0"

  config.vm.provider "virtualbox" do |vb|
    vb.memory = "4096"
  end

  config.vm.provision "shell", privileged: false, inline: <<-SHELL
    # Setup go workspace, see https://golang.org/doc/code.html
    mkdir -p ~/go/src/github.com/coverit
    ln -s /vagrant ~/go/src/github.com/coverit/coverit
  SHELL

  # Accelerate Docker
  config.vm.provision "shell", inline: <<-SHELL
    echo "DOCKER_OPTS=\\\"--registry-mirror=http://c3de84d2.m.daocloud.io \\\$DOCKER_OPTS\\\"" >>/etc/default/docker
  SHELL

end
