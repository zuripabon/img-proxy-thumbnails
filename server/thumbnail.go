package server

import (
	"bytes"
	"errors"
	"github.com/disintegration/imaging"
	"image"
	"image/jpeg"
	"image/png"
	"mime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func ThumbnailHandler(prefix string, w http.ResponseWriter, r *http.Request) {

	queryValues := r.URL.Query()

	remoteImageUrl, err := getUrlParam(queryValues.Get("url"))

	if err != nil {
		sendError(w, errors.New("INVALID_IMAGE_URL"), http.StatusBadRequest)
		return
	}

	thumbnailSize, err := getSizeParam(queryValues.Get("size"))

	if err != nil {
		sendError(w, errors.New("INVALID_IMAGE_SIZE"), http.StatusBadRequest)
		return
	}

	remoteImage, imageType, err := readImageFromUrl(remoteImageUrl)

	if err != nil {
		sendError(w, err, http.StatusBadRequest)
		return
	}

	thumbnailImage := imaging.Resize(remoteImage, thumbnailSize[0], thumbnailSize[1], imaging.Lanczos)

	sendImage(w, thumbnailImage, imageType)

}

func getUrlParam(urlInput string) (string, error) {
	u, err := url.Parse(urlInput)
	if err != nil {
		return "", err
	}
	return u.String(), nil
}

func getSizeParam(size string) ([2]int, error) {

	var sizes [2]int

	separator := "x"
	initialSize := 100

	if size == "" {
		sizes[0] = initialSize
		sizes[1] = initialSize
		return sizes, nil
	}

	if !strings.Contains(size, separator) {

		sizeNumber, err := strconv.Atoi(size)

		if err != nil {
			return sizes, err
		}

		sizes[0] = sizeNumber
		sizes[1] = sizeNumber

		return sizes, nil
	}

	sizesAsStrings := strings.Split(size, separator)

	width, err := strconv.Atoi(sizesAsStrings[0])

	if err != nil {
		return sizes, err
	}

	sizes[0] = width

	height, err := strconv.Atoi(sizesAsStrings[1])

	if err != nil {
		return sizes, err
	}

	sizes[1] = height

	return sizes, nil

}

func readImageFromUrl(url string) (image.Image, string, error) {

	remoteImageResponse, err := http.Get(url)

	if err != nil {
		return nil, "", errors.New("COULD_NOT_READ_IMAGE_FROM_URL")
	}

	defer remoteImageResponse.Body.Close()

	decodedImage, _, err := image.Decode(remoteImageResponse.Body)

	if err != nil {
		return nil, "", errors.New("COULD_NOT_DECODE_IMAGE")
	}

	imageType, _, err := mime.ParseMediaType(remoteImageResponse.Header["Content-Type"][0])

	if err != nil {
		return nil, "", errors.New("COULD_NOT_READ_CONTENT_TYPE")
	}

	return decodedImage, imageType, nil

}

func sendImage(w http.ResponseWriter, img *image.NRGBA, imageType string) {

	switch imageType {
	case "image/png":
		w.Header().Set("Content-Type", imageType)
		sendPNG(w, img)
	case "image/jpg", "image/jpeg":
		w.Header().Set("Content-Type", imageType)
		sendJPG(w, img)
	default:
		sendError(w, errors.New("CONTENT_TYPE_IMAGE_NOT_SOPPORTED"), http.StatusBadRequest)
	}

}

func sendJPG(w http.ResponseWriter, img *image.NRGBA) {

	buffer := new(bytes.Buffer)

	if err := jpeg.Encode(buffer, img, nil); err != nil {
		sendError(w, errors.New("COULD_NOT_ENCODE_IMAGE"), http.StatusBadRequest)
	}

	sendBufferResponse(w, *buffer)

}

func sendPNG(w http.ResponseWriter, img *image.NRGBA) {

	buffer := new(bytes.Buffer)

	if err := png.Encode(buffer, img); err != nil {
		sendError(w, errors.New("COULD_NOT_ENCODE_IMAGE"), http.StatusBadRequest)
	}

	sendBufferResponse(w, *buffer)

}

func sendBufferResponse(w http.ResponseWriter, buffer bytes.Buffer) {

	// Allow pre-flight cors headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Range, Content-Disposition, Content-Type, ETag")

	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))

	if _, err := w.Write(buffer.Bytes()); err != nil {
		sendError(w, errors.New("COULD_NOT_SEND_IMAGE"), http.StatusBadRequest)
		return
	}

}

func sendError(w http.ResponseWriter, message error, statusCode int) {

	w.WriteHeader(statusCode)
	w.Write([]byte(message.Error()))

}
