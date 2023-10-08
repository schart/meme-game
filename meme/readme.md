This is the application that allows us to upload the memes and meme texts needed for our project to the database, maintain them and distribute them against specific searches.

The meme uploader package is responsible for saving two important data sets, photos and texts, to the database.

When uploading a photo, the photo is uploaded to disk and at that time, i.e. before uploading, the name is changed to id, the changed name is published in the queue and the consumer receives it, saves it in the database and processes it.

When the text is uploaded, it is published to the queue and the consumer receives it, saves it to the database and processes it.

These data sets (photo, text) are independent of each other and need to be uploaded.

Text data can be saved by sending it via endpoint with access key and body parameter.

Multiple data upload feature is currently disabled.

Data can be saved in two ways - Users from pop-up endpoints. - Bots with an access token.

 
**Endpoints:**

        root endpoint: meme -> /meme/?

        POST -> meme/uploaders/upload-photo -> For upload a photo
            requirements body: keywords(file) = photo binary...

        POST -> meme/uploaders/upload-text -> For upload a text
            requirements body: keywords(Text)  = "text".

        GET ->  meme/items/photo -> For get photo by count
            requirements url parameter: keywords(count) =  0,1,2,3,4...?.

        GET ->  meme/items/text -> For get text by count
            requirements url parameter: keywords(count) =  0,1,2,3,4...?.



 