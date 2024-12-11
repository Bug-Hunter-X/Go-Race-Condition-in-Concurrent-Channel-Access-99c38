# Go Race Condition Example

This repository contains a Go program that demonstrates a race condition when using channels with multiple goroutines.  The program attempts to send integers to a channel and then receive them, but without proper synchronization, this leads to unpredictable behavior.

## Bug Description
The original code lacks sufficient synchronization mechanisms to handle concurrent access to the channel `ch`.  The `wg.Wait()` call waits for all goroutines to complete sending to the channel, but it does not guarantee that the values will be received in the correct order. If the receiving loop finishes before all goroutines have sent their values, it will prematurely close the channel causing values to be lost.

## Solution
The solution introduces a buffered channel to mitigate the race condition. This allows goroutines to send their values concurrently without blocking until the receiving goroutine is ready.  The buffered channel provides a temporary storage for the values, ensuring that no values are lost. This also addresses the premature closing of the channel.

## How to Run
1. Clone this repository.
2. Navigate to the repository's directory.
3. Run the original buggy code using `go run bug.go`.
4. Run the fixed code using `go run bugSolution.go`. 
Observe the difference in output and behavior between both runs.