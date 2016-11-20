# Polaroid

The Polaroid Virtual Machine, written in Go


| Registers | Description | Stord in AR |
|---:|:---|:---:|
| `pc` |  Program Counter | no |
| `esp` |  Stack Pointer | no |
| `tp` |  This pointer | no |
| `me` |  Shared memory pointer | no |
| `eax` `ebx` `ecx` `edx` | General purpose registers | no |
| `lpc` | Last Program Counter | yes |
| `ltp` |  Last This Pointer | yes |
| `r0 ... r32` | Local variable registers | yes |

This means that the machine itself provides the `pc`, `esp`, `tp` and `me` registers
but `eax`, `ebx`, `ecx`, `edx`, `ltp` are referenced in each Activation Record.