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
