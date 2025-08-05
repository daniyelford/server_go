package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "127.0.0.1"
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok &&
			!ipnet.IP.IsLoopback() &&
			ipnet.IP.To4() != nil {

			ip := ipnet.IP.String()

			if strings.HasPrefix(ip, "192.") || strings.HasPrefix(ip, "10.") || strings.HasPrefix(ip, "172.") {
				return ip
			}
		}
	}
	return "127.0.0.1"
}

func main() {
	port := "8000"
	dir := "./"
	localIP := getLocalIP()

	fmt.Println("🟢 سرور لوکال Go در حال اجرا است.")
	fmt.Printf("🌐 بازدید از آدرس: http://%s:%s\n", localIP, port)

	// اجرای ssh tunnel در پس‌زمینه
	sshCmd := exec.Command("ssh", "-R", "80:"+localIP+":"+port, "ssh.localhost.run")
	fmt.Println(sshCmd)

	// هدایت خروجی ssh به ترمینال جاری برای دیدن لاگ
	sshCmd.Stdout = os.Stdout
	sshCmd.Stderr = os.Stderr

	// اجرای ssh به صورت پس‌زمینه
	err := sshCmd.Start()
	if err != nil {
		fmt.Println("❌ خطا در اجرای SSH tunnel:", err)
	}

	// راه اندازی سرور HTTP
	fs := http.FileServer(http.Dir(dir))
	http.Handle("/", fs)
	err = http.ListenAndServe("0.0.0.0:"+port, nil)
	if err != nil {
		fmt.Println("❌ خطا در اجرای سرور:", err)
		os.Exit(1)
	}
}
