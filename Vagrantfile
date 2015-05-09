# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure(2) do |config|
  config.vm.box = "ubuntu-cn/trusty64"
  config.vm.box_version = "0.2.1"

  # # Keep insecure keypair
  config.ssh.insert_key = false

  config.vm.provision "shell", privileged: false, inline: <<-SHELL

    # Install docker-compose
    sudo apt-get install -y python-pip
    sudo apt-get install -y python-dev
    sudo pip install -U docker-compose
  SHELL

  # Accelerate Docker
  config.vm.provision "shell", inline: "add-docker-registry http://c3de84d2.m.daocloud.io"

  # Install latest docker
  config.vm.provision :docker do |d|
  end
end
