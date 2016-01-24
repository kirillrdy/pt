#include <stdio.h>
#include <time.h>
#include <stdlib.h>
#include <X11/Xlib.h> // Every Xlib program must include this
#include <assert.h>   // I include this to test return values the lazy way
#include <unistd.h>   // So we got the profile for 10 seconds
#define NIL (0)       // A name for the void pointer

int random_val() {
  int r = rand();
  return r % 65000;
}

// Global
Display *display;
Colormap screen_colormap;
GC graphics_context;
Window window;

void set_pixel(int x, int y, int r, int g, int b) {
  XColor xcolour;
  xcolour.red = r;
  xcolour.green = g;
  xcolour.blue = b;
  xcolour.flags = DoRed | DoGreen | DoBlue;
  XAllocColor(display, screen_colormap, &xcolour);
  XSetForeground(display, graphics_context, xcolour.pixel);
  XDrawPoint(display, window, graphics_context, x, y);
}

void create_window(int window_width, int window_height) {

  display = XOpenDisplay(NULL);
  assert(display);

  int blackColor = BlackPixel(display, DefaultScreen(display));
  // Create the window

  window = XCreateSimpleWindow(display, DefaultRootWindow(display), 0, 0,
                               window_width, window_height, 0, blackColor,
                               blackColor);

  // We want to get MapNotify events
  XSelectInput(display, window, StructureNotifyMask);

  // "Map" the window (that is, make it appear on the screen)
  XMapWindow(display, window);

  // Create a "Graphics Context"
  graphics_context = XCreateGC(display, window, 0, NIL);
  screen_colormap = DefaultColormap(display, DefaultScreen(display));

  // Wait for the MapNotify event

  for (;;) {
    XEvent e;
    XNextEvent(display, &e);
    if (e.type == MapNotify)
      break;
  }
}

void flush() {
  XFlush(display);
}
