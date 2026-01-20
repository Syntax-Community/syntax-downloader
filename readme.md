# SYNTAX-DOWNLOADER

Universal CLI video/audio downloader untuk berbagai platform social media.

### Fitur

- Video dengan kualitas full HD
- Extract audio ke MP3 dengan kualitas terbaik  
- Custom folder output
- Otomatis fallback ke folder `Downloads` di Linux/Android  
---

## Instalasi

### Linux / Termux

1. **Compile dari source:**

```bash
# Linux
git clone https://github.com/Syntax-Community/syntax-downloader.git
cd syntax-downloader
go build -o get
sudo mv get /usr/local/bin/

#Termux
git clone https://github.com/Syntax-Community/syntax-downloader.git
cd syntax-downloader
go build -o get
mv get ~/../usr/bin/
chmod +x /usr/local/bin/get
```
2. **Archlinux:**
```bash
sudo pacman -S syntax-downloader
```