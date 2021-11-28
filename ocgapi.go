package main

//#cgo CFLAGS: -I.
//#include "ocgapi.h"
import "C"
import "unsafe"

func main() {}

// TODO: demo for function pointer, change as what in ocgapi
type scriptReader func(*C.char, *C.int) *C.byte
type cardReader func(C.uint32, *C.card_data) C.uint32
type messageHandler func(unsafe.Pointer, C.uint32) C.uint32

func ScriptReader(reader C.script_reader) scriptReader {
	return func(scriptName *C.char, len *C.int) *C.byte {
		return C.read_script(reader, scriptName, len)
	}
}

func CardReader(reader C.card_reader) cardReader {
	return func(code C.uint32, data *C.card_data) C.uint32 {
		return C.read_card(reader, code, data)
	}
}

func MessageHandler(f C.message_handler) messageHandler {
	return func(pduel unsafe.Pointer, messageType C.uint32) C.uint32 {
		return C.handle_message(f, pduel, messageType)
	}
}

//export set_script_reader
func set_script_reader(f C.script_reader) {
}

//export set_card_reader
func set_card_reader(f C.card_reader) {}

//export set_message_handler
func set_message_handler(f C.message_handler) {}

//export create_duel
func create_duel(seed C.uint32) C.ptr {
	return 0
}

//export start_duel
func start_duel(pduel C.ptr, options C.uint32) {}

//export end_duel
func end_duel(pduel C.ptr) {}

//export set_player_info
func set_player_info(pduel C.ptr, playerid C.int32, lp C.int32, startcount C.int32, drawcount C.int32) {
}

//export get_log_message
func get_log_message(pduel C.ptr, buf *C.byte) {
}

//export get_message
func get_message(pduel C.ptr, buf *C.byte) C.int32 {
	return 0
}

//export process
func process(pduel C.ptr) C.int32 {
	return 0
}

//export new_card
func new_card(
	pduel C.ptr, code C.uint32, owner C.uint8,
	playerid C.uint8, location C.uint8, sequence C.uint8,
	position uint8,
) {
}

//export new_tag_card
func new_tag_card(
	pduel C.ptr, code C.uint32, owner C.uint8,
	location C.uint8,
) {
}

//export query_card
func query_card(
	pduel C.ptr, playerid C.uint8, location C.uint8,
	sequence C.uint8, query_flag C.int32, buf *C.byte,
	use_cache C.int32,
) C.int32 {
	return 0
}

//export query_field_count
func query_field_count(pduel C.ptr, playerid C.uint8, location C.uint8) C.int32 {
	return 0
}

//export query_field_card
func query_field_card(
	pduel C.ptr, playerid C.uint8, location C.uint8,
	query_flag C.int32, buf *C.byte, use_cache C.int32,
) C.int32 {
	return 0
}

//export query_field_info
func query_field_info(pduel C.ptr, buf *C.byte) C.int32 {
	return 0
}

//export set_responsei
func set_responsei(pduel C.ptr, value C.int32) {}

//export set_responseb
func set_responseb(pduel C.ptr, buf *C.byte) {}

//export preload_script
func preload_script(pduel C.ptr, script *C.char, len C.int32) C.int32 {
	return 0
}
