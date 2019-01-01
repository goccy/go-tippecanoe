#ifndef __H_BRIDGE__
#define __H_BRIDGE__

#ifdef __cplusplus
extern "C" {
#endif

#include <stdlib.h>

extern void init_cpus();
extern void bridge_json_free(void *o);
extern void *bridge_mbtiles_open(char *dbname, char **argv, int forcetable);
extern void bridge_mbtiles_close(void *outdb, const char *pgm);
extern void *bridge_parse_filter(const char *s);
extern void *bridge_read_filter(const char *fname);
extern void bridge_set_projection_or_exit(const char *optarg);

#ifdef __cplusplus
}
#endif

#endif /* __H_BRIDGE__ */
