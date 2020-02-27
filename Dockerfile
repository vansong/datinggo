FROM scratch

WORKDIR /web
ADD ./conf.toml ./conf.toml
ADD ./goinspect ./goinspect

ENTRYPOINT ["./goinspect"]