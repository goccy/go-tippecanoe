#include "bridge.h"
#include "main.hpp"
#include "mbtiles.hpp"
#include "jsonpull/jsonpull.h"
#include "evaluator.hpp"
#include "projection.hpp"

void bridge_json_free(void *o)
{
    json_free((json_object *)o);
}

void *bridge_mbtiles_open(char *dbname, char **argv, int forcetable)
{
    return (void *)mbtiles_open(dbname, argv, forcetable);
}

void bridge_mbtiles_close(void *outdb, const char *pgm)
{
    mbtiles_close((sqlite3 *)outdb, pgm);
}

void *bridge_parse_filter(const char *s)
{
    return (void *)parse_filter(s);
}

void *bridge_read_filter(const char *fname)
{
    return (void *)read_filter(fname);
}

void bridge_set_projection_or_exit(const char *optarg)
{
    set_projection_or_exit(optarg);
}
