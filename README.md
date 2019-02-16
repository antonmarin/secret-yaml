# secret-yaml
Secret yaml values

## Goal

Lightweight tool to secret yaml values, 
which you can quickly install right inside your pipeline

## Usage

### Install

`wget sourceFile > target`

### Encrypt

`syml encrypt path/to/decryptedFile.yaml --secret=${YOUR_SECRET} > 
path/to/encryptedYaml.yml`

### Decrypt

`syml decrypt path/to/encryptedYaml.yaml --secret=${YOUR_SECRET} > 
path/to/decryptedFile.yml`

### Generate secret

`syml generate-secret` outputs generated secret. 
Store it somewhere to use later `${YOUR_SERCET}` 
