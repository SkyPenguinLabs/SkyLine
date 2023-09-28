#!/bin/bash

url1="https://raw.githubusercontent.com/SkyPenguinLabs/SkyLine-Dependant/main/ConstantIdentifiersStandard.json"
url2="https://raw.githubusercontent.com/SkyPenguinLabs/SkyLine-Dependant/main/FileSignatures.json"
url3="https://github.com/SkyPenguinLabs/SkyLine-Dependant/raw/main/SkyLine"
output_dir="/tmp/SkyLine_Dependable"

mkdir -p "$output_dir"


FetchDeps() {
    local url="$1"
    local filename="$2"
    wget -q "$url" -O "$filename"
    if [ $? -eq 0 ]; then
        echo "[+] Downloaded: $filename"
    else
        echo "[!] Failed to download: $filename"
    fi
}


FetchDeps "$url1" "$output_dir/ConstantIdentifiersStandard.json"
FetchDeps "$url2" "$output_dir/FileSignatures.json"
FetchDeps "$url3" "SkyLine" 
sudo chmod +x ./SkyLine ; sudo mv SkyLine /usr/bin
