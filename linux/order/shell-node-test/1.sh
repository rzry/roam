#!/bin/bash
case $1 in
    a)       # 接受 a
        echo "a"
        ;;
    b|c)     # 接受 b 或 c
        echo "b or c"
        ;;
    ?)       # 接受任意一个字符
        echo "chat default case"
        ;;
    *)       # 接受任意的字符或字符串
        echo "default case"
        ;;
esac
