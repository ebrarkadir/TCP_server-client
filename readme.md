-TCP Server
-TCP Client
-Custom Protocol
-Profile
-Optimzation
    -Cpu Profil
    -Connection Pool(sync.Pool, WorkerPool)
    -Epoll + Netpoll (non-blocking)
    -SO_REUSEPORT
    -eBPF
    -ARM



client<---------------->server
|----------------------------|
|---------stream w/r---------| 



0 1 2 3 | 4 5 6 7 | 8 N+
uint32  | uint32  | string
type    | length  | data

with pprof
duration time: 1.1420958s

without pprof
duration time: 1.070997s
