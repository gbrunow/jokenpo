// You can edit this code!
// Click here and start typing.
package main

import ("fmt"
        "net"
        "encoding/binary"
      )

func printMove(move int64){
  switch move {
    case 0:
      fmt.Println("PEDRA")
    case 1:
      fmt.Println("PAPEL")
    case 2:
      fmt.Println("TESOURA")
  }
}

func printResult(result int64, pc int64){
  fmt.Print("CPU jogou: ")
  printMove(pc)
  if result == 0 {
    fmt.Println("Empate.");
  } else if result == 1{
    fmt.Println("Voce ganhou.")
  } else {
    fmt.Println("Voce perdeu")
  }
  fmt.Println()
}

func main() {
  fmt.Println("CLIENT")
  var move int64 = -1
  var pc int64
  var result int64

  fmt.Println("0 - PEDRA");
  fmt.Println("1 - PAPEL");
  fmt.Println("2 - TESOURA");
  fmt.Println();
  fmt.Println("3 - ENCERRAR");

  for move != 3 {
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        fmt.Println("Falha na conexão:", err)
        fmt.Println("4 - TENTAR NOVAMENTE");
        fmt.Scanf("%d", &move)
    } else {
      fmt.Print("Opção: ")
      fmt.Scanf("%d", &move)

      if move >= 0 && move <= 2 {
        fmt.Println()
        fmt.Print("Voce jogou: ")
        printMove(move)

        err = binary.Write(conn, binary.LittleEndian, &move)
        if err != nil{
          fmt.Println("Erro ao comunicar com servidor: ", err)
          continue
        }

        err = binary.Read(conn, binary.LittleEndian, &result)
        if err != nil{
          fmt.Println("Erro ao comunicar com servidor: ", err)
          continue
        }

        err = binary.Read(conn, binary.LittleEndian, &pc)
        if err != nil{
          fmt.Println("Erro ao comunicar com servidor: ", err)
          continue
        }
        printResult(result, pc)
      } else if move > 3{
        fmt.Println("Opção Inválida.");
      }
      conn.Close()
    }
  }
}
