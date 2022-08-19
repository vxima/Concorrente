#include<bits/stdc++.h>
//biblitecas de Socket
#include<arpa/inet.h>
#include<netinet/in.h>
#include<sys/socket.h>
//bibliotecas pra getch
#include<termios.h>
#include<unistd.h>
//bibliotecas de C
#include<string.h>
#include<stdlib.h>
#include<stdio.h>
//bibliteca pra thread
#include<chrono>
#include<thread>
#include <iostream>


#define PORT 8889
#define IP "127.0.0.1"

using namespace std;

//procurei na internet o getch para o linux 
int getch( ) {
  struct termios oldt, newt;
  int ch;
  tcgetattr( STDIN_FILENO, &oldt );
  newt = oldt;
  newt.c_lflag &= ~( ICANON | ECHO );
  tcsetattr( STDIN_FILENO, TCSANOW, &newt );
  ch = getchar();
  tcsetattr( STDIN_FILENO, TCSANOW, &oldt );
  return ch;
}


int main(){
    
    int server_fd , new_socket , valread;
    struct sockaddr_in address;
    int opt = 1;
    int addrlen = sizeof(address);
    
    
    //criando o socket 
    server_fd = socket(AF_INET , SOCK_STREAM , 0);
    if (server_fd == 0){
        cout << "Erro ao criar o socket" << endl;
        exit(1);
    }

    //forçando socket para a porta 8889
    if(setsockopt(server_fd , SOL_SOCKET , SO_REUSEADDR | SO_REUSEPORT , &opt , sizeof(opt))){
        cout << "Setsockopt"<< endl;
        exit(1);
    }

    //conectar a porta 8889 e IP 127.0.0.1

    address.sin_family = AF_INET;
    address.sin_addr.s_addr = inet_addr(IP);
    address.sin_port = htons(PORT);

    //bind o ip e a porta 

    if(bind(server_fd , (struct sockaddr *) &address , sizeof(address))<0){
        cout << "Erro bind" << endl;
        exit(1);
    }

    //listen 
    

    //conexao
    while(true){
        if(listen(server_fd , 1) >= 0){
            //cout << "sucesso na escuta" << endl;
            
            if ((new_socket = accept(server_fd, (struct sockaddr *)&address, (socklen_t*)&addrlen))>= 0) { 
                cout << "Sucesso em aceitar a conexao" << endl; 
                
            
                int movement=0;
                valread = read( new_socket , buffer, 1024); //valor lido do buffer
                //cout << valread << endl; 
                while(valread){
                    //cout << "entrou no while do valread" << endl;
                    
                    valread = read( new_socket , buffer, 1024);
                }
                cout << "Conexao encerrada" << endl;
            }
        }
    }
    return 0;
}
