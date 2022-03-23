# Lefebvre Network Attached Storage

## Introduction

This project is a small project to detect directory changes and upload all files to a storage unit. For now, I am just using an S3 bucket for simplicity sake, but in the future, we would like to be able to have our own file storage unit in order to own the infrastructure required to upload pictures and videos to our own server(NAS) cheaply. 

In order for this project to be successfully finished, the program will have to run as a daemon job that will both detect files to a specified folder(Photos will be the default) and download files from the S3 bucket to local. I am going to use the [fsnotify](https://pkg.go.dev/github.com/howeyc/fsnotify) package and golang in order to script this out.

Pretty straighforward... going to add documentation here when the commands are fleshed out.