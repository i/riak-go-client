package riak

type CommandImpl struct {
	Error          error
	Success        bool
	remainingTries byte
	lastNode       *Node
}

func (cmd *CommandImpl) Successful() bool {
	return cmd.Success == true
}

func (cmd *CommandImpl) onError(err error) {
	cmd.Success = false
	// NB: only set error to the *last* error (retries)
	// TODO: should we keep other errors around?
	if !cmd.hasRemainingTries() {
		cmd.Error = err
	}
}

func (cmd *CommandImpl) setRemainingTries(tries byte) {
	cmd.remainingTries = tries
}

func (cmd *CommandImpl) decrementRemainingTries() {
	cmd.remainingTries--
	logDebug("[CommandImpl]", "remainingTries: %d", cmd.remainingTries)
}

func (cmd *CommandImpl) hasRemainingTries() bool {
	return cmd.remainingTries > 0
}

func (cmd *CommandImpl) setLastNode(lastNode *Node) {
	if lastNode == nil {
		panic("[CommandImpl] nil last node")
	}
	cmd.lastNode = lastNode
}

func (cmd *CommandImpl) getLastNode() *Node {
	return cmd.lastNode
}
