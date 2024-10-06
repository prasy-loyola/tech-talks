format ELF64 executable 3

; syscalls
    SYS_EXIT = 60 
    SYS_WRITE = 1

; constants
    LN = 10
    STD_OUT = 1

    push 42
    push 23
    pop rax
    pop rdx
    add rax, rdx

    mov byte [msg], al 

    mov  rax, SYS_WRITE
    mov  rdi, STD_OUT
    mov  rsi, msg
    mov  rdx, msg_size
    syscall


    mov  rax, SYS_EXIT
    mov  rdi, 1
    syscall

msg:
    db ' ', LN, 0
    msg_size = $ - msg
