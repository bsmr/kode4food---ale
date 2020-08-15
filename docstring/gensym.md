---
title: "gensym"
date: 2019-04-06T12:19:22+02:00
description: "creates a unique symbol, useful in macros"
names: ["gensym"]
usage: "(gensym str?)"
tags: ["symbol", "macro"]
---

If a string is provided, that string will be used to qualify the uniquely generated symbol. This function provides the underlying behavior for hash-tailed symbols in syntax-highlighting macros.

#### An Example

```scheme
(let [s (gensym "var")]
  (list 'ale/let [s "hello"] s))

;; is equivalent to
``(let [var# "hello"] var#)
```

This example will return _"hello"_.
