# go-broadcast [![](https://godoc.org/github.com/oliverpool/go-broadcast?status.svg)](http://godoc.org/github.com/oliverpool/go-broadcast)
Simple library to send multiple broadcast to multiple listeners without blocking

## Usage

Simply create a broadcast object:
```go
b := NewBroadcast()
```

and start sending broadcast (without having to care if someone is listening):
```go
b.Send()
```

and wait for a broadcast to come:
```go
<-b.Receive()
```

or see if a broadcast arrived:
```go
select{
case <-b.Receive():
// broadcast recevied !
case <-time.After(time.Second):
// no broadcast
}
```
