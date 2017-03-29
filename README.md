GoWorkshop Session 1
===========================

## Helps

### Channels

channels require two process at once to work. 
One who put the content through the channel, and one who grabs that from the other one.
Otherwise the process will be blocked and wait until someon come and give/receive the content through the channel other side

#### Source

* [Go Tour](https://tour.golang.org/concurrency/2)
* [Go By Example](https://gobyexample.com/channels)

### Fanin/Fanout Pipeline pattern

https://blog.golang.org/pipelines

#### FAN OUT

Multiple functions reading from the same channel until that channel is closed

#### FAN IN

A function can read from multiple inputs and proceed until all are closed by
multiplexing the input channels onto a single channel that's closed when
all the inputs are closed.

#### PATTERN

there's a pattern to our pipeline functions:
-- stages close their outbound channels when all the send operations are done.
-- stages keep receiving values from inbound channels until those channels are closed.



