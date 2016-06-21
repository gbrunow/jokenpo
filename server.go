// You can edit this code!
// Click here and start typing.
package main

import ("fmt"
        "math/rand"
        "time"
        "net"
        "encoding/binary"
        "os"
      )

func handleConnection(conn net.Conn) {
  var pc int64
  var player int64
  var result int64

  // read value from client
  err := binary.Read(conn, binary.LittleEndian, &player)

  if err != nil{
    fmt.Println("Falha na comunicação com cliente:", err)
  } else {
    fmt.Println(player)
    s1 := rand.NewSource(time.Now().UnixNano())
  	r1 := rand.New(s1)
  	pc = int64(r1.Intn(3))

    result = player - pc % 3
    if result < 0 {
      result += 3
    }

    err = binary.Write(conn, binary.LittleEndian, &pc)
    if err != nil{
      fmt.Println("Falha na comunicação com cliente:", err)
    }

    err = binary.Write(conn, binary.LittleEndian, &result)
    if err != nil{
      fmt.Println("Falha na comunicação com cliente:", err)
    }
  }
}

func main() {
  fmt.Println("SERVER")
  ln, err := net.Listen("tcp", ":8080")
  if err != nil {
    fmt.Println("Não foi possível iniciar o servidor, erro:", err);
    os.Exit(1)
  }

  for {
    conn, err := ln.Accept()
    if err != nil {
      fmt.Println("Problema ao estabelecer conexão", err);
    } else {
      go handleConnection(conn)
    }
  }
}
