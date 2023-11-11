#############################################################
###################### Install Golang #######################
wget https://dl.google.com/go/go1.20.3.linux-amd64.tar.gz
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -zxvf go1.20.3.linux-amd64.tar.gz
mkdir -p ~/go/{bin,pkg,src}

# The following assume that your shell is bash
echo 'export GOPATH=$HOME/go' >> ~/.bashrc
echo 'export GOROOT=/usr/local/go' >> ~/.bashrc
echo 'export PATH=$PATH:$GOPATH/bin:$GOROOT/bin' >> ~/.bashrc
echo 'export GO111MODULE=auto' >> ~/.bashrc
source ~/.bashrc
rm go1.20.3.linux-amd64.tar.gz
#############################################################

