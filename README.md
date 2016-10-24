GoMarsys Workshop Session 1
===========================

## Challanges

### Replicate wait group with channel
#### single core
#### multi core
### Fan pipeline pattern
#### IN
#### OUT
#### Is this fan out and fan in ?
### Dead lock

## Help

### FAN OUT

Multiple functions reading from the same channel until that channel is closed

### FAN IN

A function can read from multiple inputs and proceed until all are closed by
multiplexing the input channels onto a single channel that's closed when
all the inputs are closed.

### PATTERN

there's a pattern to our pipeline functions:
-- stages close their outbound channels when all the send operations are done.
-- stages keep receiving values from inbound channels until those channels are closed.

### SOURCE

https://blog.golang.org/pipelines
