format ELF64 executable 3

; syscalls
    SYS_EXIT = 60 
    SYS_WRITE = 1

; constants
    LN = 10
    STD_OUT = 1

    pushw [PrintNumMsg]
    push PrintNumMsg_size
    call PrintAt

    ; push 10
    ; call PrintNum

    mov  rax, SYS_EXIT
    mov  rdi, 0
    syscall

PrintNum:
    pop rax
    mov byte [PrintNumFormatNum], al
    push PrintNumMsg
    push PrintNumMsg_size
    call PrintAt 

    cmp rax, 0
    jz PrintNumZero
    mov bl, 10
    div bl
    push rdx
    push rax
    call PrintNum
    pop rdx
    add dl, '0'
    mov byte [msg], dl
    call Print
    ret
PrintNumZero:
    mov byte [msg], '0'
    call Print
    ret

;     div 10
;     cmp rdx, 0
;     jz PrintZero
;
;     push byte dl
;     call PrintNum
;     pop byte dl, [msg]
; PrintZero:
;     mov byte [msg], al
;     call Print
; PrintNumRet:
;     ret
;

PrintAt:
    mov  rax, SYS_WRITE
    mov  rdi, STD_OUT
    ; xor rdx, rdx
    ; xor rsi, rsi
    pop rdx
    pop rsi
    ; mov  rsi, PrintNumMsg
    ; mov  rdx, PrintNumMsg_size
    ; pop rdx
    ; pop rsi
    syscall
    ret

Print:
    mov  rax, SYS_WRITE
    mov  rdi, STD_OUT
    mov  rsi, msg
    mov  rdx, msg_size
    syscall
    ret

msg:
    db ' ', LN, 0
    msg_size = $ - msg
PrintNumMsg:
    db 'PrintNum: rax '
    PrintNumFormatNum:
    db ' ', LN, 0
    PrintNumMsg_size = $ - PrintNumMsg


