#ifndef CALLBACKS_H__
#define CALLBACKS_H__

#include "R.h"

int rango_read_console(const char * p, unsigned char * buf, int buflen, int add_history);

void rango_write_console(const char* s, int bufline, int otype);

#endif /* end of include guard: CALLBACKS_H__ */
