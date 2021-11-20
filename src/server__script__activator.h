#ifndef NOX_SERVER_SCRIPT_ACTIVATOR_H
#define NOX_SERVER_SCRIPT_ACTIVATOR_H

#include "defs.h"

bool nox_script_activatorCancel_51AD60(int id);
void nox_script_activatorCancelAll_51AC60();
void nox_script_activatorTimer_51ACA0(int df, int callback, int arg);
void nox_script_activatorClearObj_51AE60(nox_object_t* obj);
int nox_script_activatorSave_51AEA0();
int nox_script_activatorLoad_51AF80();
void nox_script_activatorRun_51ADF0();
void nox_script_activatorResolveObjs_51B0C0();

#endif // NOX_SERVER_SCRIPT_ACTIVATOR_H
