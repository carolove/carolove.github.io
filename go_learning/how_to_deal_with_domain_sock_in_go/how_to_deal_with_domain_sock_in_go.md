# how to deal with domain socket in golang

# at server side
- listen unix domain sock
- accept unix domain sock
- unix domain conn write byte
- unix domain conn close

# at client side
- unix domain sock dial as a domain socket conn
- unix domain conn read byte
- unix domain conn defer close