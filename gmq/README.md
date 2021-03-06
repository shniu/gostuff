# gmq

Distribute Message Queue written in Golang.

## Design

### Namesrv

Namesrv 被作为注册中心使用，producer、broker 和 consumer 与 namesrvd 进行通信.
