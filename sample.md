```
.import stdio

.data
    s : "Hello, World! "

%start
    .const 3
    .const 5
    add     a0,     r0,     r1
    call    printf, s,      a0
```