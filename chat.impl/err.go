package impl

import (
	"runtime/debug"

	"chat.logger"
)

func CatchErr() {
	if err := recover(); err != nil {
		logger.Error(err, "\n", string(debug.Stack()))
	}
}
