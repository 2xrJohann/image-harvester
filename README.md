# image-harvester
Wrote this to migrate a discord image channel from one chord to the other. This bot just goes and gets all the images from whichever channel you want.

The discordgo library I used will let me get 100 messages at a time. If I had to scale this to a bigger channel I would but I only needed it to read 
about 900 images so I just did it one at a time

Uploader.go uploads all the images from a directory to a discord channel
