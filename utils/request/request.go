package request

import "github.com/gofiber/fiber/v2"

type Response struct {
	Status int
	Body   []byte
	Errs   []error
}

// Post 发送post请求
func Post(url string, header map[string]string, data fiber.Map) (Response, error) {
	a := fiber.Post(url)

	a.Set("Accept", "application/json")

	for k, v := range header {
		a.Set(k, v)
	}
	a.UserAgent("Gotutu by Gmero")
	a.JSON(data)
	var res Response

	if err := a.Parse(); err != nil {
		return res, err
	}

	code, body, errs := a.Bytes()

	if len(errs) > 0 {
		return res, errs[0]
	}

	res.Status = code
	res.Errs = errs
	res.Body = body

	return res, nil
}

// Put 发送put请求
func Put(url string, header map[string]string, data []byte) (Response, error) {
	a := fiber.Put(url)

	a.Set("Accept", "application/json")

	for k, v := range header {
		a.Set(k, v)
	}
	a.UserAgent("Gotutu by Gmero")

	file := &fiber.FormFile{Fieldname: "file", Name: "image", Content: data}
	a.FileData(file).MultipartForm(nil)
	var res Response

	if err := a.Parse(); err != nil {
		return res, err
	}

	code, body, errs := a.Bytes()

	if len(errs) > 0 {
		return res, errs[0]
	}

	res.Status = code
	res.Errs = errs
	res.Body = body

	return res, nil
}

// Get 发送Get请求
func Get(url string, header map[string]string) (Response, error) {
	a := fiber.Get(url)

	for k, v := range header {
		a.Set(k, v)
	}
	a.UserAgent("Gotutu by Gmero")

	var res Response

	if err := a.Parse(); err != nil {
		return res, err
	}

	code, body, errs := a.Bytes()

	if len(errs) > 0 {
		return res, errs[0]
	}

	res.Status = code
	res.Errs = errs
	res.Body = body

	return res, nil
}
