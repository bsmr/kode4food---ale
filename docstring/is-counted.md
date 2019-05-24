---
title: "counted?"
date: 2019-04-06T12:19:22+02:00
description: "tests whether the provided forms are counted sequences"
names: ["counted?", "!counted?", "is-counted"]
usage: "(counted? form+) (!counted? form+) (is-counted form)"
tags: ["sequence", "predicate"]
---
If all forms evaluate to a valid sequence than can report its length without counting, then this function will return _true_. The first non-counted sequence will result in the function returning _false_.

#### An Example

```clojure
(counted? '(1 2 3 4) [5 6 7 8])
```

This example will return _true_.

Like most predicates, this function can also be negated by prepending the `!` character. This means that all of the provided forms must not be valid counted sequences.

```clojure
(!counted? "hello" 99)
```

This example will return _false_ because strings are counted sequences.