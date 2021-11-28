#include "ocgapi.h"


byte* read_script(script_reader reader, const char* script_name, int* len) {
    return reader(script_name, len);
}


uint32 read_card(card_reader reader, uint32 code, card_data* data) {
    return reader(code, data);
}

uint32 handle_message(message_handler handler, void* pduel, uint32 message_type) {
    return handler(pduel, message_type);
}