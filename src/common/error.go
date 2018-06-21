package common

const(
	EOK = ErrorStatus(iota)
	ENotOk = 1
	EBadParam = 100
	EDirNotEmpty = 101
	ETargetDirExist = 102
	EFileExist = 103
	EReadOnly = 106
	ENoPermission = 200
	ETimeOut = 500
	EWriteError = 501
	EReadError = 502
	ENotEnoughSpace = 600
	EDirLocked = 804
	EDirUnlock = 805
	EDirLockCleaning = 806
	ENotFound = 807
)

type ErrorStatus int32

func (e ErrorStatus) String() string{
	return "no access"
}

