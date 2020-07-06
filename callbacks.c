#include "callbacks.h"

extern int GoReadConsole(char * p, unsigned char * buf, int buflen, int add_history);

int rango_read_console(const char * p, unsigned char * buf, int buflen, int add_history) {
    return GoReadConsole((char *) p, buf, buflen, add_history);
}
