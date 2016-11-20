# Polaroid

The Polaroid Virtual Machine, written in Go

## Virtual Machine registers

| Registers | Description | Stored in AR |
|---:|:---|:---:|
| `pc` |  Program Counter | no |
| `esp` |  Stack Pointer | no |
| `tp` |  This pointer | no |
| `me` |  Shared memory pointer | no |
| `eax` `ebx` `ecx` `edx` | General purpose registers | no |
| `lpc` | Last Program Counter | yes |
| `ltp` |  Last This Pointer | yes |
| `r0 ... r31` | Local variable registers | yes |

This means that the machine itself provides the `pc`, `esp`, `tp` and `me` registers
but `eax`, `ebx`, `ecx`, `edx`, `ltp` are referenced in each Activation Record.
That is, each time an AR is generated, `pc` is transfered in the new `tpc` and same
goes for `tp` being copied in `ltp`. `r1` and all successive registers are allocated to
the parameters passed to the function. `r0` holds the number of arguments that were passed.

## Assembler instructions

Each assembler instruction has from 0 to 3 operands that can either be registers, constants
or memory references (pointers). Each operand is label using this convention :

* `A` is the first operand
* `B` is the second operand
* `C` is the third operand

For a whole instruction must hold in 32 bits, each operand must be 8 bits long. An operand binary representation is as follows :

```
bits   0      2                  8
       | type |       value      | 
```

Where type is :

* `00` if the operand refers to a register
* `01` if the operand refers to a constant (up to 2^6 - 1) in the pool
* `10` if the operand refers to a the memory address stored in the given register

Let's analyze the `add eax, 21, [r2]` instruction that is just `eax = k[21] + *r2;`, where `k` is the constant pool.
Assuming the `add` opcode is right, it is compiled down to:

```
   OPCODE  A          B         C 

      add  eax        21        [r2]
0010 1100  0000 0100  0101 0101 1000 1101
           ^^         ^^        ^^
            |          |         |
            |          |        10 : refers to what is stored at the address stored in r2
            |         01 : refers to a constant (the 21st)
           00 : refers to a register
```

## Assembler specifications

* `R` is the symbol for _any register_
* `K` is the symbol for _constant index_
* `A` is the symbol for _address_

| Instruction  name | Operands | A | B | C | Description |
|:-----------------:|:---:|:-----:|:-----:|:-----:|:---|
| **Memory related**|     |       |       |       |    |
| `mov`             |   2 | `RA`  | `RKA` | -     | Moves the 2nd operand value or pointer to first operand |
| `swp`             |   2 | `RA`  | `RA`  | -     | Swaps the two operands values |
| **Arithmetic**    |     |       |       |       |    |
| `add`             |   3 | `RA`  | `RKA` | `RKA` | Adds two integers from constants or registers into the first register |
| `sub`             |   3 | `RA`  | `RKA` | `RKA` | Subtracts two integers |
| `mul`             |   3 | `RA`  | `RKA` | `RKA` | Multiplies two integers |
| `pow`             |   2 | `RA`  | `RKA` | -     | Exponentiation |
| `div`             |   3 | `RA`  | `RK`  | `RK`  | Divides two integers |
| `rem`             |   3 | `RA`  | `RK`  | `RK`  | Computes the remainder |
| `inc`             |   1 | `RA`  | -     | -     | Increments a register |
| `dec`             |   1 | `RA`  | -     | -     | Decrements a register |