# picvoter2

First, you should install buffalo:

```
go get -u -v github.com/gobuffalo/buffalo/buffalo
go get github.com/gobuffalo/pop/...
```

Make sure `$GOPATH/bin` is in your `PATH` so you can use the buffalo binaries. Next, create the following directories:

```
# choose a directory where imported and processed images will be stored
mkdir storage
cd storage
mkdir static && mkdir imports
```

Set the following env vars:

```
# make sure go uses dependency management
GO111MODULE=on

# the location you just created
STORAGE_LOCATION=storage

# wherever you want to import from
SCAN_LOCATIONS=/media/USB,/dev/sdb1
```

Now, run `yarn install`, and then `go get`.

Migrate the database:

```
soda create -a
soda migrate
```

At this point, you can start the server. To import images, use an http client like curl, httpie or postman, and do the following:  
(all requests have json formatted payload, `picvoter` represents the url of the service, if it's local it would be `localhost:3000`)

```
http GET picvoter/imports/candidates
```

This will show a list of folders that were found in your scan locations. Imagine you have a USB stick that will automatically be mounted at `/media/USB`, which you have in your scan locations. This USB stick contains the folders `Fotos Tag 1` and `Fotos Tag 2`, which will result in the following response:

```
{
    "/media/USB": [
        "Fotos Tag 1",
        "Fotos Tag 2",
    ],
}
```

If you want to import one of these, do the following request:

```
http POST picvoter/imports
{
    "location": "/media/USB",
    "directory": "Fotos Tag 1",
    "author": "Mein Name"
}
```

This will copy the photos from the scan location to the storage location and create an entry in the database. However, since these photos are probably full size they will need to be processed before they'll be served. It's recommended to do this at night, since the processing is done concurrenctly and takes a lot of resources, which totally killed our raspberry pi. To run the processing task, do the following:

```
buffalo task imports:process
```

This will process all pending imports and make the photos available.
