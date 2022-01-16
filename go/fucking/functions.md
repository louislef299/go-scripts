# The ***Fucking*** Project

This script makes it easy to work with docker in a repetitive manner. Docker has a lot of different functionalities, but the goal is to make it easy enough to work as a simple file system that is isolated from the root file system. We will be utilizing the docker API and golang to create this function.

## Functions

- `fucking -init`: This creates a new container with an ubuntu image
- `fucking -new=<image>`: Create a new container with a given image
- `fucking -restore=<imageID>`: Restores a saved container session
- `fucking -list`: Lists all active containers
- `fucking -save`: Saves the current container environment
- `fucking -saveAndClose`: Saves the environment and shuts it down
- `fucking -listSaved`: Lists all of your saved sessions and when they were committed
- `fucking -kill=<containerID>`: Kills a running container. `***WARNING***` *If* you kill a container and it was not saved, you will lose any progress made on the container. 

## Resources

- [Learn Docker in 100 seconds](https://www.youtube.com/watch?v=Gjnup-PuquQ)
- [Advanced Docker](https://www.youtube.com/watch?v=gAkwW2tuIqE)
- [Learn Go in 100 seconds](https://www.youtube.com/watch?v=446E-r0rXHI)