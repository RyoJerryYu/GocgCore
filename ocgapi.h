/*
 * interface.h
 *
 *  Created on: 2010-4-28
 *      Author: Argon
 */

#ifndef OCGAPI_H_
#define OCGAPI_H_

typedef unsigned long uptr;
typedef unsigned long long uint64;
typedef unsigned int uint32;
typedef unsigned short uint16;
typedef unsigned char uint8;
typedef unsigned char byte;
typedef long ptr;
typedef long long int64;
typedef int int32;
typedef short int16;
typedef signed char int8;
typedef int BOOL;

typedef struct {
	uint32 code;
	uint32 alias;
	uint64 setcode;
	uint32 type;
	uint32 level;
	uint32 attribute;
	uint32 race;
	int32 attack;
	int32 defense;
	uint32 lscale;
	uint32 rscale;
	uint32 link_marker;
}card_data;

typedef byte* (*script_reader)(const char*, int*);
typedef uint32 (*card_reader)(uint32, card_data*);
typedef uint32 (*message_handler)(void*, uint32);

byte* read_script(script_reader reader, const char* script_name, int* len);
uint32 read_card(card_reader reader, uint32 code, card_data* data);
uint32 handle_message(message_handler handler, void* pduel, uint32 message_type);

#endif /* OCGAPI_H_ */