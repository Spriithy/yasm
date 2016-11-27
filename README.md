# SMALL - The **SM**all **A**ssembly-**L**ike **L**anguage

This is the *SMALL* language implementation repository for both 32 and 64 bits architectures.

## Registers

```
r0 ... r7   -> general purpose registers
pc          -> program counter
sp          -> stack pointer (local variables)
ic          -> interruption code (sys calls)
rx          -> the register  
```

# Opcodes

The instructions are stored on a 32 bits format, with 

```
      0x0 if using on nothing or integers
      0x1 if using floats
      |     
      |  op    ra   rb       extra bits ------> (signed if integer)
      v______ ____ ____ ___________________ 
    XXTO OOOO AAAA BBBB EEEE EEEE EEEE EEEE
    ^^        
    |         
    0x0 usual opcode
    0x1 immediate 16 (stored in this word)
    0x2 immediate 32 (one next word)
    0x3 immediate 64 (two next words)
```

There is the complete opcodes specification

```go
// general purpose opcodes

swi    0x00    (ra | imm)       interrupts the execution to perform the expected software interrupt           
mov    0x01    ra, (rb | imm)   moves the content from rb or the immediate value into register ra
mom    0x02    ra, rb           moves in memory content from the address stored in rb to the address stored in ra     
loa    0x03    ra, rb           load the value stored at the memory address of rb into ra
str    0x04    ra, (rb | imm)   stores the content of register rb or immediate value into ra

// arithmetic operations, each one updates the rx register to the operation result

add    0x05    ra, (rb | imm)   adds the second operand to the first register
sub    0x06    ra, (rb | imm)   subtract the second operand from the first register
mul    0x07    ra, (rb | imm)   multiplies the second operand with the first register
div    0x08    ra, (rb | imm)   divide the first register by the second operand
rem    0x09    ra, (rb | imm)   adds the second operand to the first register
bsl    0x0a    ra, (rb | imm)   bit shift left the first register by the second operand 
bsr    0x0b    ra, (rb | imm)   bit shift right the first register by the second operand

inc    0x0c    ra, (rb | imm)   increments the value stored in ra by the second operand signed value
dec    0x0d    ra, (rb | imm)   decrements the value stored in ra by the second operand signed value

and    0x0e    ra, (rb | imm)   bitwise and on first register and second operand 
ior    0x0f    ra, (rb | imm)   bitwise ior on first register and second operand
xor    0x10    ra, (rb | imm)   bitwise xor on first register and second operand
not    0x11    ra               bitwise not on register ra

// all conditional jump statements update the value of rx to the comparison result (except for jmp)
// rx = 1 if condition is met, rx = 0 otherwise 

jmp    0x12    imm              performs a relative jump, ie. PC += imm
jz     0x13    imm              relative jump if rx == 0
jnz    0x14    imm              relative jump if rx != 0
jeq    0x15    (ra | imm), imm  relative jump if ra == rx (or imm, likewise)
jne    0x16    (ra | imm), imm  relative jump if ra != rx    
jlt    0x17    (ra | imm), imm  relative jump if ra <  rx    
jle    0x18    (ra | imm), imm  relative jump if ra <= rx    
jgt    0x19    (ra | imm), imm  relative jump if ra >  rx    
jge    0x1a    (ra | imm), imm  relative jump if ra >= rx    

srl    0x1b    ra, rb           subroutine link
ret    0x1c    ra, rb           return from subroutine
```

## Legacy opcodes

```
eq     0x12    ra, (rb | imm)   compares ra and the second operand, sets rx to 1 unless they are different
neq    0x13    ra, (rb | imm)   sets rx to 0 unless they are different
lt     0x14    ra, (rb | imm)   sets rx to 1 unless ra >= (rb | imm)
leq    0x15    ra, (rb | imm)   sets rx to 1 unless ra >  (rb | imm)
gt     0x16    ra, (rb | imm)   sets rx to 1 unless ra <= (rb | imm)
geq    0x17    ra, (rb | imm)   sets rx to 1 unless ra <  (rb | imm)
```

I have effectively removed them from the Insctruction Set since they can easily been translated only using conditional jumps and mov
operations. Indeed, say we have the following higher-level code:

```go
if r0 == 0 {
    r1 = r1 + 1  
} else {
    r1 = r1 - 1
}
```

It would be translated to the following *SMALL* code:

```asm
eq    r0, 0
jnz      +2
inc   r1, 1
jmp      +1
dec   r1, 1
```

However, the use of `eq` could be replaced with a simple trick. This is used to reduce even more
the instruction set. I can afford up to 32 (from `0x00` to `0x1f`) opcodes so the opcode part in 
the binary format doesn't overflow 5 bits.

The previous piece of code is therefore invalid. Here is a rewrite:

```asm
mov   rx,  0
jne   r0, +2
inc   r1,  1 
jmp       +1
dec   r1,  1
```

Sameways, `a = (a == b)` could be:

```asm
eq    r0, r1    ; assume r0 = a, r1, = b
mov   r0, rx
```

but is instead written as:

```asm
mov   rx, r1    ; assume r0 = a, r1, = b
jeq   r0, +0    ; +0 tricks the interpreter to just not update the program counter
mov   r0, rx
```