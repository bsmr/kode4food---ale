---
title: "rest"
date: 2019-04-06T12:19:22+02:00
description: "returns the rest of the sequence"
names: ["rest"]
usage: "(rest seq)"
tags: ["sequence"]
---

This function will return a sequence that excludes the first element of the specified sequence.

#### An Example

```scheme
(define x '(99 64 32 48))
(rest x)
```

This example will return _(64 32 48)_.
