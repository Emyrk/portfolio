# Pinapple

## Abstract

Hackathon submission to HackUMass V, claiming the "Best Life Hack" prize. Pinapple is simply a peer to peer file sharer using Facebook authentication. The basic idea came around having an easy way to share files without needing to download software or have file size limitations. 

The name came from mispelling the placeholder name "Pineapple", and then sticking with it. The devpost describes the project and has a short demo video.

## Design

The design is rather simple. A GoLang based webserver acts as a coodinator, managing websockets with clients that connect. Clients can request connections with other clients, and send files over these websockets. Users on both ends see the same screen and can move/upload files in the shared box.