#include<termios.h>
#include<stdio.h>
#include<stdlib.h>

#define PASSWORD_LEN 8

int main(){
    struct termios initialrsettings, newrsettings;
    char password[PASSWORD_LEN + 1];

    tcgetattr(fileno(stdin), &initialrsettings);

    newrsettings = initialrsettings;
    // 删除标志位
    newrsettings.c_lflag &= ~ECHO;
    printf("Enter password: ");

    // TCSAFLUSH 丢弃用户在程序准备好读取数据之前输入的任何内容
    if(tcsetattr(fileno(stdin), TCSAFLUSH, &newrsettings) != 0) {
        fprintf(stderr,"Could not set attributes\n");
    } else {
        fgets(password, PASSWORD_LEN, stdin);
        tcsetattr(fileno(stdin), TCSANOW, &initialrsettings);
        fprintf(stdout, "\nYou entered %s\n", password);
    }
    exit(0);
}
