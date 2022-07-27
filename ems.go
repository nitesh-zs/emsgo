package emsgo

/*
#cgo CFLAGS: -I.
#cgo CFLAGS: -I${SRCDIR}/include/tibems
#cgo LDFLAGS: -L${SRCDIR}/lib -ltibems -Wl,-rpath -Wl,\${SRCDIR}/lib
#include <tibems.h>
*/
import "C"
import (
	"errors"
	"unsafe"
)

type TibEMSConnection C.tibemsConnection
type TibEMSConnFactory C.tibemsConnectionFactory
type TibEMSDestination C.tibemsDestination
type TibEMSDestinationType C.tibemsDestinationType
type TibEMSSession C.tibemsSession
type TibEMSConsumer C.tibemsMsgConsumer
type TibEMSMsg C.tibemsMsg
type TibEMSMsgType C.tibemsMsgType

var errorContext C.tibemsErrorContext

func init() {
	// initialize the error context
	_ = C.tibemsErrorContext_Create(&errorContext)
}

func CreateConnFactory() TibEMSConnFactory {
	var cf TibEMSConnFactory = C.tibemsConnectionFactory_Create()
	return cf
}

func SetServerURL(factory TibEMSConnFactory, url string) error {
	status := C.tibemsConnectionFactory_SetServerURL(factory, C.CString(url))
	if status != TibEMSOk {
		return errors.New("error in setting server URL")
	}

	return nil
}

func CreateConnection(cf TibEMSConnFactory, conn TibEMSConnection, user, password string) error {
	status := C.tibemsConnectionFactory_CreateConnection(cf, &conn, C.CString(user), C.CString(password))
	if status != TibEMSOk {
		return errors.New("error in creating connection")
	}

	return nil
}

func StartConnection(conn TibEMSConnection) error {
	status := C.tibemsConnection_Start(conn)
	if status != TibEMSOk {
		return errors.New("error in starting connection")
	}

	return nil
}

func StopConnection(conn TibEMSConnection) error {
	status := C.tibemsConnection_Stop(conn)
	if status != TibEMSOk {
		return errors.New("error in stopping connection")
	}

	return nil
}

func CloseConnection(conn TibEMSConnection) error {
	status := C.tibemsConnection_Close(conn)
	if status != TibEMSOk {
		return errors.New("error in closing connection")
	}

	return nil
}

func CreateDestination(destType TibEMSDestinationType, destination string) (TibEMSDestination, error) {
	var dest C.tibemsDestination
	status := C.tibemsDestination_Create(&dest, destType, C.CString(destination))
	if status != TibEMSOk {
		return dest, errors.New("error in creating destination")
	}

	return dest, nil

}

func CreateSession(conn TibEMSConnection) (TibEMSSession, error) {
	var session TibEMSSession
	status := C.tibemsConnection_CreateSession(conn, &session, TibEMSFalse, TibEMSExplicitClientAcknowledge)
	if status != TibEMSOk {
		return session, errors.New("error in creating session")
	}

	return session, nil
}

func CreateConsumer(session TibEMSSession, dest TibEMSDestination) (TibEMSConsumer, error) {
	var msgConsumer TibEMSConsumer
	status := C.tibemsSession_CreateConsumer(session, &msgConsumer, dest, nil, TibEMSFalse)
	if status != TibEMSOk {
		return msgConsumer, errors.New("error in creating consumer")
	}

	return msgConsumer, nil
}

func DestroyMsg(msg TibEMSMsg) error {
	status := C.tibemsMsg_Destroy(msg)
	if status != TibEMSOk {
		return errors.New("error in destroying message")
	}

	return nil
}

func GetMsgType(msg TibEMSMsg) (TibEMSMsgType, error) {
	var msgType TibEMSMsgType
	status := C.tibemsMsg_GetBodyType(msg, &msgType)
	if status != TibEMSOk {
		return msgType, errors.New("error in getting message type")
	}

	return msgType, nil
}

func GetMsgText(msg TibEMSMsg) (string, error) {
	const numOfChars = 32768

	var buf *C.char = (*C.char)(C.calloc(numOfChars, 1))

	defer C.free(unsafe.Pointer(buf))

	status := C.tibemsTextMsg_GetText(msg, &buf)
	if status != TibEMSOk {
		return "", errors.New("error in getting message text")
	}

	return C.GoString(buf), nil
}

func ReceiveMsg(msgConsumer TibEMSConsumer) (TibEMSMsg, error) {
	var msg TibEMSMsg

	status := C.tibemsMsgConsumer_Receive(msgConsumer, &msg)
	if status != TibEMSOk {
		return msg, errors.New("error in receiving message")
	}

	return msg, nil
}

func DestroyDestination(dest TibEMSDestination) error {
	status := C.tibemsDestination_Destroy(dest)
	if status != TibEMSOk {
		return errors.New("error in destroying destination")
	}

	return nil
}

func CloseSession(session TibEMSSession) error {
	status := C.tibemsSession_Close(session)
	if status != TibEMSOk {
		return errors.New("error in closing session")
	}

	return nil
}

func AcknowledgeMsg(msg TibEMSMsg) error {
	status := C.tibemsMsg_Acknowledge(msg)
	if status != TibEMSOk {
		return errors.New("error in acknowledging message")
	}

	return nil
}
