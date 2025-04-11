# HEADER

Plain text test

---
# TABLE

| LEFT | b | rightish | 
|:---|:---:|---:|
| 1 | 2 | 3 | 
| 4 | 5 | 6 |

---

# LIST WITH CONTINUOUS NUMBERING

Numbering is continous, doesnt depend on level

1. Begining
- a
	- b
- c
	- d
6. e
	7. f
8. g
	9. h

---

# LIST WITH **NOT** CONTINUOUS NUMBERING

Numbering for ordered part depends on level of item. When you go down one level, index contunues from the last item of this level. If the next item after upper level item is met, the index for this level resets.

1.  Begining
- a
	- b
- c
	- d
2. e 	```Index doesn't reset after unordered part```
	1. f
3. g 	```Here index resets for level 2 ordered part```
	1. h
	2. X
		1. XAXAXA ```Go up to level 3```
		2. JAJAJA
	3. BACK TO LEVEL 2
		1. Cool ```After going back to level 2, index for 3rd level is reset and for level 2 is remained until we go back to level 1```
4. A
5. B

---

# QUOTE

This is a WIP. No leveling support yet

> Sweet Pumpkin Pie

---

# TEXT FORMAT

~~Strike through test~~

***Bold Italic test***

**Bold test**

*Italic test*

---

# CODE BLOCK

```go
package main
import "fmt"

func main() {
	fmt.Println("Hello, World!")
}
```
---

# LINK

[Google Test Link](https://google.com)

# ALERTS

>[!WARNING]
>Warning test

>[!IMPORTANT]
>Important test

>[!NOTE]
>Note test

>[!TIP]
>Tip test

>[!CAUTION]
>Caution test
