package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"runtime"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

func main() {
        if runtime.GOOS == "windows" {
                fmt.Println("This tool might require administrator privileges on Windows.")
        }

        devices, err := pcap.FindAllDevs()
        if err != nil {
                log.Fatal(err)
        }

        fmt.Println("Available network interfaces:")
        for i, device := range devices {
                fmt.Printf("%d: %s (%s)\n", i, device.Name, device.Description)
        }

        var ifaceIndex int
        fmt.Print("Select interface to listen on: ")
        fmt.Scanln(&ifaceIndex)

        if ifaceIndex < 0 || ifaceIndex >= len(devices) {
                log.Fatal("Invalid interface selected.")
        }

        // Declare iface here, outside the if block
        var iface pcap.Interface
        if len(devices) > 0 {
                iface = devices[ifaceIndex]
        } else {
                log.Fatal("No interfaces found.")
        }

        handle, err := pcap.OpenLive(iface.Name, 65536, false, pcap.BlockForever)
        if err != nil {
                log.Fatal(err)
        }
        defer handle.Close()

        fmt.Println("Listening on", iface.Name, "...")

        knownHosts := make(map[string]string)

        packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
        for packet := range packetSource.Packets() {
                ipLayer := packet.Layer(layers.LayerTypeIPv4)
                if ipLayer != nil {
                        ip, _ := ipLayer.(*layers.IPv4)
                        srcIP := ip.SrcIP.String()
            dstIP := ip.DstIP.String()

                        arpLayer := packet.Layer(layers.LayerTypeARP)
                        if arpLayer != nil {
                                arp, _ := arpLayer.(*layers.ARP)
                                if arp.Operation == layers.ARPRequest || arp.Operation == layers.ARPReply {
                                        ipAddr := net.IP(arp.SourceHwAddress[:4]).String()
                                        macAddr := net.HardwareAddr(arp.SourceHwAddress).String()
                                        knownHosts[ipAddr] = macAddr
                                }
                        }
            if srcIP != "0.0.0.0"{
                if _, ok := knownHosts[srcIP]; !ok{
                    macAddr := getMacAddress(iface.Name, srcIP)
                    if macAddr != ""{
                        knownHosts[srcIP] = macAddr
                    }
                }
            }
            if dstIP != "255.255.255.255"{
                if _, ok := knownHosts[dstIP]; !ok{
                    macAddr := getMacAddress(iface.Name, dstIP)
                    if macAddr != ""{
                        knownHosts[dstIP] = macAddr
                    }
                }
            }
                }
                clearScreen()
                printKnownHosts(knownHosts)
        }
}

// ... (getMacAddress, printKnownHosts, clearScreen functions remain the same)
func getMacAddress(ifaceName, ipAddress string) string {
    interfaces, err := net.Interfaces()
    if err != nil {
        return ""
    }

    for _, iface := range interfaces {
        if iface.Name == ifaceName {
            addrs, err := iface.Addrs()
            if err != nil {
                continue
            }
            for _, addr := range addrs {
                ipnet, ok := addr.(*net.IPNet)
                if ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil && ipnet.IP.String() != ipAddress{
                    return iface.HardwareAddr.String()
                }
            }
        }
    }
    return ""
}

func printKnownHosts(hosts map[string]string) {
        fmt.Println("Discovered Hosts:")
        if len(hosts) == 0 {
                fmt.Println("No hosts discovered yet.")
                return
        }
        for ip, mac := range hosts {
                fmt.Printf("IP: %s, MAC: %s\n", ip, mac)
        }
}

func clearScreen() {
        if runtime.GOOS == "windows" {
                cmd := exec.Command("cmd", "/c", "cls")
                cmd.Stdout = os.Stdout
                cmd.Run()
        } else { // Linux/macOS
                cmd := exec.Command("clear")
                cmd.Stdout = os.Stdout
                cmd.Run()
        }
}