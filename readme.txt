go run server.go

or

ویندوز وقتی .exe می‌سازی و اجراش می‌کنی، بسته به تنظیمات ممکنه:

اجازه دسترسی به شبکه رو به اون فایل خاص نده.

یا پورت‌هایی مثل 8000 رو برای برنامه‌های ناشناس ببنده.

ولی وقتی با go run اجراش می‌کنی، چون پروسه از طریق go.exe هست و شناخته‌شده‌تره، اجازه دسترسی می‌ده.

✅ راه‌حل: اجازه کامل دسترسی به شبکه رو برای فایل .exe بده:

وارد بخش زیر شو:

mathematica
Copy
Edit
Control Panel > System and Security > Windows Defender Firewall > Allow an app or feature through Windows Defender Firewall
دکمه "Change settings" بزن

دکمه "Allow another app" بزن

فایل server.exe رو انتخاب کن

تیک‌ها رو برای Private و Public بزن

یا موقتاً برای تست فایروال رو خاموش کن و امتحان کن:

cmd
Copy
Edit
netsh advfirewall set allprofiles state off
2. 🔌 تفاوت سطح دسترسی بین اجرای مستقیم و go run
زمانی که .exe رو اجرا می‌کنی:

ممکنه به عنوان کاربر محدود (بدون دسترسی شبکه) اجرا بشه

ولی از VSCode یا Terminal با دسترسی بیشتر باشه

✅ راه‌حل: فایل .exe رو با دسترسی ادمین اجرا کن (Right Click → Run as Administrator)

3. ❌ اختلاف بین IPهایی که تونل می‌زنی
دستور زیر رو فقط زمانی بزن که مطمئنی سرور داره روی IP داخلی گوش می‌ده:

bash
Copy
Edit
ssh -R 80:192.168.1.50:8000 ssh.localhost.run
اگه مطمئن نیستی، بهتره بزنی:

bash
Copy
Edit
ssh -R 80:localhost:8000 ssh.localhost.run
یا حتی امن‌تر:

bash
Copy
Edit
ssh -R 80:127.0.0.1:8000 ssh.localhost.run
و توی Go مطمئن شو که روی 0.0.0.0:8000 گوش می‌ده نه فقط localhost.

✅ نسخه نهایی اصلاح‌شده‌ی server.go با اجرای خودکار تونل:
go
Copy
Edit
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

	// اجرای تونل به صورت خودکار (بدون بلاک شدن)
	go func() {
		cmd := exec.Command("C:\\Windows\\System32\\OpenSSH\\ssh.exe", "-R", "80:localhost:"+port, "ssh.localhost.run")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			fmt.Println("❌ خطا در اجرای تونل:", err)
		}
	}()

	fs := http.FileServer(http.Dir(dir))
	http.Handle("/", fs)
	err := http.ListenAndServe("0.0.0.0:"+port, nil)
	if err != nil {
		fmt.Println("❌ خطا در اجرای سرور:", err)
		os.Exit(1)
	}
}
اگر باز هم مشکل داشتی
از منوی استارت Windows Defender Firewall with Advanced Security رو باز کن و یه Inbound Rule بساز برای:

پورت 8000

یا فایل server.exe

go build server.go
./server

in git batsh
ssh -R 80:192.168.1.50:8000 ssh.localhost.run


or
وارد پنل مودمت شو:

تو مرورگر برو به: 192.168.1.1 یا 192.168.0.1

یوزر و پسورد معمولاً admin / admin یا روی مودم نوشته شده

بگرد دنبال قسمت:

Port Forwarding یا Virtual Server

(تو بعضی مودم‌ها داخل قسمت Advanced Settings یا NAT هست)

یه Port Forward جدید بساز:

External Port: 8000 (یا هرچی دوست داری)

Internal IP: 192.168.1.51 ← آی‌پی سیستم

Internal Port: 8000

Protocol: TCP (یا All)

ذخیره و ری‌استارت مودم 
and buy vps and static ip
