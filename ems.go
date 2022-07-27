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

func CreateConnFactory() TibEMSConnFactory {
	cf := C.tibemsConnectionFactory_Create()
	return TibEMSConnFactory(cf)
}

func SetServerURL(factory TibEMSConnFactory, url string) error {
	status := C.tibemsConnectionFactory_SetServerURL(C.tibemsConnectionFactory(factory), C.CString(url))
	if status != TibEMSOk {
		return errors.New("error in setting server URL")
	}

	return nil
}

func CreateConnection(cf TibEMSConnFactory, user, password string) (TibEMSConnection, error) {
	var conn C.tibemsConnection
	status := C.tibemsConnectionFactory_CreateConnection(C.tibemsConnectionFactory(cf), &conn, C.CString(user), C.CString(password))
	if status != TibEMSOk {
		return TibEMSConnection(conn), errors.New("error in creating connection")
	}

	return TibEMSConnection(conn), nil
}

func StartConnection(conn TibEMSConnection) error {
	status := C.tibemsConnection_Start(C.tibemsConnection(conn))
	if status != TibEMSOk {
		return errors.New("error in starting connection")
	}

	return nil
}

func StopConnection(conn TibEMSConnection) error {
	status := C.tibemsConnection_Stop(C.tibemsConnection(conn))
	if status != TibEMSOk {
		return errors.New("error in stopping connection")
	}

	return nil
}

func CloseConnection(conn TibEMSConnection) error {
	status := C.tibemsConnection_Close(C.tibemsConnection(conn))
	if status != TibEMSOk {
		return errors.New("error in closing connection")
	}

	return nil
}

func CreateDestination(destType TibEMSDestinationType, destination string) (TibEMSDestination, error) {
	var dest C.tibemsDestination
	status := C.tibemsDestination_Create(&dest, C.tibemsDestinationType(destType), C.CString(destination))
	if status != TibEMSOk {
		return TibEMSDestination(dest), errors.New("error in creating destination")
	}

	return TibEMSDestination(dest), nil

}

func CreateSession(conn TibEMSConnection) (TibEMSSession, error) {
	var session C.tibemsSession
	status := C.tibemsConnection_CreateSession(C.tibemsConnection(conn), &session, TibEMSFalse, TibEMSExplicitClientAcknowledge)
	if status != TibEMSOk {
		return TibEMSSession(session), errors.New("error in creating session")
	}

	return TibEMSSession(session), nil
}

func CreateConsumer(session TibEMSSession, dest TibEMSDestination) (TibEMSConsumer, error) {
	var msgConsumer C.tibemsMsgConsumer
	status := C.tibemsSession_CreateConsumer(C.tibemsSession(session), &msgConsumer, C.tibemsDestination(dest), nil, TibEMSFalse)
	if status != TibEMSOk {
		return TibEMSConsumer(msgConsumer), errors.New("error in creating consumer")
	}

	return TibEMSConsumer(msgConsumer), nil
}

func DestroyMsg(msg TibEMSMsg) error {
	status := C.tibemsMsg_Destroy(C.tibemsMsg(msg))
	if status != TibEMSOk {
		return errors.New("error in destroying message")
	}

	return nil
}

func getMsgType(msg TibEMSMsg) (TibEMSMsgType, error) {
	var msgType C.tibemsMsgType
	status := C.tibemsMsg_GetBodyType(msg, &msgType)
	if status != TibEMSOk {
		return TibEMSMsgType(msgType), errors.New("error in getting message type")
	}

	return TibEMSMsgType(msgType), nil
}

func GetMsgText(msg TibEMSMsg) (string, error) {
	msgType, err := getMsgType(msg)
	if err != nil {
		return "", err
	}

	if C.tibemsMsgType(msgType) != TibEMSTextMessage {
		return "", errors.New("unsupported message type")
	}

	const numOfChars = 32768

	var buf *C.char = (*C.char)(C.calloc(numOfChars, 1))

	defer C.free(unsafe.Pointer(buf))

	status := C.tibemsTextMsg_GetText(C.tibemsMsg(msg), &buf)
	if status != TibEMSOk {
		return "", errors.New("error in getting message text")
	}

	return C.GoString(buf), nil
}

func ReceiveMsg(msgConsumer TibEMSConsumer) (TibEMSMsg, error) {
	var msg C.tibemsMsg

	status := C.tibemsMsgConsumer_Receive(C.tibemsMsgConsumer(msgConsumer), &msg)
	if status != TibEMSOk {
		return TibEMSMsg(msg), errors.New("error in receiving message")
	}

	// return special message in case nil message is received
	if msg == nil {
		return "", errors.New("nil msg received")
	}

	return TibEMSMsg(msg), nil
}

func DestroyDestination(dest TibEMSDestination) error {
	status := C.tibemsDestination_Destroy(C.tibemsDestination(dest))
	if status != TibEMSOk {
		return errors.New("error in destroying destination")
	}

	return nil
}

func CloseSession(session TibEMSSession) error {
	status := C.tibemsSession_Close(C.tibemsSession(session))
	if status != TibEMSOk {
		return errors.New("error in closing session")
	}

	return nil
}

func AcknowledgeMsg(msg TibEMSMsg) error {
	status := C.tibemsMsg_Acknowledge(C.tibemsMsg(msg))
	if status != TibEMSOk {
		return errors.New("error in acknowledging message")
	}

	return nil
}
