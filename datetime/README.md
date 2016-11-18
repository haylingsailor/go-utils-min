# datetime
A collection of cross-project date/time utilities.

## TimeProvider
Do not use the go built-in Time.Now(). Instead, inject into your app something
that implements interface **TimeProvider**  - for example **NowTimeProvider**
to get and manipulate the current date/time, so you can test aspects of the
system whose behaviour changes with the passage of time (e.g. expiring entities).
