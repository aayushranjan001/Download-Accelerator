# 🚀 Download-Accelerator

A high-performance download accelerator built with Go to speed up file downloads efficiently.

## 📌 Features
- ✅ Multi-threaded downloading
- ✅ Resume interrupted downloads
- ✅ Supports large file downloads
- ✅ Cross-platform support (Windows, macOS, Linux)

---

## 🚀 Getting Started

### 1️⃣ Prerequisites
Ensure you have the following installed:
- **Go** (>=1.21): [Download Here](https://go.dev/dl/)
- **Git** (if cloning from GitHub)

### 2️⃣ Clone the Repository
```sh
git clone https://github.com/aayushranjan001/Download-Accelerator.git
cd Download-Accelerator
```

### 3️⃣ Build the Project
For Windows:
```sh
go build -o download-accelerator.exe ./cmd/myapp/
```
For Linux/macOS:
```sh
go build -o download-accelerator ./cmd/myapp/
```

### 4️⃣ Run the Project
For Windows:
```sh
.\download-accelerator.exe
```
For Linux/macOS:
```sh
./download-accelerator
```

---

## 🔧 Configuration
The application can be configured using environment variables or a config file.

Example `.env` file:
```
DOWNLOAD_THREADS=4
TIMEOUT=30
LOG_LEVEL=info
```

---

## ⚡ Usage
### Download a File
```sh
download-accelerator -url "https://example.com/file.zip" -o "file.zip"
```

### Resume Download
```sh
download-accelerator -url "https://example.com/file.zip" -o "file.zip" --resume
```

### Available Flags
| Flag | Description | Example |
|------|------------|---------|
| `-url` | URL of the file to download | `-url "https://example.com/file.zip"` |
| `-o` | Output file name | `-o "myfile.zip"` |
| `--resume` | Resume an interrupted download | `--resume` |
| `-t` | Number of download threads | `-t 4` |

---

## 🛠 Development
### Run Locally (Without Building)
```sh
go run ./cmd/myapp
```

### Format Code
```sh
go fmt ./...
```

### Run Tests
```sh
go test ./...
```

---

## 🤝 Contributing
Contributions are welcome! Please follow these steps:
1. Fork the repository
2. Create a feature branch (`git checkout -b feature-new-feature`)
3. Commit your changes (`git commit -m "Add new feature"`)
4. Push to the branch (`git push origin feature-new-feature`)
5. Open a pull request

---

## 📜 License
This project is licensed under the **MIT License**. See `LICENSE` for details.

---

### 📧 Contact
For any issues or feature requests, please create an issue in the repository.

---
🚀 Happy Downloading!
