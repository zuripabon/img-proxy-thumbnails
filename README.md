# Thumbnail Image Proxy

This generates thumbnail images on the fly by requesting to the `/thumbnail` endpoint. It accepts `url` and `size` query params. The size can be `<width>x<height>` or can be a number for both width and height (i.e `size=150`, `size=100x80`). 

The url query param must point to a valid JPG or PNG image. It start up by default at port 9092. Use `IMAGE_PROXY_PORT` to change it. For example `http://localhost:9092/thumbnail?size=160x120&url=https://d1052pu3rm1xk9.cloudfront.net/fspt_640_640/cc370bb350187eb74737ab488feb8038d51fae29bd03430be1f9c3f8.jpg`

## Setup

To start up the server, move to the root directory and create the docker image by running `docker build -t image-proxy-server .` then run it by issuing `docker run -p 9092:9092 image-proxy-server`

## Todos

  * Goroutines support for thumbnail conversion
  * Use a message queue broker  
  * Caching Policies
