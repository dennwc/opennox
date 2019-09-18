#include "vqa_file.h"

extern "C" SDL_Surface* movieSurface;
extern "C" Uint32 g_format;

extern "C" int PlayMovieCallback(byte* frame, dword cx, dword cy);

extern "C" int PlayMovie(char* filename)
{
    char* backslash = strchr(filename, '\\');
    if (backslash != NULL)
    {
        backslash[0] = '/';
    }
    Cvqa_file file(filename);
    file.register_decode(&PlayMovieCallback);
    file.post_open();
    bool valid = file.is_valid();
    printf("Movie filename: %s, is valid: %d\n", filename, valid);
    if (valid)
    {
        movieSurface = SDL_CreateRGBSurfaceWithFormat(0, file.get_cx(), file.get_cy(), 16, g_format);
        file.extract_both();
        movieSurface = NULL;
    }
    return 0;
}
