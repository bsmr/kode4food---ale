---
title: "get"
date: 2019-04-06T12:19:22+02:00
description: "retrieves a value by key"
names: ["get"]
usage: "(get form key default?)"
tags: ["sequence"]
---
Returns the value within a sequence that is associated with the specified key. If the key does not exist within the sequence, then either the default value is returned or an error is raised.

#### An Example

```clojure
(def robert {:name "Bob" :age 45})
(get robert :address "wrong")
```

This example returns _"wrong"_ because the associative doesn't contain an :address property.