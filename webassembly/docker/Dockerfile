FROM alpine:latest
MAINTAINER Sebastian Macke <sebastian@macke.de>

RUN echo "http://dl-cdn.alpinelinux.org/alpine/edge/testing" >> /etc/apk/repositories
RUN echo "http://dl-cdn.alpinelinux.org/alpine/edge/community" >> /etc/apk/repositories

RUN apk update && apk upgrade
RUN apk add --no-cache bash file curl
RUN apk add --no-cache go go-doc binaryen wabt
RUN apk add --no-cache npm
RUN apk add --no-cache nano nano-syntax tmux

# for C support 
#RUN apk add git llvm clang lld make
#RUN git clone https://github.com/CraneStation/wasi-libc.git && cd wasi-libc && make install INSTALL_DIR=/tmp/wasi-libc
#RUN mkdir -p /usr/lib/clang/10.0.0/lib/wasi && cd /usr/lib/clang/10.0.0/lib/wasi && wget "https://github.com/jedisct1/libclang_rt.builtins-wasm32.a/raw/master/precompiled/libclang_rt.builtins-wasm32.a"

# bash
RUN echo 'include "/usr/share/nano/*.nanorc"' > /etc/nanorc
RUN mkdir -p /root/wasm
RUN echo 'export PS1=" [\w] > "' >> ~/.bashrc

ENTRYPOINT ["/bin/bash"]
