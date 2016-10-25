/*
CHALLENGE #1
Is this fan out?

YES
Multiple functions reading from the same channel until that channel is closed

CHALLENGE #2
Is this fan in?

NO
A function can read from multiple inputs and proceed until all are closed by
multiplexing the input channels onto a single channel that's closed when
all the inputs are closed.

*/
