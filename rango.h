#ifndef RANGO_H__
#define RANGO_H__

#include "R.h"
#include <stdlib.h>
#include <string.h>

#ifndef _WIN32
#include <dlfcn.h>
#else
#define WIN32_LEAN_AND_MEAN 1
#include <windows.h>
#endif

char* rango_last_loaded_symbol(void);
char* rango_dl_error_message(void);
int rango_load(const char* rhome);
int rango_is_initialized(void);
int rango_load_symbols(void);
int rango_load_constants(void);


int rango_init(int ac, char **av);
void rango_setuploop(void);
void rango_runloop(void);
void rango_set_callback(char* name, void* cb);

#endif /* end of include guard: RANGO_H__ */
