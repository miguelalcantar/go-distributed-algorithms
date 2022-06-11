package les

import "sync"

// Execution of the entire system begins with all the processes in arbitrary start states,
// and all channel empty. Then the processes, in lock-step, repeatedly perform the following
// to steps:
//	1. 	Apply message-generation function to the current state. Put these messages in the
//		appropiate channels.
//	2. 	Apply the state-transition function to the current state and the incoming messages
//		to obtain the new state. Remove all messages from the channels.

var (
	ChannelMap = sync.Map{}
	wg         sync.WaitGroup
)

func RoundSync(nodes []node) {

	for !endSignal.end {
		// single round

		// message-generation
		for _, n := range nodes {
			wg.Add(1)
			go func(n node, wg *sync.WaitGroup) {
				defer wg.Done()
				n.Msgs()
			}(n, &wg)
		}
		wg.Wait()

		// state-transition
		for _, n := range nodes {
			wg.Add(1)
			go func(n node, wg *sync.WaitGroup) {
				defer wg.Done()
				n.Trans()
			}(n, &wg)
		}
		wg.Wait()
	}
	
}
