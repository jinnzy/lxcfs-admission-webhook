FROM ubuntu:20.04 as build
RUN echo 'deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ jammy main restricted universe multiverse' > /etc/apt/sources.list && \
    echo 'deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ jammy-updates main restricted universe multiverse' >> /etc/apt/sources.list && \
    echo 'deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ jammy-backports main restricted universe multiverse' >> /etc/apt/sources.list && \
    echo 'deb https://mirrors.tuna.tsinghua.edu.cn/ubuntu/ jammy-security main restricted universe multiverse' >> /etc/apt/sources.list && \
    touch /etc/apt/apt.conf.d/99verify-peer.conf && \
    echo >>/etc/apt/apt.conf.d/99verify-peer.conf "Acquire { https::Verify-Peer false }"

RUN apt -y update
RUN apt -y install libfuse-dev libpam-dev wget pkg-config gcc automake autoconf libtool make meson
ENV LXCFS_VERSION 5.0.0
RUN DEBIAN_FRONTEND=noninteractive apt -y install python3 python3-jinja2
RUN wget https://linuxcontainers.org/downloads/lxcfs/lxcfs-$LXCFS_VERSION.tar.gz && \
	mkdir /lxcfs && tar xzvf lxcfs-$LXCFS_VERSION.tar.gz -C /lxcfs  --strip-components=1 && \
    cd /lxcfs && \
    make


FROM ubuntu:22.04
STOPSIGNAL SIGINT

COPY --from=build /lxcfs/build/lxcfs /lxcfs/lxcfs
COPY --from=build /lxcfs/build/liblxcfs.so /lxcfs/liblxcfs.so

COPY start.sh /
CMD ["/start.sh"]