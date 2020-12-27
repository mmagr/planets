# Planetary weather forecast

This repo holds a simple implementation of a weather forecast logic for some galaxy composed of
three planets. The weather in all of them is based on their position relative to each other and
the galaxy's sun.

## Building

The easiest way to build this project is to use its provided makefile. You should only
require `go` (>= `1.15`) and `make` installed on your machine.

To build:
```
make build
```

Once that is done, you should find the binary created in `$PWD/bin/service`.

When run, this binary will fire up a HTTP server that exposes an API for calculating the expected
metereologic conditions for the planets (as they all share the same forecast).

## API usage

Once started, one can use the API as follows (will assume `curl` is available):

```
curl -sSL localhost:8080/clima?dia=123
```

> Please note that `localhost` might need to be replaced with the IP address or DNS entry of the
machine where the process is running.

In the sample above, one expects the forecast for the 123th day to be returned:

```
{
  "dia": 123,
  "clima": "nublado"
}
```

## Notes

This implementation makes some assumptions.

First, all planetary orbits are expected to be circular. Each planet may have its own angular speed,
and may move clockwise or counter-clockwise.

The system's sun is fixed at the origin of the 2D plane that holds all planets, and all planets'
orbits are centered on it.

By default, on `Day 0` all planets are aligned with the sun along the `x` axis of the cartesian
system centered on the sun. This (the planets initial position) may be overriden through
configuration, along with their angular speeds, direction and orbit radius.

## Known issues

This is currently only evaluating the expected forecast in terms of days. Some states however are
very short lived (time wise) - specially draughts and perfect weather conditions, as they require
all planets (and the sun in the earlier scenario) to be aligned. This is currently assuming that
such alignment is only observed at integer values of time `d` and when they happen on such discrete
conditions, the entire day is marked as having the forecast. Notice this also means that depending
on the initial position and speed values provided, one may never notice either of those scenarios.

To make matters worse, this implementation also relies on some not that complicated, but also not
that trivial math to work - using `float64` to do so. This means on top of the rounding errors
associated with the operations involved themselves, we also incur on some extra
representation-related errors that originate on the set of values one can effectively represent
with `float64`. This, combined with the granularity of evaluation mentioned earlier, make for some
quite hard to achieve scenarios, specially for the perfect weather condition.

## Pushing to cf

With a valid cloudfoundry environment in place, one can deploy this using:

```
cf push -b go_buildpack -m 64M --random-route
```

One such environment is available at:

```
planets-egregious-sable-dj.cap.explore.suse.dev
```

To access (requires `curl`):

```
curl -sSL planets-egregious-sable-dj.cap.explore.suse.dev/clima?dia=123
```