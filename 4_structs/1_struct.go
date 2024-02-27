package main

type msgToSend struct {
	msg       string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func canSendMsg(mToSend, msgToSend) bool {
	if mToSend.sender.name == "" {
		return false
	}
	if mToSend.recipient.name == "" {
		return false
	}
	if mToSend.sender.number == "" {
		return false
	}
	if mToSend.recipient.number == "" {
		return false
	}
	return true
}

func test(mToSend, msgToSend)
