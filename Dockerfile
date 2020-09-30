FROM ubuntu:20.04


RUN apt-get update -yqq && DEBIAN_FRONTEND="noninteractive" apt-get install -yqq \
    build-essential \
    emacs \
    vim \
    wget \
    curl \
    snap \
    software-properties-common \
    git \
    gdb
RUN add-apt-repository -y ppa:longsleep/golang-backports && apt update -yqq && apt install -yqq golang-1.15-go
RUN git clone -b 2020.06 --single-branch https://github.com/hugsy/gef.git
RUN echo source `pwd`/gef/gef.py >> ~/.gdbinit
RUN echo "export PATH=$PATH:/usr/lib/go-1.15/bin/" >> ~/.bashrc
WORKDIR /repo/
