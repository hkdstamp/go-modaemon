#!/bin/bash

yum install -y git gcc

archive=go1.8.linux-amd64.tar.gz
if [ ! -f $archive ]; then
  wget -q https://storage.googleapis.com/golang/$archive
fi

if [ ! -d /home/ec2-user/go ]; then
  tar -C /home/ec2-user -xzf $archive
  chown -R ec2-user.ec2-user /home/ec2-user/go
fi

if ! grep GOPATH /home/ec2-user/.bashrc; then
  echo 'export GOROOT=/home/ec2-user/go' >> /home/ec2-user/.bashrc
  echo 'export PATH=$PATH:$GOROOT/bin:$HOME/bin' >> /home/ec2-user/.bashrc
  echo 'export GOPATH=$HOME' >> /home/ec2-user/.bashrc
  echo 'sudo chown -R ec2-user.ec2-user $HOME/src' >> /home/ec2-user/.bashrc
fi
