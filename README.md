# Polaroid

The Polaroid Virtual Machine, written in Go


| Registers | Description | Stord in AR |
|---:|:---|:---:|
| `pc` |  Program Counter | [ ] |
| `esp` |  Stack Pointer | [ ] |
| `tp` |  This pointer | [ ] |
| `me` |  Shared memory pointer | [ ] |
| `eax` `ebx` `ecx` `edx` | General purpose registers | [ ] |
| `lpc` | Last Program Counter | [x] |
| `ltp` |  Last This Pointer | [x] |
| `r0 ... r32` | Local variable registers | [x] |

This means that the machine itself provides the `pc`, `esp`, `tp` and `me` registers
but `eax`, `ebx`, `ecx`, `edx`, `ltp` are referenced in each Activation Record.