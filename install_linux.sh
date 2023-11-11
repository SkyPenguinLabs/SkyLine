#!/bin/bash


# Verify before creating directories
DirExistsQ() {
    if [ $# -ne 1 ]; then
        echo "[!] DEVELOPER ERROR -> FUNCTION MISUSE ( $0 <directory>)"
        exit 1
    fi
    local directory=$1
    if [ -d "$directory" ]; then
        return 0  #
    else
        return 1  
    fi
}

echo "[!]]] Checking if /usr/share/SkyLineDep EXISTS!"
if DirExistsQ "/usr/share/SkyLineDep"; then
    echo -e "\t WARNING: DIRECTORY ALREADY EXISTS....Skipping Creation"
else
    mkdir "/usr/share/SkyLineDep"
fi

# using git to download directories- way easier
# + we can verify the endpoints
Download_Directories() {
    if [ $# -lt 2 ]; then
        echo "[!] DEVELOPER ERROR -> FUNCTION MISUSE ( $0 <REPO> <destination directory>)"
        exit 1
    fi
    local repo_url=$1
    local directory=$2
    local destination=${3:-.}  
    local temp_dir=$(mktemp -d)
    echo "[!]]] Checking $destination$directory"
    if DirExistsQ $destination$directory; then 
        echo -e " \t WARNING: DIRECTORY ALREADY EXISTS....Skipping"
    else
        git clone "$repo_url" "$temp_dir"
        if [ $? -ne 0 ]; then
            echo "Error: Unable to clone the repository. Check the repository URL."
            exit 1
        fi
        mv "$temp_dir/$directory" "$destination"
        rm -rf "$temp_dir"
        echo "Downloaded $directory from $repo_url to $destination."
    fi
}

# We use WGET here over GIT because WGET provides a way easier 
# interface for single files rather than directories
Download_Files() {
    if [ $# -ne 2 ]; then
        echo "[!] DEVELOPER ERROR -> FUNCTION MISUSE ( $0 <file URL> <destination directory>)"
        exit 1
    fi
    local file_url=$1
    local destination=$2
    wget "$file_url" -O "$destination"
    if [ $? -ne 0 ]; then
        echo "Error: Unable to download the file. Check the file URL."
        exit 1
    fi
    echo "Installed file from $file_url to $destination."
}

Download_Directories https://github.com/SkyPenguinLabs/SkyLine.git Demonstrations /usr/share/SkyLineDep/
Download_Directories https://github.com/SkyPenguinLabs/SkyLine.git Primary /usr/share/SkyLineDep/
Download_Directories https://github.com/SkyPenguinLabs/SkyLine.git Utilities /usr/share/SkyLineDep/
Download_Directories https://github.com/SkyPenguinLabs/SkyLine.git Required /usr/share/SkyLineDep/
Download_Files https://github.com/SkyPenguinLabs/SkyLine/raw/main/SkyLine $(pwd)/SkyLine
chmod +x ./SkyLine
sudo mv SkyLine /usr/bin
