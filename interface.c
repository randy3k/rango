#include "interface.h"


int rango_init(int ac, char **av) {
  return Rf_initialize_R(ac, av);
}

void rango_setuploop(void) {
  setup_Rmainloop();
}
